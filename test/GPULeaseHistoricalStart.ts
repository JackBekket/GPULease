import { expect } from "chai";
import { network } from "hardhat";

const { ethers } = await network.connect();

describe("GPULease historical start timestamp", function () {
  let owner: any;
  let user: any;
  let provider: any;
  let treasury: any;
  let token: any;
  let wallet: any;
  let lease: any;

  const storagePrice = 1_000_000n;
  const computePrice = 2_000_000n;
  const duration = 30 * 24 * 60 * 60;

  beforeEach(async () => {
    [owner, user, provider, treasury] = await ethers.getSigners();

    const Token = await ethers.getContractFactory("MockERC20");
    token = await Token.deploy("Mock", "MOCK");

    const Wallet = await ethers.getContractFactory("GPULeaseWallet");
    wallet = await Wallet.deploy(token.target);

    const Lease = await ethers.getContractFactory("GPULease");
    lease = await Lease.deploy(wallet.target, treasury.address);
    await wallet.setLeaseManager(lease.target);

    await token.mint(user.address, ethers.parseEther("1000"));
    await token.connect(user).approve(wallet.target, ethers.MaxUint256);
    await wallet.connect(user).deposit(ethers.parseEther("1000"));
  });

  async function latestTimestamp() {
    const block = await ethers.provider.getBlock("latest");
    if (!block) throw new Error("Cannot fetch latest block");
    return block.timestamp;
  }

  async function startHistoricalLease(
    startTimestamp: number,
    leaseDuration = duration
  ) {
    return lease["startLease(uint256,uint256,uint256,uint256,address,address)"](
      startTimestamp,
      leaseDuration,
      storagePrice,
      computePrice,
      provider.address,
      user.address
    );
  }

  it("charges only storage between the supplied start and activation", async () => {
    const startTimestamp = (await latestTimestamp()) - 6 * 60 * 60;
    await startHistoricalLease(startTimestamp);

    const activationTime = await lease.leaseActivationTime(0);
    const [storageCost, computeCost] = await lease.calculateActualCost(0);

    expect(storageCost).to.equal(
      (activationTime - BigInt(startTimestamp)) * storagePrice
    );
    expect(computeCost).to.equal(0);
  });

  it("charges storage plus compute after activation", async () => {
    const startTimestamp = (await latestTimestamp()) - 6 * 60 * 60;
    await startHistoricalLease(startTimestamp);
    const activationTime = await lease.leaseActivationTime(0);

    await ethers.provider.send("evm_increaseTime", [2 * 60 * 60]);
    await ethers.provider.send("evm_mine", []);
    const now = BigInt(await latestTimestamp());
    const [storageCost, computeCost] = await lease.calculateActualCost(0);

    expect(storageCost).to.equal(
      (now - BigInt(startTimestamp)) * storagePrice
    );
    expect(computeCost).to.equal((now - activationTime) * computePrice);
  });

  it("stores both supplied start and on-chain activation timestamps", async () => {
    const startTimestamp = (await latestTimestamp()) - 60 * 60;
    const tx = await startHistoricalLease(startTimestamp);
    const receipt = await tx.wait();
    const activationBlock = await ethers.provider.getBlock(receipt.blockNumber);
    if (!activationBlock) throw new Error("Cannot fetch activation block");

    const leaseData = await lease.leases(0);
    expect(leaseData.startTime).to.equal(startTimestamp);
    expect(await lease.leaseActivationTime(0)).to.equal(
      activationBlock.timestamp
    );
  });

  it("emits start and activation timestamps in LeaseStarted", async () => {
    const startTimestamp = (await latestTimestamp()) - 60 * 60;
    const tx = await startHistoricalLease(startTimestamp);
    const receipt = await tx.wait();
    const activationBlock = await ethers.provider.getBlock(receipt.blockNumber);
    if (!activationBlock) throw new Error("Cannot fetch activation block");
    const event = receipt.logs
      .map((log: any) => {
        try {
          return lease.interface.parseLog(log);
        } catch {
          return null;
        }
      })
      .find((log: any) => log?.name === "LeaseStarted");

    expect(event.args.startTimestamp).to.equal(startTimestamp);
    expect(event.args.activationTimestamp).to.equal(activationBlock.timestamp);
    expect(event.args.duration).to.equal(duration);
  });

  it("rejects zero and future start timestamps", async () => {
    await expect(startHistoricalLease(0)).to.be.revertedWith(
      "invalid start timestamp"
    );

    const futureTimestamp = (await latestTimestamp()) + 60 * 60;
    await expect(startHistoricalLease(futureTimestamp)).to.be.revertedWith(
      "start in future"
    );
  });

  it("caps storage at duration and charges no compute if activation is after lease end", async () => {
    const shortDuration = 4 * 60 * 60;
    const startTimestamp = (await latestTimestamp()) - 10 * 60 * 60;
    await startHistoricalLease(startTimestamp, shortDuration);

    const [storageCost, computeCost] = await lease.calculateActualCost(0);
    expect(storageCost).to.equal(BigInt(shortDuration) * storagePrice);
    expect(computeCost).to.equal(0);
    expect(await lease.isSettlementDue(0)).to.equal(true);
  });

  it("keeps pause accounting limited to post-activation compute", async () => {
    const startTimestamp = (await latestTimestamp()) - 6 * 60 * 60;
    await startHistoricalLease(startTimestamp);
    const activationTime = await lease.leaseActivationTime(0);

    await ethers.provider.send("evm_increaseTime", [60 * 60]);
    await lease.pauseLease(0);
    await ethers.provider.send("evm_increaseTime", [2 * 60 * 60]);
    await lease.resumeLease(0);
    await ethers.provider.send("evm_increaseTime", [60 * 60]);
    await ethers.provider.send("evm_mine", []);

    const now = BigInt(await latestTimestamp());
    const leaseData = await lease.leases(0);
    const [storageCost, computeCost] = await lease.calculateActualCost(0);
    const computeDuration = now - activationTime - leaseData.pausedDuration;

    expect(storageCost).to.equal(
      (now - BigInt(startTimestamp)) * storagePrice
    );
    expect(computeCost).to.equal(computeDuration * computePrice);
  });

  it("does not count a pause extending beyond lease end twice", async () => {
    const shortDuration = 3 * 60 * 60;
    const startTimestamp = await latestTimestamp();
    await startHistoricalLease(startTimestamp, shortDuration);

    await ethers.provider.send("evm_increaseTime", [60 * 60]);
    await lease.pauseLease(0);
    await ethers.provider.send("evm_increaseTime", [5 * 60 * 60]);
    await lease.resumeLease(0);

    const leaseData = await lease.leases(0);
    const activationTime = await lease.leaseActivationTime(0);
    const [storageCost, computeCost] = await lease.calculateActualCost(0);
    expect(storageCost).to.equal(BigInt(shortDuration) * storagePrice);
    expect(computeCost).to.equal(
      (
        BigInt(startTimestamp + shortDuration) -
        activationTime -
        leaseData.pausedDuration
      ) * computePrice
    );
  });

  it("settles historical storage together with post-activation earnings", async () => {
    const startTimestamp = (await latestTimestamp()) - 6 * 60 * 60;
    await startHistoricalLease(startTimestamp);
    await ethers.provider.send("evm_increaseTime", [24 * 60 * 60]);
    await ethers.provider.send("evm_mine", []);

    await lease.settleLease(0);

    const [storageCost, computeCost] = await lease.calculateActualCost(0);
    const settlement = await lease.leaseSettlements(0);
    expect(settlement.providerPaid + settlement.feePaid).to.equal(
      storageCost + computeCost
    );
  });

});

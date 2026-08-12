import { expect } from "chai";
import { network } from "hardhat";
const { ethers } = await network.connect();

describe("GPULease duration edge cases", function () {
  let owner: any;
  let user: any;
  let provider: any;
  let treasury: any;

  let token: any;
  let wallet: any;
  let lease: any;

  async function startLeaseNow(
    target: any,
    duration: number,
    storagePricePerSecond: bigint,
    computePricePerSecond: bigint,
    providerAddress: string,
    userAddress: string
  ) {
    const block = await ethers.provider.getBlock("latest");
    if (!block) throw new Error("Cannot fetch latest block");
    const startTimestamp = block.timestamp + 1;
    await ethers.provider.send("evm_setNextBlockTimestamp", [startTimestamp]);
    return target.startLease(
      startTimestamp,
      duration,
      storagePricePerSecond,
      computePricePerSecond,
      providerAddress,
      userAddress
    );
  }

  beforeEach(async () => {
    [owner, user, provider, treasury] = await ethers.getSigners();

    const Token = await ethers.getContractFactory("MockERC20");
    token = await Token.deploy("Mock", "MOCK");
    await token.mint(user.address, ethers.parseEther("1000"));

    const Wallet = await ethers.getContractFactory("GPULeaseWallet");
    wallet = await Wallet.deploy(token.target);

    const Lease = await ethers.getContractFactory("GPULease");
    lease = await Lease.deploy(wallet.target, treasury.address);
    await wallet.setLeaseManager(lease.target);

    await token.connect(user).approve(wallet.target, ethers.parseEther("1000"));
    await wallet.connect(user).deposit(ethers.parseEther("1000"));
  });

  // Хелпер для расчёта expected payout по контрактной логике
  async function expectedPayout(leaseId: number) {
    const info = await lease.leases(leaseId);
    const block = await ethers.provider.getBlock("latest");
    if (!block) throw new Error("Cannot fetch latest block");
  
    const now = block.timestamp;
  
    let storageDuration = now - Number(info.startTime);
    if (storageDuration > Number(info.duration)) {
      storageDuration = Number(info.duration);
    }

    const activationTime = Number(await lease.leaseActivationTime(leaseId));
    const calculationTime = Math.min(
      now,
      Number(info.startTime) + Number(info.duration)
    );
    const computeDuration = Math.max(0, calculationTime - activationTime);
  
    let totalPaused = Number(info.pausedDuration);
    if (info.paused && calculationTime > Number(info.pausedAt)) {
      totalPaused += calculationTime - Number(info.pausedAt);
    }
    if (totalPaused > computeDuration) totalPaused = computeDuration;
  
    const activeDuration = computeDuration - totalPaused;
  
    const actualStorageCost = BigInt(storageDuration) * BigInt(info.storagePricePerSecond);
    const actualComputeCost = BigInt(activeDuration) * BigInt(info.computePricePerSecond);
    const actualTotal = actualStorageCost + actualComputeCost;
  
    const fee = (actualTotal * BigInt(info.leaseFeePercentage)) / 100n;
    const providerAmount = actualTotal - fee;
  
    return { fee, providerAmount, total: actualTotal };
  }

  it("duration < declared, no pause", async () => {
    const duration = 1000;
    const price = ethers.parseEther("0.001");

    await startLeaseNow(lease, duration, price, price, provider.address, user.address);

    // fast-forward чуть меньше duration
    await ethers.provider.send("evm_increaseTime", [900]);
    await ethers.provider.send("evm_mine", []);

    await lease.completeLease(0);

    const { fee, providerAmount } = await expectedPayout(0);
    expect(await wallet.userBalance(provider.address)).to.equal(providerAmount);
    expect(await wallet.userBalance(treasury.address)).to.equal(fee);
  });

  it("duration < declared, with pause", async () => {
    const duration = 1000;
    const price = ethers.parseEther("0.001");

    await startLeaseNow(lease, duration, price, price, provider.address, user.address);

    await lease.pauseLease(0);
    await ethers.provider.send("evm_increaseTime", [400]);
    await lease.resumeLease(0);

    await ethers.provider.send("evm_increaseTime", [400]);
    await ethers.provider.send("evm_mine", []);

    await lease.completeLease(0);

    const { fee, providerAmount } = await expectedPayout(0);
    expect(await wallet.userBalance(provider.address)).to.equal(providerAmount);
    expect(await wallet.userBalance(treasury.address)).to.equal(fee);
  });

  it("duration > declared, no pause", async () => {
    const duration = 1000;
    const price = ethers.parseEther("0.001");

    await startLeaseNow(lease, duration, price, price, provider.address, user.address);

    // fast-forward больше duration
    await ethers.provider.send("evm_increaseTime", [1500]);
    await ethers.provider.send("evm_mine", []);

    await lease.completeLease(0);

    const { fee, providerAmount } = await expectedPayout(0);
    expect(await wallet.userBalance(provider.address)).to.equal(providerAmount);
    expect(await wallet.userBalance(treasury.address)).to.equal(fee);
  });

  it("duration > declared, with pause", async () => {
    const duration = 1000;
    const price = ethers.parseEther("0.001");

    await startLeaseNow(lease, duration, price, price, provider.address, user.address);

    await lease.pauseLease(0);
    await ethers.provider.send("evm_increaseTime", [500]);
    await lease.resumeLease(0);

    await ethers.provider.send("evm_increaseTime", [800]);
    await ethers.provider.send("evm_mine", []);

    await lease.completeLease(0);

    const { fee, providerAmount } = await expectedPayout(0);
    expect(await wallet.userBalance(provider.address)).to.equal(providerAmount);
    expect(await wallet.userBalance(treasury.address)).to.equal(fee);
  });

  it("duration < declared, paused at completion", async () => {
    const duration = 1000;
    const price = ethers.parseEther("0.001");

    await startLeaseNow(lease, duration, price, price, provider.address, user.address);

    await lease.pauseLease(0);
    await ethers.provider.send("evm_increaseTime", [700]);
    await ethers.provider.send("evm_mine", []);

    await lease.completeLease(0); // paused на момент завершения

    const { fee, providerAmount } = await expectedPayout(0);
    expect(await wallet.userBalance(provider.address)).to.equal(providerAmount);
    expect(await wallet.userBalance(treasury.address)).to.equal(fee);
  });

  it("duration > declared, paused at completion", async () => {
    const duration = 1000;
    const price = ethers.parseEther("0.001");

    await startLeaseNow(lease, duration, price, price, provider.address, user.address);

    await lease.pauseLease(0);
    await ethers.provider.send("evm_increaseTime", [1500]);
    await ethers.provider.send("evm_mine", []);

    await lease.completeLease(0); // paused на момент завершения

    const { fee, providerAmount } = await expectedPayout(0);
    expect(await wallet.userBalance(provider.address)).to.equal(providerAmount);
    expect(await wallet.userBalance(treasury.address)).to.equal(fee);
  });
});

import { expect } from "chai";
import { network } from "hardhat";

const { ethers } = await network.connect();

describe("GPULease daily settlements", function () {
  let owner: any;
  let user: any;
  let provider: any;
  let treasury: any;
  let operator: any;
  let other: any;
  let token: any;
  let wallet: any;
  let lease: any;

  const duration = 3 * 24 * 60 * 60;
  const storagePrice = 1_000_000n;
  const computePrice = 2_000_000n;

  beforeEach(async () => {
    [owner, user, provider, treasury, operator, other] = await ethers.getSigners();

    const Token = await ethers.getContractFactory("MockERC20");
    token = await Token.deploy("Mock", "MOCK");

    const Wallet = await ethers.getContractFactory("GPULeaseWallet");
    wallet = await Wallet.deploy(token.target);

    const Lease = await ethers.getContractFactory("GPULease");
    lease = await Lease.deploy(wallet.target, treasury.address);
    await wallet.setLeaseManager(lease.target);
    await lease.setSettlementOperator(operator.address);

    await token.mint(user.address, ethers.parseEther("1000"));
    await token.connect(user).approve(wallet.target, ethers.parseEther("1000"));
    await wallet.connect(user).deposit(ethers.parseEther("1000"));
  });

  async function startLease(providerAddress = provider.address) {
    const block = await ethers.provider.getBlock("latest");
    if (!block) throw new Error("Cannot fetch latest block");
    const startTimestamp = block.timestamp + 1;
    await ethers.provider.send("evm_setNextBlockTimestamp", [startTimestamp]);
    await lease.startLease(
      startTimestamp,
      duration,
      storagePrice,
      computePrice,
      providerAddress,
      user.address
    );
  }

  it("allows the configured operator to settle once per day", async () => {
    await startLease();

    await expect(lease.connect(operator).settleLease(0)).to.be.revertedWith(
      "settlement not due"
    );

    await ethers.provider.send("evm_increaseTime", [24 * 60 * 60]);
    await ethers.provider.send("evm_mine", []);

    const before = await lease.settlementEntitlement(0);
    await expect(lease.connect(operator).settleLease(0))
      .to.emit(lease, "LeaseSettled");

    const settlement = await lease.leaseSettlements(0);
    expect(settlement.providerPaid).to.be.gte(before.providerEarned);
    expect(settlement.feePaid).to.be.gte(before.feeEarned);
    expect(await wallet.userBalance(provider.address)).to.equal(
      settlement.providerPaid
    );
    expect(await wallet.userBalance(treasury.address)).to.equal(
      settlement.feePaid
    );

    await expect(lease.connect(operator).settleLease(0)).to.be.revertedWith(
      "settlement not due"
    );
  });

  it("catches up all accrued earnings when a day is missed", async () => {
    await startLease();
    await ethers.provider.send("evm_increaseTime", [2 * 24 * 60 * 60]);
    await ethers.provider.send("evm_mine", []);

    await lease.connect(operator).settleLease(0);

    const entitlement = await lease.settlementEntitlement(0);
    const settlement = await lease.leaseSettlements(0);
    expect(settlement.providerPaid).to.equal(entitlement.providerEarned);
    expect(settlement.feePaid).to.equal(entitlement.feeEarned);
  });

  it("settles a bounded batch", async () => {
    await startLease(provider.address);
    await startLease(other.address);
    await ethers.provider.send("evm_increaseTime", [24 * 60 * 60]);
    await ethers.provider.send("evm_mine", []);

    await lease.connect(operator).settleLeases([0, 1]);

    expect(await wallet.userBalance(provider.address)).to.be.gt(0);
    expect(await wallet.userBalance(other.address)).to.be.gt(0);
    expect((await lease.leaseSettlements(0)).providerPaid).to.be.gt(0);
    expect((await lease.leaseSettlements(1)).providerPaid).to.be.gt(0);
  });

  it("rejects unauthorized settlement callers", async () => {
    await startLease();
    await ethers.provider.send("evm_increaseTime", [24 * 60 * 60]);
    await ethers.provider.send("evm_mine", []);

    await expect(lease.connect(other).settleLease(0)).to.be.revertedWith(
      "not settlement operator"
    );
  });

  it("does not accrue compute during the paused period", async () => {
    await startLease();
    await lease.pauseLease(0);
    await ethers.provider.send("evm_increaseTime", [24 * 60 * 60]);
    await ethers.provider.send("evm_mine", []);

    await lease.connect(operator).settleLease(0);

    const settlement = await lease.leaseSettlements(0);
    const totalPaid = settlement.providerPaid + settlement.feePaid;
    const leaseData = await lease.leases(0);
    const block = await ethers.provider.getBlock("latest");
    if (!block) throw new Error("Cannot fetch latest block");
    const elapsed = Math.min(
      block.timestamp - Number(leaseData.startTime),
      Number(leaseData.duration)
    );
    const pausedDuration =
      Number(leaseData.pausedDuration) +
      (block.timestamp - Number(leaseData.pausedAt));
    const activeDuration = elapsed - Math.min(elapsed, pausedDuration);
    const storageCost = BigInt(elapsed) * storagePrice;
    const computeCost = BigInt(activeDuration) * computePrice;

    expect(totalPaid).to.equal(storageCost + computeCost);
    expect(computeCost).to.be.lt(10n * 60n * computePrice);
  });

  it("completes after a daily settlement without paying twice", async () => {
    await startLease();
    await ethers.provider.send("evm_increaseTime", [24 * 60 * 60]);
    await ethers.provider.send("evm_mine", []);
    await lease.connect(operator).settleLease(0);

    await ethers.provider.send("evm_increaseTime", [6 * 60 * 60]);
    await ethers.provider.send("evm_mine", []);
    await lease.completeLease(0);

    const entitlement = await lease.settlementEntitlement(0);
    expect(await wallet.userBalance(provider.address)).to.equal(
      entitlement.providerEarned
    );
    expect(await wallet.userBalance(treasury.address)).to.equal(
      entitlement.feeEarned
    );
    expect(await lease.frozenFunds(0)).to.equal(0);
    expect((await lease.leases(0)).completed).to.equal(true);
  });

  it("uses cumulative referral accounting across daily settlements", async () => {
    const Referral = await ethers.getContractFactory("GPULeaseReferral");
    const referral = await Referral.deploy();
    await referral.setReferrer(user.address, owner.address);
    await referral.setReferralShareBps(3333);
    await lease.setReferralManager(referral.target);
    await startLease();

    for (let day = 0; day < 2; day++) {
      await ethers.provider.send("evm_increaseTime", [24 * 60 * 60]);
      await ethers.provider.send("evm_mine", []);
      await lease.connect(operator).settleLease(0);
    }

    const settlement = await lease.leaseSettlements(0);
    expect(await wallet.userBalance(owner.address)).to.equal(
      settlement.referralPaid
    );
    expect(await wallet.userBalance(treasury.address)).to.equal(
      settlement.feePaid - settlement.referralPaid
    );
    expect(await wallet.userBalance(provider.address)).to.equal(
      settlement.providerPaid
    );
  });

  it("allows the owner to settle even when another operator is configured", async () => {
    await startLease();
    await ethers.provider.send("evm_increaseTime", [24 * 60 * 60]);
    await ethers.provider.send("evm_mine", []);

    await expect(lease.connect(owner).settleLease(0)).to.emit(
      lease,
      "LeaseSettled"
    );
  });

  it("updates the operator and rejects zero address and non-owner updates", async () => {
    await expect(lease.setSettlementOperator(other.address))
      .to.emit(lease, "SettlementOperatorUpdated")
      .withArgs(operator.address, other.address);
    expect(await lease.settlementOperator()).to.equal(other.address);

    await expect(
      lease.connect(user).setSettlementOperator(operator.address)
    ).to.be.revertedWithCustomError(lease, "OwnableUnauthorizedAccount");
    await expect(
      lease.setSettlementOperator(ethers.ZeroAddress)
    ).to.be.revertedWith("zero settlement operator");
  });

  it("does not give the settlement operator lease lifecycle permissions", async () => {
    const block = await ethers.provider.getBlock("latest");
    if (!block) throw new Error("Cannot fetch latest block");
    await expect(
      lease.connect(operator).startLease(
        block.timestamp,
        duration,
        storagePrice,
        computePrice,
        provider.address,
        user.address
      )
    ).to.be.revertedWithCustomError(lease, "OwnableUnauthorizedAccount");
  });

  it("becomes due exactly at the 24 hour boundary", async () => {
    await startLease();
    const settlement = await lease.leaseSettlements(0);
    const dueAt = Number(settlement.lastSettledAt) + 24 * 60 * 60;

    await ethers.provider.send("evm_setNextBlockTimestamp", [dueAt - 1]);
    await ethers.provider.send("evm_mine", []);
    expect(await lease.isSettlementDue(0)).to.equal(false);

    await ethers.provider.send("evm_setNextBlockTimestamp", [dueAt]);
    await ethers.provider.send("evm_mine", []);
    expect(await lease.isSettlementDue(0)).to.equal(true);
  });

  it("allows the final short-period settlement when the lease has ended", async () => {
    const shortDuration = 2 * 60 * 60;
    const block = await ethers.provider.getBlock("latest");
    if (!block) throw new Error("Cannot fetch latest block");
    await lease.startLease(
      block.timestamp,
      shortDuration,
      storagePrice,
      computePrice,
      provider.address,
      user.address
    );
    await ethers.provider.send("evm_increaseTime", [shortDuration]);
    await ethers.provider.send("evm_mine", []);

    expect(await lease.isSettlementDue(0)).to.equal(true);
    await lease.connect(operator).settleLease(0);
    expect((await lease.leaseSettlements(0)).providerPaid).to.be.gt(0);
  });

  it("returns false for unknown and completed leases", async () => {
    expect(await lease.isSettlementDue(999)).to.equal(false);
    await expect(lease.settlementEntitlement(999)).to.be.revertedWith(
      "Lease not started"
    );

    await startLease();
    await lease.completeLease(0);
    expect(await lease.isSettlementDue(0)).to.equal(false);
    await expect(lease.connect(operator).settleLease(0)).to.be.revertedWith(
      "settlement not due"
    );
  });

  it("rejects empty and oversized batches", async () => {
    await expect(
      lease.connect(operator).settleLeases([])
    ).to.be.revertedWith("empty batch");

    await expect(
      lease.connect(operator).settleLeases(Array(51).fill(0))
    ).to.be.revertedWith("batch too large");
  });

  it("rejects batch calls from unauthorized accounts", async () => {
    await expect(lease.connect(other).settleLeases([0])).to.be.revertedWith(
      "not settlement operator"
    );
  });

  it("reverts the entire batch if one lease is not due", async () => {
    await startLease(provider.address);
    await startLease(other.address);
    await ethers.provider.send("evm_increaseTime", [24 * 60 * 60]);
    await ethers.provider.send("evm_mine", []);

    await lease.connect(operator).settleLease(0);
    await expect(
      lease.connect(operator).settleLeases([0, 1])
    ).to.be.revertedWith("settlement not due");

    expect((await lease.leaseSettlements(1)).providerPaid).to.equal(0);
    expect(await wallet.userBalance(other.address)).to.equal(0);
  });

  it("decreases frozen funds by exactly the distributed amount", async () => {
    await startLease();
    const frozenBefore = await lease.frozenFunds(0);
    await ethers.provider.send("evm_increaseTime", [24 * 60 * 60]);
    await ethers.provider.send("evm_mine", []);

    await lease.connect(operator).settleLease(0);
    const settlement = await lease.leaseSettlements(0);
    const frozenAfter = await lease.frozenFunds(0);

    expect(frozenBefore - frozenAfter).to.equal(
      settlement.providerPaid + settlement.feePaid
    );
  });

  it("emits settlement deltas that match credited balances", async () => {
    await startLease();
    await ethers.provider.send("evm_increaseTime", [24 * 60 * 60]);
    await ethers.provider.send("evm_mine", []);

    const tx = await lease.connect(operator).settleLease(0);
    const receipt = await tx.wait();
    const event = receipt.logs
      .map((log: any) => {
        try {
          return lease.interface.parseLog(log);
        } catch {
          return null;
        }
      })
      .find((log: any) => log?.name === "LeaseSettled");

    expect(event).not.to.equal(undefined);
    expect(event.args.leaseId).to.equal(0);
    expect(event.args.providerAmount).to.equal(
      await wallet.userBalance(provider.address)
    );
    expect(event.args.platformFee).to.equal(
      await wallet.userBalance(treasury.address)
    );
    expect(event.args.referralAmount).to.equal(0);
    expect(event.args.settledAt).to.equal(
      (await lease.leaseSettlements(0)).lastSettledAt
    );
  });

  it("credits only the new delta on the second daily settlement", async () => {
    await startLease();
    await ethers.provider.send("evm_increaseTime", [24 * 60 * 60]);
    await ethers.provider.send("evm_mine", []);
    await lease.connect(operator).settleLease(0);
    const firstPaid = (await lease.leaseSettlements(0)).providerPaid;

    await ethers.provider.send("evm_increaseTime", [24 * 60 * 60]);
    await ethers.provider.send("evm_mine", []);
    await lease.connect(operator).settleLease(0);
    const secondSettlement = await lease.leaseSettlements(0);

    expect(secondSettlement.providerPaid).to.be.gt(firstPaid);
    expect(await wallet.userBalance(provider.address)).to.equal(
      secondSettlement.providerPaid
    );
  });

  it("uses the custom user fee captured when the lease starts", async () => {
    await lease.setUserFee(user.address, 25);
    await startLease();
    await lease.setUserFee(user.address, 1);
    await ethers.provider.send("evm_increaseTime", [24 * 60 * 60]);
    await ethers.provider.send("evm_mine", []);
    await lease.connect(operator).settleLease(0);

    const settlement = await lease.leaseSettlements(0);
    const distributed = settlement.providerPaid + settlement.feePaid;
    expect(settlement.feePaid).to.equal((distributed * 25n) / 100n);
  });

  it("handles a 100 percent platform fee", async () => {
    await lease.setUserFee(user.address, 100);
    await startLease();
    await ethers.provider.send("evm_increaseTime", [24 * 60 * 60]);
    await ethers.provider.send("evm_mine", []);

    expect(await lease.isSettlementDue(0)).to.equal(true);
    await lease.connect(operator).settleLease(0);

    const settlement = await lease.leaseSettlements(0);
    expect(settlement.providerPaid).to.equal(0);
    expect(settlement.feePaid).to.be.gt(0);
    expect(await wallet.userBalance(provider.address)).to.equal(0);
    expect(await wallet.userBalance(treasury.address)).to.equal(
      settlement.feePaid
    );
  });

  it("settles paused and resumed portions cumulatively", async () => {
    await startLease();
    await lease.pauseLease(0);
    await ethers.provider.send("evm_increaseTime", [24 * 60 * 60]);
    await ethers.provider.send("evm_mine", []);
    await lease.connect(operator).settleLease(0);
    const afterPausedDay = (await lease.leaseSettlements(0)).providerPaid;

    await lease.resumeLease(0);
    await ethers.provider.send("evm_increaseTime", [24 * 60 * 60]);
    await ethers.provider.send("evm_mine", []);
    await lease.connect(operator).settleLease(0);
    const afterActiveDay = (await lease.leaseSettlements(0)).providerPaid;

    expect(afterActiveDay - afterPausedDay).to.be.gt(afterPausedDay);
  });

  it("performs an immediate final settlement before one day has passed", async () => {
    await startLease();
    await ethers.provider.send("evm_increaseTime", [60 * 60]);
    await ethers.provider.send("evm_mine", []);

    expect(await lease.isSettlementDue(0)).to.equal(false);
    await lease.completeLease(0);

    expect((await lease.leaseSettlements(0)).providerPaid).to.be.gt(0);
    expect(await wallet.userBalance(provider.address)).to.be.gt(0);
    expect(await lease.frozenFunds(0)).to.equal(0);
  });
});

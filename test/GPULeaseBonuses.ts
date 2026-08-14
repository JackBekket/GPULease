import { expect } from "chai";
import { network } from "hardhat";

const { ethers } = await network.connect();

describe("GPULease bonus balances", function () {
  let owner: any;
  let user: any;
  let provider: any;
  let treasury: any;
  let other: any;
  let token: any;
  let wallet: any;
  let lease: any;

  const unit = 1_000_000n;

  beforeEach(async () => {
    [owner, user, provider, treasury, other] = await ethers.getSigners();

    const Token = await ethers.getContractFactory("MockERC20");
    token = await Token.deploy("Mock USDC", "USDC");

    const Wallet = await ethers.getContractFactory("GPULeaseWallet");
    wallet = await Wallet.deploy(token.target);

    const Lease = await ethers.getContractFactory("GPULease");
    lease = await Lease.deploy(wallet.target, treasury.address);
    await wallet.setLeaseManager(lease.target);
    await lease.setPlatformFee(0);

    await token.mint(owner.address, 10_000n * unit);
    await token.mint(user.address, 10_000n * unit);
    await token.approve(wallet.target, ethers.MaxUint256);
    await token.connect(user).approve(wallet.target, ethers.MaxUint256);
  });

  async function fundAndGrant(amount: bigint) {
    await wallet.fundBonusPool(amount);
    await wallet.grantBonus(user.address, amount);
  }

  async function startLeaseNow(
    duration: number,
    storagePrice: bigint,
    computePrice = 0n,
  ) {
    const block = await ethers.provider.getBlock("latest");
    if (!block) throw new Error("Cannot fetch latest block");
    const startTimestamp = block.timestamp + 1;
    await ethers.provider.send("evm_setNextBlockTimestamp", [startTimestamp]);
    return lease.startLease(
      startTimestamp,
      duration,
      storagePrice,
      computePrice,
      provider.address,
      user.address,
    );
  }

  it("funds a real USDC reserve before issuing bonuses", async () => {
    const amount = 100n * unit;

    await expect(wallet.fundBonusPool(amount))
      .to.emit(wallet, "BonusPoolFunded")
      .withArgs(owner.address, amount);

    expect(await wallet.bonusReserve()).to.equal(amount);
    expect(await token.balanceOf(wallet.target)).to.equal(amount);
  });

  it("grants non-withdrawable bonuses from the funded reserve", async () => {
    const amount = 100n * unit;
    await wallet.fundBonusPool(amount);

    await expect(wallet.grantBonus(user.address, amount))
      .to.emit(wallet, "BonusGranted")
      .withArgs(user.address, amount);

    expect(await wallet.bonusReserve()).to.equal(0);
    expect(await wallet.bonusBalance(user.address)).to.equal(amount);
    expect(await wallet.withdrawableBalance(user.address)).to.equal(0);
    expect(await wallet.spendableBalance(user.address)).to.equal(amount);
    expect(await wallet.userBalance(user.address)).to.equal(amount);
    await expect(wallet.connect(user).withdraw(1)).to.be.revertedWith(
      "insufficient balance",
    );
  });

  it("keeps deposited cash withdrawable and separate from bonuses", async () => {
    await wallet.connect(user).deposit(40n * unit);
    await fundAndGrant(60n * unit);

    expect(await wallet.withdrawableBalance(user.address)).to.equal(40n * unit);
    expect(await wallet.bonusBalance(user.address)).to.equal(60n * unit);
    expect(await wallet.spendableBalance(user.address)).to.equal(100n * unit);

    await wallet.connect(user).withdraw(40n * unit);
    expect(await wallet.withdrawableBalance(user.address)).to.equal(0);
    expect(await wallet.bonusBalance(user.address)).to.equal(60n * unit);
  });

  it("spends bonuses before cash when starting a lease", async () => {
    await wallet.connect(user).deposit(50n * unit);
    await fundAndGrant(60n * unit);

    await startLeaseNow(100, unit);

    expect(await lease.frozenFunds(0)).to.equal(100n * unit);
    expect(await lease.frozenBonusFunds(0)).to.equal(60n * unit);
    expect(await lease.frozenCashFunds(0)).to.equal(40n * unit);
    expect(await wallet.bonusBalance(user.address)).to.equal(0);
    expect(await wallet.withdrawableBalance(user.address)).to.equal(10n * unit);
  });

  it("allows a lease funded entirely by bonuses", async () => {
    await fundAndGrant(100n * unit);

    await startLeaseNow(100, unit);

    expect(await lease.frozenBonusFunds(0)).to.equal(100n * unit);
    expect(await lease.frozenCashFunds(0)).to.equal(0);
  });

  it("turns spent bonuses into withdrawable provider earnings", async () => {
    await fundAndGrant(100n * unit);
    await startLeaseNow(100, unit);
    await ethers.provider.send("evm_increaseTime", [100]);
    await ethers.provider.send("evm_mine", []);

    await lease.completeLease(0);

    const providerBalance = await wallet.withdrawableBalance(provider.address);
    expect(providerBalance).to.equal(100n * unit);
    expect(await wallet.bonusBalance(provider.address)).to.equal(0);

    const before = await token.balanceOf(provider.address);
    await wallet.connect(provider).withdraw(providerBalance);
    expect(await token.balanceOf(provider.address)).to.equal(
      before + providerBalance,
    );
  });

  it("returns unused cash and bonuses to their original balance types", async () => {
    await wallet.connect(user).deposit(50n * unit);
    await fundAndGrant(60n * unit);
    await startLeaseNow(100, unit);

    await ethers.provider.send("evm_increaseTime", [25]);
    await ethers.provider.send("evm_mine", []);
    await lease.completeLease(0);

    const paid = (await lease.leaseSettlements(0)).providerPaid;
    expect(await wallet.bonusBalance(user.address)).to.equal(60n * unit - paid);
    expect(await wallet.withdrawableBalance(user.address)).to.equal(50n * unit);
    expect(await wallet.withdrawableBalance(provider.address)).to.equal(paid);
    expect(await lease.frozenFunds(0)).to.equal(0);
    expect(await lease.frozenBonusFunds(0)).to.equal(0);
    expect(await lease.frozenCashFunds(0)).to.equal(0);
  });

  it("uses bonus funds across daily settlements before touching frozen cash", async () => {
    const duration = 3 * 24 * 60 * 60;
    const price = 1_000n;
    const dailyAmount = 24n * 60n * 60n * price;
    await wallet.connect(user).deposit(2n * dailyAmount);
    await fundAndGrant(dailyAmount);
    await startLeaseNow(duration, price);

    await ethers.provider.send("evm_increaseTime", [24 * 60 * 60]);
    await ethers.provider.send("evm_mine", []);
    await lease.settleLease(0);

    const paid = (await lease.leaseSettlements(0)).providerPaid;
    expect(await lease.frozenBonusFunds(0)).to.equal(0);
    expect(await lease.frozenCashFunds(0)).to.equal(
      2n * dailyAmount - (paid - dailyAmount),
    );
    expect(await wallet.withdrawableBalance(provider.address)).to.equal(paid);
  });

  it("supports batch grants with reserve and input validation", async () => {
    await wallet.fundBonusPool(100n * unit);

    await wallet.grantBonuses(
      [user.address, other.address],
      [60n * unit, 40n * unit],
    );

    expect(await wallet.bonusBalance(user.address)).to.equal(60n * unit);
    expect(await wallet.bonusBalance(other.address)).to.equal(40n * unit);
    expect(await wallet.bonusReserve()).to.equal(0);

    await expect(wallet.grantBonuses([], [])).to.be.revertedWith("empty batch");
    await expect(
      wallet.grantBonuses([user.address], [1, 2]),
    ).to.be.revertedWith("length mismatch");
    await expect(
      wallet.grantBonuses(Array(101).fill(user.address), Array(101).fill(1)),
    ).to.be.revertedWith("batch too large");
  });

  it("rejects unauthorized, unbacked, zero and invalid bonus grants", async () => {
    await expect(
      wallet.connect(user).grantBonus(user.address, unit),
    ).to.be.revertedWithCustomError(wallet, "OwnableUnauthorizedAccount");
    await expect(wallet.grantBonus(user.address, unit)).to.be.revertedWith(
      "insufficient bonus reserve",
    );

    await wallet.fundBonusPool(10n * unit);
    await expect(wallet.grantBonus(ethers.ZeroAddress, unit)).to.be.revertedWith(
      "zero user",
    );
    await expect(wallet.grantBonus(user.address, 0)).to.be.revertedWith(
      "zero amount",
    );
  });

  it("revokes only unused bonuses and allows withdrawing unallocated reserve", async () => {
    await fundAndGrant(100n * unit);

    await expect(wallet.revokeBonus(user.address, 40n * unit))
      .to.emit(wallet, "BonusRevoked")
      .withArgs(user.address, 40n * unit);

    expect(await wallet.bonusBalance(user.address)).to.equal(60n * unit);
    expect(await wallet.bonusReserve()).to.equal(40n * unit);

    const before = await token.balanceOf(other.address);
    await wallet.withdrawBonusReserve(other.address, 40n * unit);
    expect(await token.balanceOf(other.address)).to.equal(before + 40n * unit);
    expect(await wallet.bonusReserve()).to.equal(0);
  });

  it("guards lease-only debit and refund functions", async () => {
    await expect(
      wallet.connect(user).debitForLease(user.address, 1),
    ).to.be.revertedWith("not lease manager");
    await expect(
      wallet.connect(user).refundLeaseBalance(user.address, 1, 1),
    ).to.be.revertedWith("not lease manager");
  });
});

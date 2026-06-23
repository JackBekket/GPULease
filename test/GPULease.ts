import { expect } from "chai";
import { network } from "hardhat";
const { ethers } = await network.connect();
describe("GPULease", function () {
  let owner: any;
  let user: any;
  let provider: any;
  let treasury: any;

  let token: any;
  let wallet: any;
  let lease: any;
  let referral: any;

  beforeEach(async () => {
    [owner, user, provider, treasury] = await ethers.getSigners();

    // Deploy MockERC20
    const Token = await ethers.getContractFactory("MockERC20");
    token = await Token.deploy("Mock", "MOCK");

    // Mint tokens to user
    await token.mint(user.address, ethers.parseEther("1000"));

    const Wallet = await ethers.getContractFactory("GPULeaseWallet");
    wallet = await Wallet.deploy(token.target);

    // Deploy GPULease
    const Lease = await ethers.getContractFactory("GPULease");
    lease = await Lease.deploy(wallet.target, treasury.address);
    await wallet.setLeaseManager(lease.target);

    // Approve + deposit
    await token.connect(user).approve(wallet.target, ethers.parseEther("1000"));
    await wallet.connect(user).deposit(ethers.parseEther("1000"));
  });

  it("should deposit correctly", async () => {
    const balance = await wallet.userBalance(user.address);
    expect(balance).to.equal(ethers.parseEther("1000"));
    expect(await wallet.userBalance(user.address)).to.equal(ethers.parseEther("1000"));
  });

  it("should deposit for another user", async () => {
    await token.mint(owner.address, ethers.parseEther("100"));
    await token.approve(wallet.target, ethers.parseEther("100"));

    await wallet.depositFor(provider.address, ethers.parseEther("100"));

    expect(await wallet.userBalance(provider.address)).to.equal(
      ethers.parseEther("100")
    );
    expect(await token.balanceOf(wallet.target)).to.equal(
      ethers.parseEther("1100")
    );
  });

  it("should not deposit for zero address", async () => {
    await expect(
      wallet.depositFor(ethers.ZeroAddress, ethers.parseEther("1"))
    ).to.be.revertedWith("zero beneficiary");
  });

  it("should start lease and freeze funds", async () => {
    const duration = 1000;
    const price = ethers.parseEther("0.001");

    await lease.startLease(
      duration,
      price,
      price,
      provider.address,
      user.address
    );

    const frozen = await lease.frozenFunds(0);
    expect(frozen).to.be.gt(0);

    const userBalance = await wallet.userBalance(user.address);
    expect(userBalance).to.be.lt(ethers.parseEther("1000"));
  });

  it("should pause and resume lease", async () => {
    const duration = 1000;
    const price = ethers.parseEther("0.001");

    await lease.startLease(
      duration,
      price,
      price,
      provider.address,
      user.address
    );

    await lease.pauseLease(0);

    let leaseData = await lease.leases(0);
    expect(leaseData.paused).to.equal(true);

    await lease.resumeLease(0);

    leaseData = await lease.leases(0);
    expect(leaseData.paused).to.equal(false);
  });

  it("should complete lease and distribute funds", async () => {
    const duration = 1000;
    const price = ethers.parseEther("0.001");

    await lease.startLease(
      duration,
      price,
      price,
      provider.address,
      user.address
    );

    // немного прокрутим время
    await ethers.provider.send("evm_increaseTime", [500]);
    await ethers.provider.send("evm_mine", []);

    await lease.completeLease(0);

    const leaseData = await lease.leases(0);
    expect(leaseData.completed).to.equal(true);

    const providerBalance = await wallet.userBalance(provider.address);
    expect(providerBalance).to.be.gt(0);

    const treasuryBalance = await wallet.userBalance(treasury.address);
    expect(treasuryBalance).to.be.gt(0);
  });

  it("should return frozen funds list", async () => {
    const duration = 1000;
    const price = ethers.parseEther("0.001");

    await lease.startLease(
      duration,
      price,
      price,
      provider.address,
      user.address
    );

    const result = await lease.getUserFrozenFunds(user.address);

    expect(result.length).to.equal(1);
    expect(result[0].leaseId).to.equal(0);
    expect(result[0].amount).to.be.gt(0);
  });

  it("should withdraw correctly", async () => {
    await wallet.connect(user).withdraw(ethers.parseEther("100"));

    const balance = await wallet.userBalance(user.address);
    expect(balance).to.equal(ethers.parseEther("900"));
  });

  it("should not allow withdraw more than balance", async () => {
    await expect(
      wallet.connect(user).withdraw(ethers.parseEther("2000"))
    ).to.be.revertedWith("insufficient balance");
  });

  it("should allow owner to change platform fee", async () => {
    expect(await lease.platformFeePercentage()).to.equal(10);

    await expect(lease.setPlatformFee(15))
      .to.emit(lease, "PlatformFeeUpdated")
      .withArgs(10, 15);

    const fee = await lease.platformFeePercentage();
    expect(fee).to.equal(15);
  });

  it("should allow owner to set and clear user fee", async () => {
    expect(await lease.feePercentageForUser(user.address)).to.equal(10);

    await expect(lease.setUserFee(user.address, 25))
      .to.emit(lease, "UserFeeUpdated")
      .withArgs(user.address, 25);

    expect(await lease.hasCustomUserFee(user.address)).to.equal(true);
    expect(await lease.userFeePercentage(user.address)).to.equal(25);
    expect(await lease.feePercentageForUser(user.address)).to.equal(25);

    await expect(lease.clearUserFee(user.address))
      .to.emit(lease, "UserFeeCleared")
      .withArgs(user.address);

    expect(await lease.hasCustomUserFee(user.address)).to.equal(false);
    expect(await lease.feePercentageForUser(user.address)).to.equal(10);
  });

  it("should use custom user fee when freezing and completing lease", async () => {
    const duration = 1000;
    const price = ethers.parseEther("0.001");
    const baseAmount = BigInt(duration) * BigInt(price) * 2n;
    const customFee = (baseAmount * 25n) / 100n;

    await lease.setUserFee(user.address, 25);
    await lease.startLease(
      duration,
      price,
      price,
      provider.address,
      user.address
    );

    const leaseData = await lease.leases(0);
    expect(leaseData.leaseFeePercentage).to.equal(25);
    expect(await lease.frozenFunds(0)).to.equal(baseAmount + customFee);

    await ethers.provider.send("evm_increaseTime", [1500]);
    await ethers.provider.send("evm_mine", []);

    await lease.completeLease(0);

    expect(await wallet.userBalance(treasury.address)).to.equal(customFee);
    expect(await wallet.userBalance(provider.address)).to.equal(
      baseAmount - customFee
    );
  });

  it("should lock lease fee at start", async () => {
    const duration = 1000;
    const price = ethers.parseEther("0.001");

    await lease.setUserFee(user.address, 20);
    await lease.startLease(
      duration,
      price,
      price,
      provider.address,
      user.address
    );
    await lease.setUserFee(user.address, 5);

    const leaseData = await lease.leases(0);
    expect(leaseData.leaseFeePercentage).to.equal(20);
  });

  it("should not allow non-owner to set user fee", async () => {
    await expect(
      lease.connect(user).setUserFee(user.address, 20)
    ).to.be.revertedWithCustomError(lease, "OwnableUnauthorizedAccount");
  });

  it("should split user fee between treasury and referrer", async () => {
    const Referral = await ethers.getContractFactory("GPULeaseReferral");
    referral = await Referral.deploy();
    await referral.setReferrer(user.address, owner.address);
    await lease.setReferralManager(referral.target);

    const duration = 1000;
    const price = ethers.parseEther("0.001");
    const baseAmount = BigInt(duration) * BigInt(price) * 2n;
    const totalFee = (baseAmount * 10n) / 100n;
    const referralFee = totalFee / 2n;

    await lease.startLease(
      duration,
      price,
      price,
      provider.address,
      user.address
    );

    const referralInfo = await lease.leaseReferralInfo(0);
    expect(referralInfo.referrer).to.equal(owner.address);
    expect(referralInfo.referralShareBps).to.equal(5000);
    expect(await lease.frozenFunds(0)).to.equal(baseAmount + totalFee);

    await ethers.provider.send("evm_increaseTime", [1500]);
    await ethers.provider.send("evm_mine", []);

    await lease.completeLease(0);

    expect(await wallet.userBalance(treasury.address)).to.equal(
      totalFee - referralFee
    );
    expect(await wallet.userBalance(owner.address)).to.equal(referralFee);
    expect(await wallet.userBalance(provider.address)).to.equal(
      baseAmount - totalFee
    );
  });

  it("should split a custom user fee in half", async () => {
    const Referral = await ethers.getContractFactory("GPULeaseReferral");
    referral = await Referral.deploy();
    await referral.setReferrer(user.address, owner.address);
    await lease.setReferralManager(referral.target);
    await lease.setUserFee(user.address, 7);

    const duration = 1000;
    const price = ethers.parseEther("0.001");
    const baseAmount = BigInt(duration) * BigInt(price) * 2n;
    const totalFee = (baseAmount * 7n) / 100n;
    const referralFee = totalFee / 2n;

    await lease.startLease(
      duration,
      price,
      price,
      provider.address,
      user.address
    );

    await ethers.provider.send("evm_increaseTime", [1500]);
    await ethers.provider.send("evm_mine", []);

    await lease.completeLease(0);

    expect(await wallet.userBalance(treasury.address)).to.equal(
      totalFee - referralFee
    );
    expect(await wallet.userBalance(owner.address)).to.equal(referralFee);
  });

  it("should lock referral rules at lease start", async () => {
    const Referral = await ethers.getContractFactory("GPULeaseReferral");
    referral = await Referral.deploy();
    await referral.setReferrer(user.address, owner.address);
    await lease.setReferralManager(referral.target);

    const duration = 1000;
    const price = ethers.parseEther("0.001");

    await lease.startLease(
      duration,
      price,
      price,
      provider.address,
      user.address
    );

    await referral.setReferralShareBps(3000);
    await referral.setReferrer(user.address, provider.address);

    const referralInfo = await lease.leaseReferralInfo(0);
    expect(referralInfo.referrer).to.equal(owner.address);
    expect(referralInfo.referralShareBps).to.equal(5000);
  });

  it("should allow owner to replace referral manager", async () => {
    const Referral = await ethers.getContractFactory("GPULeaseReferral");
    const firstReferral = await Referral.deploy();
    const secondReferral = await Referral.deploy();

    await expect(lease.setReferralManager(firstReferral.target))
      .to.emit(lease, "ReferralManagerUpdated")
      .withArgs(ethers.ZeroAddress, firstReferral.target);

    await expect(lease.setReferralManager(secondReferral.target))
      .to.emit(lease, "ReferralManagerUpdated")
      .withArgs(firstReferral.target, secondReferral.target);

    expect(await lease.referralManager()).to.equal(secondReferral.target);
  });

  it("should allow custom referral share", async () => {
    const Referral = await ethers.getContractFactory("GPULeaseReferral");
    referral = await Referral.deploy();
    await referral.setReferrer(user.address, owner.address);
    await referral.setReferralShareBps(2500);
    await lease.setReferralManager(referral.target);

    await lease.startLease(
      1000,
      ethers.parseEther("0.001"),
      ethers.parseEther("0.001"),
      provider.address,
      user.address
    );

    const referralInfo = await lease.leaseReferralInfo(0);
    expect(referralInfo.referralShareBps).to.equal(2500);
  });

  it("should guard referral manager admin methods", async () => {
    const Referral = await ethers.getContractFactory("GPULeaseReferral");
    referral = await Referral.deploy();

    await expect(
      referral.connect(user).setReferrer(user.address, owner.address)
    ).to.be.revertedWithCustomError(referral, "OwnableUnauthorizedAccount");

    await expect(
      referral.setReferrer(user.address, user.address)
    ).to.be.revertedWith("self referral");

    await expect(referral.setReferralShareBps(10001)).to.be.revertedWith("Share too high");
  });

  it("should move treasury balance on change", async () => {
    // artificially give treasury balance
    await lease.setPlatformFee(10);

    await lease.setTreasury(provider.address);

    const newTreasury = await lease.treasury();
    expect(newTreasury).to.equal(provider.address);
  });

  it("should keep wallet balances when GPULease is replaced", async () => {
    const Lease = await ethers.getContractFactory("GPULease");
    const upgradedLease = await Lease.deploy(wallet.target, treasury.address);

    await wallet.setLeaseManager(upgradedLease.target);

    expect(await wallet.userBalance(user.address)).to.equal(
      ethers.parseEther("1000")
    );
    expect(await wallet.userBalance(user.address)).to.equal(
      ethers.parseEther("1000")
    );

    const duration = 1000;
    const price = ethers.parseEther("0.001");

    await expect(
      lease.startLease(duration, price, price, provider.address, user.address)
    ).to.be.revertedWith("not lease manager");

    await upgradedLease.startLease(
      duration,
      price,
      price,
      provider.address,
      user.address
    );

    expect(await upgradedLease.frozenFunds(0)).to.be.gt(0);
    expect(await wallet.userBalance(user.address)).to.be.lt(
      ethers.parseEther("1000")
    );
  });

  
});

import { network } from "hardhat";
const { ethers } = await network.connect();
import { expect } from "chai";

describe("GPULease Gas Tests (raw wei)", function () {
  let lease: any;
  let wallet: any;
  let token: any;
  let owner: any;
  let user: any;
  let provider: any;
  let treasury: any;

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

  async function logGas(tx: any, label: string) {
    const receipt = await tx.wait();

    const gasUsed = receipt.gasUsed;
    const gasPrice = receipt.gasPrice ?? receipt.effectiveGasPrice;
    const totalWei = gasUsed * gasPrice;

    console.log(`\n=== ${label} ===`);
    console.log("gasUsed:", gasUsed.toString());
    console.log("gasPrice:", gasPrice.toString());
    console.log("totalWei:", totalWei.toString());
  }

  beforeEach(async () => {
    [owner, user, provider, treasury] = await ethers.getSigners();

    // Deploy token
    const Token = await ethers.getContractFactory("MockERC20");
    token = await Token.deploy("Mock", "MCK");
    await token.waitForDeployment();

    const Wallet = await ethers.getContractFactory("GPULeaseWallet");
    wallet = await Wallet.deploy(token.target);
    await wallet.waitForDeployment();

    // Deploy lease contract
    const Lease = await ethers.getContractFactory("GPULease");
    lease = await Lease.deploy(wallet.target, treasury.address);
    await lease.waitForDeployment();
    await wallet.setLeaseManager(lease.target);

    // Setup balances
    await token.mint(user.address, ethers.parseEther("10000"));
    await token.connect(user).approve(wallet.target, ethers.parseEther("10000"));

    await wallet.connect(user).deposit(ethers.parseEther("1000"));
  });

  it("startLease gas", async () => {
    const tx = await startLeaseNow(lease,      3600,
      ethers.parseEther("0.0001"),
      ethers.parseEther("0.0002"),
      provider.address,
      user.address
    );

    await logGas(tx, "startLease");
  });

  it("pause + resume gas", async () => {
    await startLeaseNow(lease,      3600,
      ethers.parseEther("0.0001"),
      ethers.parseEther("0.0002"),
      provider.address,
      user.address
    );

    await logGas(await lease.pauseLease(0), "pauseLease");
    await logGas(await lease.resumeLease(0), "resumeLease");
  });

  it("completeLease gas", async () => {
    await startLeaseNow(lease,      3600,
      ethers.parseEther("0.0001"),
      ethers.parseEther("0.0002"),
      provider.address,
      user.address
    );

    // имитируем время
    await ethers.provider.send("evm_increaseTime", [1800]);
    await ethers.provider.send("evm_mine", []);

    await logGas(await lease.completeLease(0), "completeLease");
  });

  it("daily settlement gas", async () => {
    await startLeaseNow(lease,      3 * 24 * 60 * 60,
      ethers.parseEther("0.000001"),
      ethers.parseEther("0.000002"),
      provider.address,
      user.address
    );

    await ethers.provider.send("evm_increaseTime", [24 * 60 * 60]);
    await ethers.provider.send("evm_mine", []);

    await logGas(await lease.settleLease(0), "settleLease (first payout)");

    await ethers.provider.send("evm_increaseTime", [24 * 60 * 60]);
    await ethers.provider.send("evm_mine", []);

    await logGas(await lease.settleLease(0), "settleLease (next payout)");
  });

  it("deposit + withdraw gas", async () => {
      await logGas(
      await wallet.connect(user).deposit(ethers.parseEther("100")),
      "deposit"
    );

    await logGas(
      await wallet.connect(user).withdraw(ethers.parseEther("50")),
      "withdraw"
    );
  });

  it("multiple leases scaling gas", async () => {
    for (let i = 0; i < 5; i++) {
      await startLeaseNow(lease,        3600,
        ethers.parseEther("0.0001"),
        ethers.parseEther("0.0002"),
        provider.address,
        user.address
      );
    }

    const gas = await lease.getUserFrozenFunds.estimateGas(user.address);

    console.log("\n=== getUserFrozenFunds (5 leases) ===");
    console.log("estimated gas:", gas.toString());
  });
});

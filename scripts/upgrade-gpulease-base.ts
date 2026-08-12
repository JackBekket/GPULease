import { network } from "hardhat";

const { ethers } = await network.connect();

const OLD_GPULEASE_ADDRESS = "0xB6E47Eb260160BD6A18A246CC0b27D9240706401";
const GPULEASE_WALLET_ADDRESS = "0xD4352D14Ba7928f6066dd7ec6031C7c0CCF13340";
const GPULEASE_REFERRAL_ADDRESS = "0x2695d98bF8b233539f5a1Fb823298AA055f2a143";
const TREASURY_ADDRESS = "0x80CFfF5C710E3e21050AaB272201B574a53A66B3";

async function main() {
  const [deployer] = await ethers.getSigners();
  const provider = ethers.provider;

  const wallet = await ethers.getContractAt(
    "GPULeaseWallet",
    GPULEASE_WALLET_ADDRESS,
  );
  const oldLease = await ethers.getContractAt(
    "GPULease",
    OLD_GPULEASE_ADDRESS,
  );

  const [walletOwner, oldLeaseOwner, currentManager, leaseCount, balance] =
    await Promise.all([
      wallet.owner(),
      oldLease.owner(),
      wallet.leaseManager(),
      oldLease.leaseCount(),
      provider.getBalance(deployer.address),
    ]);

  if (walletOwner.toLowerCase() !== deployer.address.toLowerCase()) {
    throw new Error(`deployer is not wallet owner: ${walletOwner}`);
  }
  if (oldLeaseOwner.toLowerCase() !== deployer.address.toLowerCase()) {
    throw new Error(`deployer is not old GPULease owner: ${oldLeaseOwner}`);
  }
  if (currentManager.toLowerCase() !== OLD_GPULEASE_ADDRESS.toLowerCase()) {
    throw new Error(`unexpected current lease manager: ${currentManager}`);
  }

  for (let leaseId = 0n; leaseId < leaseCount; leaseId++) {
    const leaseData = await oldLease.leases(leaseId);
    if (leaseData.active && !leaseData.completed) {
      throw new Error(`active lease prevents upgrade: ${leaseId}`);
    }
    await new Promise((resolve) => setTimeout(resolve, 250));
  }

  console.log("Deploying from:", deployer.address);
  console.log("Deployer ETH:", ethers.formatEther(balance));
  console.log("Existing wallet:", GPULEASE_WALLET_ADDRESS);
  console.log("Existing referral:", GPULEASE_REFERRAL_ADDRESS);
  console.log("Treasury:", TREASURY_ADDRESS);
  console.log("Checked leases:", leaseCount.toString(), "(0 active)");

  const GPULease = await ethers.getContractFactory("GPULease");
  const newLease = await GPULease.deploy(
    GPULEASE_WALLET_ADDRESS,
    TREASURY_ADDRESS,
  );
  await newLease.waitForDeployment();
  const newLeaseAddress = await newLease.getAddress();
  const deploymentTx = newLease.deploymentTransaction();
  const deploymentReceipt = await deploymentTx?.wait(1);

  const referralTx = await newLease.setReferralManager(
    GPULEASE_REFERRAL_ADDRESS,
  );
  await referralTx.wait(1);

  const managerTx = await wallet.setLeaseManager(newLeaseAddress);
  await managerTx.wait(1);

  const [verifiedManager, configuredReferral, configuredWallet, owner, operator] =
    await Promise.all([
      wallet.leaseManager(),
      newLease.referralManager(),
      newLease.wallet(),
      newLease.owner(),
      newLease.settlementOperator(),
    ]);

  if (verifiedManager.toLowerCase() !== newLeaseAddress.toLowerCase()) {
    throw new Error(`wallet manager verification failed: ${verifiedManager}`);
  }
  if (
    configuredReferral.toLowerCase() !== GPULEASE_REFERRAL_ADDRESS.toLowerCase()
  ) {
    throw new Error(`referral verification failed: ${configuredReferral}`);
  }
  if (configuredWallet.toLowerCase() !== GPULEASE_WALLET_ADDRESS.toLowerCase()) {
    throw new Error(`wallet verification failed: ${configuredWallet}`);
  }

  const deploymentBlock = deploymentReceipt
    ? await provider.getBlock(deploymentReceipt.blockNumber)
    : null;

  console.log("New GPULease:", newLeaseAddress);
  console.log("Deployment tx:", deploymentTx?.hash);
  console.log("Deployment block:", deploymentReceipt?.blockNumber);
  console.log(
    "Deployment date:",
    deploymentBlock
      ? new Date(deploymentBlock.timestamp * 1000).toISOString()
      : new Date().toISOString(),
  );
  console.log("Referral tx:", referralTx.hash);
  console.log("Wallet manager tx:", managerTx.hash);
  console.log("Owner:", owner);
  console.log("Settlement operator:", operator);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});

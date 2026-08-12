import { network } from "hardhat";

const { ethers } = await network.connect();

const OLD_GPULEASE_ADDRESS = "0xB6E47Eb260160BD6A18A246CC0b27D9240706401";
const NEW_GPULEASE_ADDRESS = "0xCCD732200366886e04F508D12F561ee94Eb03110";
const GPULEASE_WALLET_ADDRESS = "0xD4352D14Ba7928f6066dd7ec6031C7c0CCF13340";
const GPULEASE_REFERRAL_ADDRESS = "0x2695d98bF8b233539f5a1Fb823298AA055f2a143";

async function main() {
  const [deployer] = await ethers.getSigners();
  const wallet = await ethers.getContractAt(
    "GPULeaseWallet",
    GPULEASE_WALLET_ADDRESS,
  );
  const oldLease = await ethers.getContractAt(
    "GPULease",
    OLD_GPULEASE_ADDRESS,
  );
  const newLease = await ethers.getContractAt(
    "GPULease",
    NEW_GPULEASE_ADDRESS,
  );

  const [walletOwner, newLeaseOwner, currentManager, leaseCount] =
    await Promise.all([
      wallet.owner(),
      newLease.owner(),
      wallet.leaseManager(),
      oldLease.leaseCount(),
    ]);

  if (walletOwner.toLowerCase() !== deployer.address.toLowerCase()) {
    throw new Error(`deployer is not wallet owner: ${walletOwner}`);
  }
  if (newLeaseOwner.toLowerCase() !== deployer.address.toLowerCase()) {
    throw new Error(`deployer is not new GPULease owner: ${newLeaseOwner}`);
  }
  if (
    currentManager.toLowerCase() !== OLD_GPULEASE_ADDRESS.toLowerCase() &&
    currentManager.toLowerCase() !== NEW_GPULEASE_ADDRESS.toLowerCase()
  ) {
    throw new Error(`unexpected lease manager: ${currentManager}`);
  }

  if (currentManager.toLowerCase() === OLD_GPULEASE_ADDRESS.toLowerCase()) {
    for (let leaseId = 0n; leaseId < leaseCount; leaseId++) {
      const leaseData = await oldLease.leases(leaseId);
      if (leaseData.active && !leaseData.completed) {
        throw new Error(`active lease prevents upgrade: ${leaseId}`);
      }
      await new Promise((resolve) => setTimeout(resolve, 250));
    }
  }

  const currentReferral = await newLease.referralManager();
  if (currentReferral.toLowerCase() !== GPULEASE_REFERRAL_ADDRESS.toLowerCase()) {
    const referralTx = await newLease.setReferralManager(
      GPULEASE_REFERRAL_ADDRESS,
    );
    console.log("Referral tx:", referralTx.hash);
    await referralTx.wait(1);
  } else {
    console.log("Referral already configured");
  }

  const managerBefore = await wallet.leaseManager();
  if (managerBefore.toLowerCase() !== NEW_GPULEASE_ADDRESS.toLowerCase()) {
    const managerTx = await wallet.setLeaseManager(NEW_GPULEASE_ADDRESS);
    console.log("Wallet manager tx:", managerTx.hash);
    await managerTx.wait(1);
  } else {
    console.log("Wallet manager already configured");
  }

  const [manager, referral, configuredWallet, owner, operator, treasury] =
    await Promise.all([
      wallet.leaseManager(),
      newLease.referralManager(),
      newLease.wallet(),
      newLease.owner(),
      newLease.settlementOperator(),
      newLease.treasury(),
    ]);

  console.log("New GPULease:", NEW_GPULEASE_ADDRESS);
  console.log("Wallet manager:", manager);
  console.log("Referral:", referral);
  console.log("Wallet:", configuredWallet);
  console.log("Owner:", owner);
  console.log("Settlement operator:", operator);
  console.log("Treasury:", treasury);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});

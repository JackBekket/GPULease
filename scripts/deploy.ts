import  { network }  from "hardhat";


const { ethers } = await network.connect();

async function main() {
  // Получаем токен (например, USDC)
  const USDC = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"; // USDC на Sepolia

  const GPULeaseWallet = await ethers.getContractFactory("GPULeaseWallet");
  const wallet = await GPULeaseWallet.deploy(USDC);
  await wallet.waitForDeployment();

  const GPULeaseReferral = await ethers.getContractFactory("GPULeaseReferral");
  const referral = await GPULeaseReferral.deploy();
  await referral.waitForDeployment();

  const GPULease = await ethers.getContractFactory("GPULease");
  const treasury = await (await ethers.getSigners())[0].getAddress();
  const gpuLease = await GPULease.deploy(await wallet.getAddress(), treasury);
  await gpuLease.waitForDeployment();

  await wallet.setLeaseManager(await gpuLease.getAddress());
  await gpuLease.setReferralManager(await referral.getAddress());

  console.log("GPULeaseWallet deployed to:", await wallet.getAddress());
  console.log("GPULeaseReferral deployed to:", await referral.getAddress());
  console.log("GPULease deployed to:", await gpuLease.getAddress());
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});

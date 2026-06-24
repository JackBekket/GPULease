import  { network }  from "hardhat";


const { ethers } = await network.connect();

async function main() {
  // Получаем токен (например, USDC)
  const USDC = "0x833589fcd6edb6e08f4c7c32d4f71b54bda02913"; // USDC на Base
  const TREASURY = "0x80CFfF5C710E3e21050AaB272201B574a53A66B3";

  const GPULeaseWallet = await ethers.getContractFactory("GPULeaseWallet");
  const wallet = await GPULeaseWallet.deploy(USDC);
  await wallet.waitForDeployment();

  const GPULeaseReferral = await ethers.getContractFactory("GPULeaseReferral");
  const referral = await GPULeaseReferral.deploy();
  await referral.waitForDeployment();

  const GPULease = await ethers.getContractFactory("GPULease");
  const gpuLease = await GPULease.deploy(await wallet.getAddress(), TREASURY);
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

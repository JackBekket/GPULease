import { writeFileSync } from "node:fs";
import { network } from "hardhat";

const { ethers } = await network.connect();

const USDC_ADDRESS = "0x833589fcd6edb6e08f4c7c32d4f71b54bda02913";
const TREASURY_ADDRESS = "0x80CFfF5C710E3e21050AaB272201B574a53A66B3";

async function main() {
  const [deployer] = await ethers.getSigners();
  const deployedAt = new Date().toISOString();

  console.log("Deploying from:", deployer.address);
  console.log("USDC:", USDC_ADDRESS);
  console.log("Treasury:", TREASURY_ADDRESS);

  const GPULeaseWallet = await ethers.getContractFactory("GPULeaseWallet");
  const wallet = await GPULeaseWallet.deploy(USDC_ADDRESS);
  await wallet.waitForDeployment();
  const walletAddress = await wallet.getAddress();
  console.log("GPULeaseWallet:", walletAddress);

  const GPULeaseReferral = await ethers.getContractFactory("GPULeaseReferral");
  const referral = await GPULeaseReferral.deploy();
  await referral.waitForDeployment();
  const referralAddress = await referral.getAddress();
  console.log("GPULeaseReferral:", referralAddress);

  const GPULease = await ethers.getContractFactory("GPULease");
  const gpuLease = await GPULease.deploy(walletAddress, TREASURY_ADDRESS);
  await gpuLease.waitForDeployment();
  const gpuLeaseAddress = await gpuLease.getAddress();
  console.log("GPULease:", gpuLeaseAddress);

  await (await wallet.setLeaseManager(gpuLeaseAddress)).wait();
  await (await gpuLease.setReferralManager(referralAddress)).wait();
  console.log("Linked GPULeaseWallet leaseManager and GPULease referralManager");

  const CampaignMetadataRenderer = await ethers.getContractFactory(
    "CampaignMetadataRenderer"
  );
  const metadataRenderer = await CampaignMetadataRenderer.deploy();
  await metadataRenderer.waitForDeployment();
  const metadataRendererAddress = await metadataRenderer.getAddress();
  console.log("CampaignMetadataRenderer:", metadataRendererAddress);

  const LLMFundraisingFactory = await ethers.getContractFactory(
    "LLMFundraisingFactory"
  );
  const campaignFactory = await LLMFundraisingFactory.deploy(
    USDC_ADDRESS,
    walletAddress,
    metadataRendererAddress,
    TREASURY_ADDRESS
  );
  await campaignFactory.waitForDeployment();
  await campaignFactory.deploymentTransaction()?.wait(1);
  const campaignFactoryAddress = await campaignFactory.getAddress();
  const campaignImplementationAddress =
    await campaignFactory.campaignImplementation();
  console.log("LLMFundraisingFactory:", campaignFactoryAddress);
  console.log("LLMFundraising implementation:", campaignImplementationAddress);

  const Verifier = await ethers.getContractFactory("Verifier");
  const verifier = await Verifier.deploy();
  await verifier.waitForDeployment();
  const verifierAddress = await verifier.getAddress();
  console.log("Verifier:", verifierAddress);

  const addressesText = [
    `# GPULeaseSC Base deployment`,
    `# Last updated at: ${deployedAt}`,
    ``,
    `${gpuLeaseAddress} -- GPULease main SC base address`,
    `${walletAddress} -- GPULeaseWallet base address`,
    `${referralAddress} -- GPULeaseReferral base address`,
    `${campaignFactoryAddress} -- LLMFundraisingFactory base address`,
    `${campaignImplementationAddress} -- LLMFundraising implementation base address`,
    `${metadataRendererAddress} -- CampaignMetadataRenderer base address`,
    `${verifierAddress} -- Verifier base address`,
    ``,
    `${USDC_ADDRESS} -- USDC on base`,
    `${TREASURY_ADDRESS} -- treasury on base`,
    `${deployer.address} -- deployer/admin on base`,
    ``,
  ].join("\n");

  writeFileSync("addresses.txt", addressesText);
  console.log("Updated addresses.txt");
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});

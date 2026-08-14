import { network } from "hardhat";

const USDC_ADDRESS = "0x833589fcd6edb6e08f4c7c32d4f71b54bda02913";
const GPULEASE_WALLET_ADDRESS = "0xf6d56d64938b65c6Ad58cFD447Cd1d74b39eEeF2";
const TREASURY_ADDRESS = "0x80CFfF5C710E3e21050AaB272201B574a53A66B3";

async function main() {
  const { ethers } = await network.connect();

  if (!ethers.isAddress(USDC_ADDRESS)) {
    throw new Error(`Invalid USDC address: ${USDC_ADDRESS}`);
  }

  if (!ethers.isAddress(GPULEASE_WALLET_ADDRESS)) {
    throw new Error(`Invalid GPULeaseWallet address: ${GPULEASE_WALLET_ADDRESS}`);
  }

  const [deployer] = await ethers.getSigners();
  console.log("Deploying LLMFundraisingFactory from:", deployer.address);
  console.log("USDC:", USDC_ADDRESS);
  console.log("GPULeaseWallet:", GPULEASE_WALLET_ADDRESS);

  const MetadataRenderer = await ethers.getContractFactory("CampaignMetadataRenderer");
  const metadataRenderer = await MetadataRenderer.deploy();
  await metadataRenderer.waitForDeployment();
  const metadataRendererAddress = await metadataRenderer.getAddress();
  console.log("Metadata renderer:", metadataRendererAddress);

  const Factory = await ethers.getContractFactory("LLMFundraisingFactory");
  const factory = await Factory.deploy(
    USDC_ADDRESS,
    GPULEASE_WALLET_ADDRESS,
    metadataRendererAddress,
    TREASURY_ADDRESS
  );
  await factory.waitForDeployment();

  const factoryAddress = await factory.getAddress();
  console.log("LLMFundraisingFactory deployed to:", factoryAddress);
  console.log("Constructor args:");
  console.log(`  usdc: ${USDC_ADDRESS}`);
  console.log(`  gpuLeaseWallet: ${GPULEASE_WALLET_ADDRESS}`);
  console.log(`  metadataRenderer: ${metadataRendererAddress}`);
  console.log(`  feeRecipient: ${TREASURY_ADDRESS}`);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});

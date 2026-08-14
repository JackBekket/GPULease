import { network } from "hardhat";

const { ethers } = await network.connect();

const USDC_ADDRESS = "0x833589fcd6edb6e08f4c7c32d4f71b54bda02913";
const GPULEASE_WALLET_ADDRESS = "0xf6d56d64938b65c6Ad58cFD447Cd1d74b39eEeF2";
const METADATA_RENDERER_ADDRESS = "0x48a862d4ca9aE88699251B38097064B8967F405c";
const TREASURY_ADDRESS = "0x80CFfF5C710E3e21050AaB272201B574a53A66B3";

async function main() {
  const [deployer] = await ethers.getSigners();
  console.log("Deploying fee-enabled CampaignFactory from:", deployer.address);

  const Factory = await ethers.getContractFactory("LLMFundraisingFactory");
  const factory = await Factory.deploy(
    USDC_ADDRESS,
    GPULEASE_WALLET_ADDRESS,
    METADATA_RENDERER_ADDRESS,
    TREASURY_ADDRESS,
  );
  await factory.waitForDeployment();

  const deploymentTx = factory.deploymentTransaction();
  const receipt = await deploymentTx?.wait(1);
  const block = receipt
    ? await ethers.provider.getBlock(receipt.blockNumber)
    : null;
  const factoryAddress = await factory.getAddress();
  console.log("Factory:", factoryAddress);
  console.log("Deploy tx:", deploymentTx?.hash);

  let implementationAddress: string | undefined;
  for (let attempt = 0; attempt < 5; attempt++) {
    try {
      implementationAddress = await factory.campaignImplementation();
      break;
    } catch {
      await new Promise((resolve) => setTimeout(resolve, 2_000));
    }
  }
  if (!implementationAddress) {
    throw new Error(
      `RPC did not expose getters yet; deployment succeeded at ${factoryAddress}`,
    );
  }

  const configuration = {
    factory: factoryAddress,
    implementation: implementationAddress,
    usdc: await factory.usdc(),
    wallet: await factory.gpuLeaseWallet(),
    metadataRenderer: await factory.metadataRenderer(),
    feeRecipient: await factory.feeRecipient(),
    deployTx: deploymentTx?.hash,
    blockNumber: receipt?.blockNumber,
    deployedAt: block
      ? new Date(block.timestamp * 1000).toISOString()
      : new Date().toISOString(),
  };

  console.log(JSON.stringify(configuration, null, 2));

  if (
    configuration.usdc.toLowerCase() !== USDC_ADDRESS.toLowerCase() ||
    configuration.wallet.toLowerCase() !==
      GPULEASE_WALLET_ADDRESS.toLowerCase() ||
    configuration.metadataRenderer.toLowerCase() !==
      METADATA_RENDERER_ADDRESS.toLowerCase() ||
    configuration.feeRecipient.toLowerCase() !== TREASURY_ADDRESS.toLowerCase()
  ) {
    throw new Error("Deployed factory configuration mismatch");
  }
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});

import { network } from "hardhat";

const { ethers } = await network.connect();

const USDC_ADDRESS = "0x833589fcd6edb6e08f4c7c32d4f71b54bda02913";
const TREASURY_ADDRESS = "0x80CFfF5C710E3e21050AaB272201B574a53A66B3";
const REFERRAL_ADDRESS = "0x2695d98bF8b233539f5a1Fb823298AA055f2a143";
const METADATA_RENDERER_ADDRESS = "0x48a862d4ca9aE88699251B38097064B8967F405c";

async function deploymentInfo(contract: any) {
  const tx = contract.deploymentTransaction();
  const receipt = await tx?.wait(1);
  const block = receipt ? await ethers.provider.getBlock(receipt.blockNumber) : null;
  return {
    address: await contract.getAddress(),
    txHash: tx?.hash,
    blockNumber: receipt?.blockNumber,
    deployedAt: block
      ? new Date(block.timestamp * 1000).toISOString()
      : new Date().toISOString(),
  };
}

async function main() {
  const [deployer] = await ethers.getSigners();
  console.log("Deploying bonus-enabled stack from:", deployer.address);
  console.log("Deployer ETH:", ethers.formatEther(await ethers.provider.getBalance(deployer.address)));

  const Wallet = await ethers.getContractFactory("GPULeaseWallet");
  const wallet = await Wallet.deploy(USDC_ADDRESS);
  await wallet.waitForDeployment();
  const walletInfo = await deploymentInfo(wallet);
  console.log("GPULeaseWallet:", JSON.stringify(walletInfo));

  const Lease = await ethers.getContractFactory("GPULease");
  const lease = await Lease.deploy(walletInfo.address, TREASURY_ADDRESS);
  await lease.waitForDeployment();
  const leaseInfo = await deploymentInfo(lease);
  console.log("GPULease:", JSON.stringify(leaseInfo));

  const managerTx = await wallet.setLeaseManager(leaseInfo.address);
  await managerTx.wait(1);
  console.log("Wallet manager tx:", managerTx.hash);

  const referralTx = await lease.setReferralManager(REFERRAL_ADDRESS);
  await referralTx.wait(1);
  console.log("Referral tx:", referralTx.hash);

  const Factory = await ethers.getContractFactory("LLMFundraisingFactory");
  const factory = await Factory.deploy(
    USDC_ADDRESS,
    walletInfo.address,
    METADATA_RENDERER_ADDRESS,
    TREASURY_ADDRESS,
  );
  await factory.waitForDeployment();
  const factoryInfo = await deploymentInfo(factory);
  const implementationAddress = await factory.campaignImplementation();
  console.log("LLMFundraisingFactory:", JSON.stringify(factoryInfo));
  console.log("LLMFundraising implementation:", implementationAddress);

  const [manager, configuredWallet, referral, treasury, operator, factoryWallet] =
    await Promise.all([
      wallet.leaseManager(),
      lease.wallet(),
      lease.referralManager(),
      lease.treasury(),
      lease.settlementOperator(),
      factory.gpuLeaseWallet(),
    ]);

  console.log("Verification:", JSON.stringify({
    manager,
    configuredWallet,
    referral,
    treasury,
    operator,
    factoryWallet,
  }));
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});

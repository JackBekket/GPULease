import { network } from "hardhat";

async function main() {
  const { ethers } = await network.connect();
  const [deployer] = await ethers.getSigners();

  console.log("Deploying Verifier from:", deployer.address);

  const Verifier = await ethers.getContractFactory("Verifier");
  const verifier = await Verifier.deploy();
  await verifier.waitForDeployment();

  const verifierAddress = await verifier.getAddress();
  console.log("Verifier deployed to:", verifierAddress);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});

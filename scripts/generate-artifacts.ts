import { mkdirSync, writeFileSync } from "node:fs";
import { artifacts } from "hardhat";

const contracts = [
  {
    artifact: "GPULease",
    abiFile: "GPULease.abi",
  },
  {
    artifact: "GPULeaseWallet",
    abiFile: "GPULeaseWallet.abi",
  },
  {
    artifact: "GPULeaseReferral",
    abiFile: "GPULeaseReferral.abi",
  },
  {
    artifact: "LLMFundraising",
    abiFile: "LLMFundraising.abi",
  },
  {
    artifact: "LLMFundraisingFactory",
    abiFile: "LLMFundraisingFactory.abi",
  },
  {
    artifact: "CampaignMetadataRenderer",
    abiFile: "CampaignMetadataRenderer.abi",
  },
  {
    artifact: "Verifier",
    abiFile: "Verifier.abi",
  },
  {
    artifact: "MockERC20",
    abiFile: "MockERC20.abi",
  },
];

mkdirSync("artifacts/abi", { recursive: true });

for (const contract of contracts) {
  const artifact = await artifacts.readArtifact(contract.artifact);
  writeFileSync(
    `artifacts/abi/${contract.abiFile}`,
    `${JSON.stringify(artifact.abi)}\n`
  );
}

console.log(`Generated ${contracts.length} ABI files in artifacts/abi`);

import { expect } from "chai";
import { network } from "hardhat";

const { ethers } = await network.connect();

describe("CampaignMetadataRenderer", function () {
  it("renders text metadata and SVG image for wallets", async () => {
    const Renderer = await ethers.getContractFactory("CampaignMetadataRenderer");
    const renderer = await Renderer.deploy();

    const campaignName = "Medical LLM crowdfunding";
    const tokenURI = await renderer.tokenURI(campaignName, 7, 4);
    const encodedJson = tokenURI.replace("data:application/json;base64,", "");
    const metadata = JSON.parse(Buffer.from(encodedJson, "base64").toString());

    expect(metadata.name).to.equal(`${campaignName} - Lead Backer`);
    expect(metadata.description).to.equal(
      "Backer reward NFT for a successful LLM fundraising campaign."
    );
    expect(metadata.image).to.match(/^data:image\/svg\+xml;base64,/);
    expect(metadata.attributes).to.deep.include({
      trait_type: "Campaign",
      value: campaignName,
    });
    expect(metadata.attributes).to.deep.include({
      trait_type: "Backer Level",
      value: "Lead Backer",
    });

    const encodedSvg = metadata.image.replace("data:image/svg+xml;base64,", "");
    const svg = Buffer.from(encodedSvg, "base64").toString();
    expect(svg).to.include(campaignName);
    expect(svg).to.include("Lead Backer");
  });
});

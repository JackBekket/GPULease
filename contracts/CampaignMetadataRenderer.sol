// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/utils/Base64.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

contract CampaignMetadataRenderer {
    using Strings for uint256;

    bytes1 private constant JSON_QUOTE = 0x22;
    bytes1 private constant JSON_BACKSLASH = 0x5c;
    bytes1 private constant ASCII_0 = 0x30;
    bytes1 private constant ASCII_B = 0x62;
    bytes1 private constant ASCII_F = 0x66;
    bytes1 private constant ASCII_N = 0x6e;
    bytes1 private constant ASCII_R = 0x72;
    bytes1 private constant ASCII_T = 0x74;
    bytes1 private constant ASCII_U = 0x75;

    function tokenURI(
        string memory campaignName,
        uint256 campaignId,
        uint8 grade
    ) external pure returns (string memory) {
        string memory gradeName = _backerGradeName(grade);
        string memory escapedCampaignName = _escapeJson(campaignName);
        string memory escapedGradeName = _escapeJson(gradeName);

        bytes memory metadata = abi.encodePacked(
            '{"name":"',
            escapedCampaignName,
            " - ",
            escapedGradeName,
            '","description":"Backer reward NFT for a successful LLM fundraising campaign.",',
            '"image":"',
            _imageDataUri(campaignName, gradeName),
            '","attributes":[{"trait_type":"Campaign","value":"',
            escapedCampaignName,
            '"},{"trait_type":"Backer Level","value":"',
            escapedGradeName,
            '"},{"trait_type":"Campaign ID","value":"',
            campaignId.toString(),
            '"}]}'
        );

        return string.concat("data:application/json;base64,", Base64.encode(metadata));
    }

    function _imageDataUri(
        string memory campaignName,
        string memory gradeName
    ) internal pure returns (string memory) {
        bytes memory svg = abi.encodePacked(
            '<svg xmlns="http://www.w3.org/2000/svg" width="800" height="800" viewBox="0 0 800 800">',
            '<rect width="800" height="800" fill="#111827"/>',
            '<rect x="48" y="48" width="704" height="704" rx="40" fill="#f8fafc"/>',
            '<text x="80" y="170" fill="#111827" font-family="Arial, sans-serif" font-size="42" font-weight="700">LLM Fundraising</text>',
            '<text x="80" y="325" fill="#111827" font-family="Arial, sans-serif" font-size="52" font-weight="700">',
            _escapeXml(campaignName),
            '</text>',
            '<text x="80" y="460" fill="#2563eb" font-family="Arial, sans-serif" font-size="48" font-weight="700">',
            _escapeXml(gradeName),
            '</text>',
            '<text x="80" y="610" fill="#64748b" font-family="Arial, sans-serif" font-size="30">Backer reward NFT</text>',
            "</svg>"
        );

        return string.concat("data:image/svg+xml;base64,", Base64.encode(svg));
    }

    function _backerGradeName(uint8 grade) internal pure returns (string memory) {
        if (grade == 4) {
            return "Lead Backer";
        }

        if (grade == 3) {
            return "Founding Backer";
        }

        if (grade == 2) {
            return "Contributor";
        }

        if (grade == 1) {
            return "Supporter";
        }

        return "None";
    }

    function _escapeXml(string memory value) internal pure returns (string memory) {
        bytes memory input = bytes(value);
        bytes memory output = new bytes(input.length * 6);
        uint256 outputLength;

        for (uint256 i = 0; i < input.length; i++) {
            bytes1 char = input[i];

            if (char == 0x26) {
                outputLength = _append(output, outputLength, "&amp;");
            } else if (char == 0x3c) {
                outputLength = _append(output, outputLength, "&lt;");
            } else if (char == 0x3e) {
                outputLength = _append(output, outputLength, "&gt;");
            } else if (char == JSON_QUOTE) {
                outputLength = _append(output, outputLength, "&quot;");
            } else {
                output[outputLength++] = char;
            }
        }

        assembly ("memory-safe") {
            mstore(output, outputLength)
        }

        return string(output);
    }

    function _escapeJson(string memory value) internal pure returns (string memory) {
        bytes memory input = bytes(value);
        bytes memory output = new bytes(input.length * 6);
        uint256 outputLength;

        for (uint256 i = 0; i < input.length; i++) {
            bytes1 char = input[i];

            if (char == JSON_QUOTE) {
                output[outputLength++] = JSON_BACKSLASH;
                output[outputLength++] = JSON_QUOTE;
            } else if (char == JSON_BACKSLASH) {
                output[outputLength++] = JSON_BACKSLASH;
                output[outputLength++] = JSON_BACKSLASH;
            } else if (char == 0x08) {
                output[outputLength++] = JSON_BACKSLASH;
                output[outputLength++] = ASCII_B;
            } else if (char == 0x09) {
                output[outputLength++] = JSON_BACKSLASH;
                output[outputLength++] = ASCII_T;
            } else if (char == 0x0a) {
                output[outputLength++] = JSON_BACKSLASH;
                output[outputLength++] = ASCII_N;
            } else if (char == 0x0c) {
                output[outputLength++] = JSON_BACKSLASH;
                output[outputLength++] = ASCII_F;
            } else if (char == 0x0d) {
                output[outputLength++] = JSON_BACKSLASH;
                output[outputLength++] = ASCII_R;
            } else if (uint8(char) < 0x20) {
                bytes16 hexSymbols = "0123456789abcdef";
                output[outputLength++] = JSON_BACKSLASH;
                output[outputLength++] = ASCII_U;
                output[outputLength++] = ASCII_0;
                output[outputLength++] = ASCII_0;
                output[outputLength++] = hexSymbols[uint8(char) >> 4];
                output[outputLength++] = hexSymbols[uint8(char) & 0x0f];
            } else {
                output[outputLength++] = char;
            }
        }

        assembly ("memory-safe") {
            mstore(output, outputLength)
        }

        return string(output);
    }

    function _append(
        bytes memory output,
        uint256 offset,
        string memory value
    ) internal pure returns (uint256) {
        bytes memory input = bytes(value);
        for (uint256 i = 0; i < input.length; i++) {
            output[offset++] = input[i];
        }
        return offset;
    }
}

// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/proxy/Clones.sol";

import {
    ICampaignParticipantRegistry,
    IGPULeaseWallet,
    LLMFundraising
} from "./Campaign.sol";

contract LLMFundraisingFactory is Ownable, ICampaignParticipantRegistry {
    using Clones for address;

    uint256 public nextCampaignId = 1;

    IERC20 public immutable usdc;
    IGPULeaseWallet public immutable gpuLeaseWallet;
    address public immutable metadataRenderer;
    address public immutable feeRecipient;
    address public immutable campaignImplementation;

    address[] public campaigns;
    mapping(uint256 => address) public campaignById;
    mapping(address => bool) public isCampaign;
    mapping(address => address[]) private _campaignsByCreator;
    mapping(address => address[]) private _campaignsByParticipant;
    mapping(address => mapping(address => bool)) private _participantInCampaign;

    event CampaignCreated(
        uint256 indexed campaignId,
        address indexed campaign,
        address indexed creator,
        uint256 targetAmount,
        uint256 startTimestamp,
        uint256 duration,
        uint256 templateId,
        string campaignName
    );
    event CampaignParticipantRegistered(
        address indexed participant,
        address indexed campaign
    );

    constructor(
        address _usdc,
        address _gpuLeaseWallet,
        address _metadataRenderer,
        address _feeRecipient
    ) Ownable(msg.sender) {
        require(_usdc != address(0), "zero usdc");
        require(_gpuLeaseWallet != address(0), "zero wallet");
        require(_metadataRenderer != address(0), "zero renderer");
        require(_feeRecipient != address(0), "zero fee recipient");

        usdc = IERC20(_usdc);
        gpuLeaseWallet = IGPULeaseWallet(_gpuLeaseWallet);
        metadataRenderer = _metadataRenderer;
        feeRecipient = _feeRecipient;

        campaignImplementation = address(new LLMFundraising());
    }

    function createCampaign(
        uint256 targetAmount,
        uint256 duration,
        uint256 startTimestamp,
        uint256 templateId,
        string calldata campaignName
    ) external returns (uint256 campaignId, address campaign) {
        campaignId = nextCampaignId;
        nextCampaignId += 1;

        campaign = campaignImplementation.clone();
        LLMFundraising(campaign).initialize(
            campaignId,
            targetAmount,
            duration,
            startTimestamp,
            templateId,
            campaignName,
            address(usdc),
            address(gpuLeaseWallet),
            address(this),
            metadataRenderer,
            msg.sender
        );

        campaigns.push(campaign);
        campaignById[campaignId] = campaign;
        isCampaign[campaign] = true;
        _campaignsByCreator[msg.sender].push(campaign);

        emit CampaignCreated(
            campaignId,
            campaign,
            msg.sender,
            LLMFundraising(campaign).targetAmount(),
            startTimestamp,
            duration,
            templateId,
            campaignName
        );
    }

    function campaignsCount() external view returns (uint256) {
        return campaigns.length;
    }

    function campaignsByCreator(
        address creator
    ) external view returns (address[] memory) {
        return _campaignsByCreator[creator];
    }

    function registerParticipant(address participant) external {
        require(isCampaign[msg.sender], "not campaign");
        require(participant != address(0), "zero participant");

        if (_participantInCampaign[participant][msg.sender]) {
            return;
        }

        _participantInCampaign[participant][msg.sender] = true;
        _campaignsByParticipant[participant].push(msg.sender);

        emit CampaignParticipantRegistered(participant, msg.sender);
    }

    function campaignsByParticipant(
        address participant
    ) external view returns (address[] memory) {
        return _campaignsByParticipant[participant];
    }

    function participantCampaignsCount(
        address participant
    ) external view returns (uint256) {
        return _campaignsByParticipant[participant].length;
    }

    function participantCampaignAt(
        address participant,
        uint256 index
    ) external view returns (address) {
        return _campaignsByParticipant[participant][index];
    }

    function hasParticipatedInCampaign(
        address participant,
        address campaign
    ) external view returns (bool) {
        return _participantInCampaign[participant][campaign];
    }
}

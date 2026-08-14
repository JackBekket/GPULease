// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

interface IGPULeaseWallet {
    function depositFor(address beneficiary, uint256 amount) external;
}

interface ICampaignParticipantRegistry {
    function registerParticipant(address participant) external;
    function feeRecipient() external view returns (address);
}

interface ICampaignMetadataRenderer {
    function tokenURI(
        string memory campaignName,
        uint256 campaignId,
        uint8 grade
    ) external view returns (string memory);
}

contract LLMFundraising is Ownable, ERC721, ReentrancyGuard {
    using SafeERC20 for IERC20;

    error ZeroFeeRecipient();
    error DonationExceedsTarget();
    error InvalidCampaignBalance();

    // Backer tiers are calculated in basis points (bps) of the campaign target.
    // BPS is 100%, so 100 bps = 1%, 500 bps = 5%, and 1,500 bps = 15%.
    // A donor's total donations are divided by targetAmount to find their donated
    // percentage, then compared with these minimum thresholds to assign a grade.
    uint256 public constant BPS = 10_000;
    uint256 public constant CROWDFUNDING_FEE_BPS = 500;
    uint256 private constant CONTRIBUTOR_MIN_BPS = 100;
    uint256 private constant FOUNDING_BACKER_MIN_BPS = 500;
    uint256 private constant LEAD_BACKER_MIN_BPS = 1_500;

    enum CampaignState {
        ACTIVE,
        SUCCESS,
        FAILED
    }

    enum BackerGrade {
        NONE,
        SUPPORTER,
        CONTRIBUTOR,
        FOUNDING_BACKER,
        LEAD_BACKER
    }

    uint256 public campaignId;
    uint256 public creatorTargetAmount;
    uint256 public feeAmount;
    uint256 public targetAmount;
    uint256 public startTimestamp;
    uint256 public duration;
    uint256 public templateId;
    string public campaignName;

    IERC20 public usdc;
    IGPULeaseWallet public gpuLeaseWallet;
    ICampaignParticipantRegistry public participantRegistry;
    ICampaignMetadataRenderer public metadataRenderer;
    address public feeRecipient;

    CampaignState public state;
    uint256 public totalRaised;
    uint256 public nextRewardTokenId;
    bool public initialized;

    address[] private _donors;
    mapping(address => uint256) public donations;
    mapping(address => bool) public refunded;
    mapping(address => BackerGrade) public backerGrades;
    mapping(address => bool) public hasDonated;
    mapping(address => uint256) public rewardTokenByDonor;
    mapping(uint256 => BackerGrade) public rewardTokenGrades;

    event Donated(
        address indexed donor,
        uint256 amount,
        uint256 totalDonated,
        BackerGrade grade
    );
    event BackerGradeUpdated(
        address indexed donor,
        BackerGrade previousGrade,
        BackerGrade newGrade,
        uint256 totalDonated,
        uint256 targetShareBps
    );
    event CampaignSucceeded(uint256 totalRaised);
    event CampaignFailed(uint256 totalRaised);
    event Refunded(address indexed donor, uint256 amount);
    event FundsTransferred(address indexed wallet, uint256 amount);
    event CampaignFeePaid(address indexed recipient, uint256 amount);
    event BackerRewardMinted(
        address indexed donor,
        uint256 indexed tokenId,
        string campaignName,
        BackerGrade grade
    );

    constructor() Ownable(msg.sender) ERC721("LLM Fundraising Backer", "LLMBACKER") {
        initialized = true;
    }

    function name() public pure override returns (string memory) {
        return "LLM Fundraising Backer";
    }

    function symbol() public pure override returns (string memory) {
        return "LLMBACKER";
    }

    function initialize(
        uint256 _campaignId,
        uint256 _creatorTargetAmount,
        uint256 _duration,
        uint256 _startTimestamp,
        uint256 _templateId,
        string memory _campaignName,
        address _usdc,
        address _gpuLeaseWallet,
        address _participantRegistry,
        address _metadataRenderer,
        address _campaignOwner
    ) external {
        require(!initialized, "already initialized");
        initialized = true;

        require(_usdc != address(0), "zero usdc");
        require(_gpuLeaseWallet != address(0), "zero wallet");
        require(_participantRegistry != address(0), "zero registry");
        require(_metadataRenderer != address(0), "zero renderer");
        require(_creatorTargetAmount > 0, "zero target");
        require(_duration > 0, "zero duration");
        require(bytes(_campaignName).length > 0, "empty name");

        campaignId = _campaignId;
        creatorTargetAmount = _creatorTargetAmount;
        feeAmount = (_creatorTargetAmount * CROWDFUNDING_FEE_BPS) / BPS;
        targetAmount = _creatorTargetAmount + feeAmount;
        duration = _duration;
        startTimestamp = _startTimestamp;
        templateId = _templateId;
        campaignName = _campaignName;

        usdc = IERC20(_usdc);
        gpuLeaseWallet = IGPULeaseWallet(_gpuLeaseWallet);
        participantRegistry = ICampaignParticipantRegistry(
            _participantRegistry
        );
        metadataRenderer = ICampaignMetadataRenderer(_metadataRenderer);
        feeRecipient = ICampaignParticipantRegistry(_participantRegistry)
            .feeRecipient();
        if (feeRecipient == address(0)) revert ZeroFeeRecipient();

        state = CampaignState.ACTIVE;
        nextRewardTokenId = 1;
        _transferOwnership(_campaignOwner);
    }

    function deadline() public view returns (uint256) {
        return startTimestamp + duration;
    }

    function isExpired() public view returns (bool) {
        return block.timestamp >= deadline();
    }

    function isTargetReached() public view returns (bool) {
        return totalRaised >= targetAmount;
    }

    function checkConditions()
        public
        view
        returns (bool expired, bool reached)
    {
        expired = isExpired();
        reached = isTargetReached();
    }

    function donorShareBps(address donor) public view returns (uint256) {
        return (donations[donor] * BPS) / targetAmount;
    }

    function gradeForDonation(
        address donor
    ) external view returns (BackerGrade) {
        return _bestAvailableGrade(donations[donor]);
    }

    function donorsCount() external view returns (uint256) {
        return _donors.length;
    }

    function donorAt(uint256 index) external view returns (address) {
        return _donors[index];
    }

    function donors() external view returns (address[] memory) {
        return _donors;
    }

    function donorsSlice(
        uint256 offset,
        uint256 limit
    ) external view returns (address[] memory donors_) {
        uint256 donorCount = _donors.length;
        if (offset >= donorCount) {
            return new address[](0);
        }

        uint256 end = offset + limit;
        if (end > donorCount) {
            end = donorCount;
        }

        donors_ = new address[](end - offset);
        for (uint256 i = offset; i < end; i++) {
            donors_[i - offset] = _donors[i];
        }
    }

    function donorInfo(
        address donor
    )
        external
        view
        returns (
            bool participated,
            uint256 donatedAmount,
            uint256 targetShareBps,
            BackerGrade grade,
            bool wasRefunded,
            uint256 rewardTokenId
        )
    {
        participated = hasDonated[donor];
        donatedAmount = donations[donor];
        targetShareBps = donorShareBps(donor);
        grade = backerGrades[donor];
        wasRefunded = refunded[donor];
        rewardTokenId = rewardTokenByDonor[donor];
    }

    function claimBackerReward() external nonReentrant returns (uint256 tokenId) {
        return _mintBackerReward(msg.sender);
    }

    function mintBackerRewards(
        uint256 offset,
        uint256 limit
    ) external nonReentrant returns (uint256 minted) {
        require(limit > 0, "zero limit");
        uint256 donorCount = _donors.length;
        require(offset < donorCount, "offset out of range");

        uint256 end = offset + limit;
        if (end > donorCount) {
            end = donorCount;
        }

        for (uint256 i = offset; i < end; i++) {
            address donor = _donors[i];
            if (rewardTokenByDonor[donor] == 0 && donations[donor] > 0) {
                _mintBackerReward(donor);
                minted += 1;
            }
        }
    }

    function tokenURI(
        uint256 tokenId
    ) public view override returns (string memory) {
        _requireOwned(tokenId);
        return metadataRenderer.tokenURI(
            campaignName,
            campaignId,
            uint8(rewardTokenGrades[tokenId])
        );
    }

    function donate(uint256 amount) external nonReentrant {
        require(state == CampaignState.ACTIVE, "not active");
        require(block.timestamp >= startTimestamp, "not started");
        require(!isExpired(), "expired");
        require(amount > 0, "zero amount");
        if (amount > targetAmount - totalRaised) revert DonationExceedsTarget();

        usdc.safeTransferFrom(msg.sender, address(this), amount);

        if (!hasDonated[msg.sender]) {
            hasDonated[msg.sender] = true;
            _donors.push(msg.sender);
            participantRegistry.registerParticipant(msg.sender);
        }

        donations[msg.sender] += amount;
        totalRaised += amount;

        BackerGrade grade = _updateBackerGrade(msg.sender);

        emit Donated(msg.sender, amount, donations[msg.sender], grade);

        _evaluateState();
    }

    function checkState() external {
        require(state == CampaignState.ACTIVE, "already closed");
        _evaluateState();
    }

    function refund() external nonReentrant {
        require(state == CampaignState.FAILED, "not failed");

        uint256 amount = donations[msg.sender];
        require(amount > 0, "nothing to refund");
        require(!refunded[msg.sender], "already refunded");

        refunded[msg.sender] = true;
        donations[msg.sender] = 0;
        _setBackerGrade(msg.sender, BackerGrade.NONE);

        usdc.safeTransfer(msg.sender, amount);

        emit Refunded(msg.sender, amount);
    }

    function _evaluateState() internal {
        if (isTargetReached()) {
            _markSuccess();
        } else if (isExpired()) {
            _markFailed();
        }
    }

    function _markSuccess() internal {
        require(state == CampaignState.ACTIVE, "not active");

        uint256 balance = usdc.balanceOf(address(this));
        if (balance != targetAmount) revert InvalidCampaignBalance();

        state = CampaignState.SUCCESS;

        _transferToWallet(creatorTargetAmount);
        if (feeAmount > 0) {
            usdc.safeTransfer(feeRecipient, feeAmount);
            emit CampaignFeePaid(feeRecipient, feeAmount);
        }

        emit CampaignSucceeded(balance);
    }

    function _markFailed() internal {
        require(state == CampaignState.ACTIVE, "not active");

        state = CampaignState.FAILED;

        emit CampaignFailed(totalRaised);
    }

    function _transferToWallet(uint256 amount) internal {
        require(amount > 0, "no funds");

        usdc.forceApprove(address(gpuLeaseWallet), amount);

        gpuLeaseWallet.depositFor(owner(), amount);

        emit FundsTransferred(address(gpuLeaseWallet), amount);
    }

    function _mintBackerReward(
        address donor
    ) internal returns (uint256 tokenId) {
        require(state == CampaignState.SUCCESS, "not successful");
        require(donations[donor] > 0, "not donor");
        require(rewardTokenByDonor[donor] == 0, "already minted");

        BackerGrade grade = backerGrades[donor];
        require(grade != BackerGrade.NONE, "no grade");

        tokenId = nextRewardTokenId;
        nextRewardTokenId += 1;

        rewardTokenByDonor[donor] = tokenId;
        rewardTokenGrades[tokenId] = grade;

        _safeMint(donor, tokenId);

        emit BackerRewardMinted(donor, tokenId, campaignName, grade);
    }

    function _targetShareBps(uint256 amount) internal view returns (uint256) {
        return (amount * BPS) / targetAmount;
    }

    function _bestAvailableGrade(
        uint256 amount
    ) internal view returns (BackerGrade) {
        if (amount == 0) {
            return BackerGrade.NONE;
        }

        uint256 shareBps = _targetShareBps(amount);

        if (shareBps >= LEAD_BACKER_MIN_BPS) {
            return BackerGrade.LEAD_BACKER;
        }

        if (shareBps >= FOUNDING_BACKER_MIN_BPS) {
            return BackerGrade.FOUNDING_BACKER;
        }

        if (shareBps >= CONTRIBUTOR_MIN_BPS) {
            return BackerGrade.CONTRIBUTOR;
        }

        return BackerGrade.SUPPORTER;
    }

    function _updateBackerGrade(
        address donor
    ) internal returns (BackerGrade) {
        BackerGrade nextGrade = _bestAvailableGrade(donations[donor]);
        _setBackerGrade(donor, nextGrade);
        return nextGrade;
    }

    function _setBackerGrade(address donor, BackerGrade nextGrade) internal {
        BackerGrade previousGrade = backerGrades[donor];
        if (previousGrade == nextGrade) {
            return;
        }

        backerGrades[donor] = nextGrade;

        emit BackerGradeUpdated(
            donor,
            previousGrade,
            nextGrade,
            donations[donor],
            donorShareBps(donor)
        );
    }

}

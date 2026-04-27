// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface IGPULease {
    function deposit(uint256 amount) external;
}

contract LLMFundraising is Ownable, ReentrancyGuard {
    using SafeERC20 for IERC20;

    enum CampaignState {
        ACTIVE,
        SUCCESS,
        FAILED
    }

    // immutable config
    uint256 public immutable campaignId;
    uint256 public immutable targetAmount;
    uint256 public immutable startTimestamp;
    uint256 public immutable duration;
    uint256 public immutable templateId;

    IERC20 public immutable usdc;
    IGPULease public immutable gpuLease;

    // state
    CampaignState public state;
    uint256 public totalRaised;

    mapping(address => uint256) public donations;
    mapping(address => bool) public refunded;

    // events
    event Donated(address indexed donor, uint256 amount);
    event CampaignSucceeded(uint256 totalRaised);
    event CampaignFailed(uint256 totalRaised);
    event Refunded(address indexed donor, uint256 amount);
    event FundsTransferred(address indexed gpuLease, uint256 amount);

    constructor(
        uint256 _campaignId,
        uint256 _targetAmount,
        uint256 _duration,
        uint256 _startTimestamp,
        uint256 _templateId,
        address _usdc,
        address _gpuLease
    ) Ownable(msg.sender) {
        require(_usdc != address(0), "zero usdc");
        require(_gpuLease != address(0), "zero gpuLease");
        require(_targetAmount > 0, "zero target");
        require(_duration > 0, "zero duration");

        campaignId = _campaignId;
        targetAmount = _targetAmount;
        duration = _duration;
        startTimestamp = _startTimestamp;
        templateId = _templateId;

        usdc = IERC20(_usdc);
        gpuLease = IGPULease(_gpuLease);

        state = CampaignState.ACTIVE;
    }

    // =========================
    // VIEW LOGIC
    // =========================

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

    // =========================
    // DONATION LOGIC
    // =========================

    function donate(uint256 amount) external nonReentrant {
        require(state == CampaignState.ACTIVE, "not active");
        require(block.timestamp >= startTimestamp, "not started");
        require(!isExpired(), "expired");
        require(amount > 0, "zero amount");

        usdc.safeTransferFrom(msg.sender, address(this), amount);

        donations[msg.sender] += amount;
        totalRaised += amount;

        emit Donated(msg.sender, amount);

        _evaluateState();
    }

    // =========================
    // PUBLIC TICK FUNCTION
    // =========================

    function checkState() external {
        require(state == CampaignState.ACTIVE, "already closed");
        _evaluateState();
    }

    function _evaluateState() internal {
        if (isTargetReached()) {
            _markSuccess();
        } else if (isExpired()) {
            _markFailed();
        }
    }

    // =========================
    // SUCCESS / FAILURE
    // =========================

    function _markSuccess() internal {
        require(state == CampaignState.ACTIVE, "not active");

        uint256 balance = usdc.balanceOf(address(this));
        require(balance >= targetAmount, "insufficient funds");

        state = CampaignState.SUCCESS;

        _transferToGPULease(balance);

        emit CampaignSucceeded(balance);
    }

    function _markFailed() internal {
        require(state == CampaignState.ACTIVE, "not active");

        state = CampaignState.FAILED;

        emit CampaignFailed(totalRaised);
    }

    function _transferToGPULease(uint256 amount) internal {
        require(amount > 0, "no funds");

        // reset approve (USDC safety pattern)
        usdc.safeApprove(address(gpuLease), 0);
        usdc.safeApprove(address(gpuLease), amount);

        gpuLease.deposit(amount);

        emit FundsTransferred(address(gpuLease), amount);
    }

    // =========================
    // REFUND LOGIC
    // =========================

    function refund() external nonReentrant {
        require(state == CampaignState.FAILED, "not failed");

        uint256 amount = donations[msg.sender];
        require(amount > 0, "nothing to refund");
        require(!refunded[msg.sender], "already refunded");

        refunded[msg.sender] = true;
        donations[msg.sender] = 0;

        usdc.safeTransfer(msg.sender, amount);

        emit Refunded(msg.sender, amount);
    }
}
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

contract GPULease is AccessControl, ReentrancyGuard {
    
    struct Lease {
        address user;
        address provider;
        uint startTime;
        uint duration;
        uint storagePricePerSecond; // Price per second for storage
        uint computePricePerSecond; // Price per second for computation
        uint totalAmount; // Total amount to be paid (both storage and compute)
        bool active;
        bool completed;
        bool paused; // Lease can be paused during execution
        uint lastPausedTime; // Time when lease was last paused
        uint pausedDuration; // Cumulative duration of pauses in seconds
    }

    // User balances mapping
    mapping(address => uint256) public userBalances;
    
    // Locked funds for leases mapping  
    mapping(uint => uint256) public lockedFunds;
    
    IERC20 public token;
    mapping(uint => Lease) public leases;
    uint public leaseCount = 0;
    
    // Platform fee parameters
    uint public platformFeePercentage = 5; // 5% platform fee
    
    event LeaseStarted(uint leaseId, address user, address provider, uint duration, uint amount);
    event LeaseCompleted(uint leaseId, uint refund, uint providerReward);
    event LeaseCancelled(uint leaseId, uint refund);
    event PaymentReceived(uint leaseId, uint amount);
    event PlatformFeeCollected(uint leaseId, uint feeAmount);
    event UserDeposited(address user, uint amount);
    event UserWithdrawn(address user, uint amount);
    event LeasePaused(uint leaseId);
    event LeaseResumed(uint leaseId);
    
    address public deployer;


    // operation with leaseID shoud be available only by deal participant or the admin of the contract (in case of delegated call)
    modifier onlyLeaseOwner(uint _leaseId) {
        require(leases[_leaseId].user == msg.sender || leases[_leaseId].provider == msg.sender || msg.sender == deployer, "Only user or provider can call this");
        _;
    }

    // Only admin can start lease on behalf of a user
    modifier onlyAdminOrContract() {
        require(msg.sender == deployer || msg.sender == address(this), "Only admin or contract itself can call this");
        _;
    }
    
    constructor(address _token) {
        token = IERC20(_token);
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender); // Grant default admin role to deployer
        deployer = msg.sender;
    }
    
    function setPlatformFee(uint _feePercentage) external onlyRole(DEFAULT_ADMIN_ROLE) {
        platformFeePercentage = _feePercentage;
    }
    
    /**
     * @dev Deposit tokens to user balance
     */
    function deposit(uint256 amount) external nonReentrant {
        require(amount > 0, "Deposit amount must be > 0");
        token.transferFrom(msg.sender, address(this), amount);
        userBalances[msg.sender] = userBalances[msg.sender] + amount;
        emit UserDeposited(msg.sender, amount);
    }
    
    /**
     * @dev Withdraw tokens from user balance
     */
    function withdraw(uint256 amount) external nonReentrant {
        require(userBalances[msg.sender] >= amount, "Insufficient balance");
        userBalances[msg.sender] = userBalances[msg.sender] - amount;
        token.transfer(msg.sender, amount);
        emit UserWithdrawn(msg.sender, amount);
    }
    
    function startLeaseWithUser(
        uint _duration,
        uint _storagePricePerSecond,
        uint _computePricePerSecond,
        address _provider,
        address _user
    ) public onlyAdminOrContract nonReentrant returns (uint leaseId) {
        require(_duration > 0, "Duration must be > 0");
        require(_storagePricePerSecond > 0 || _computePricePerSecond > 0, "At least one price must be > 0");
        
        // Calculate total amounts for both storage and compute
        uint totalStorageAmount = _duration * _storagePricePerSecond;
        uint totalComputeAmount = _duration * _computePricePerSecond;
        uint totalAmount = totalStorageAmount + totalComputeAmount;
        
        require(userBalances[_user] >= totalAmount, "Insufficient token balance");
        
        // Calculate platform fee
        uint platformFee = (totalAmount * platformFeePercentage) / 100;
        
        // Deduct funds from user balance and lock them in lockedFunds mapping by leaseId
        userBalances[_user] = userBalances[_user] - totalAmount;
        lockedFunds[leaseCount] = totalAmount;
        
        leaseId = leaseCount;
        leaseCount++;
        
        leases[leaseId] = Lease({
            user: _user,
            provider: _provider,
            startTime: block.timestamp,
            duration: _duration,
            storagePricePerSecond: _storagePricePerSecond,
            computePricePerSecond: _computePricePerSecond,
            totalAmount: totalAmount,
            active: true,
            completed: false,
            paused: false,
            lastPausedTime: 0,
            pausedDuration: 0
        });
        
        emit LeaseStarted(leaseId, _user, _provider, _duration, totalAmount);
        emit PaymentReceived(leaseId, totalAmount);
        emit PlatformFeeCollected(leaseId, platformFee);
        
        return leaseId;
    }
    
    function startLease(
        uint _duration,
        uint _storagePricePerSecond,
        uint _computePricePerSecond,
        address _provider
    ) external nonReentrant returns (uint leaseId) {
        // Anyone can call this - they start a lease for themselves
        uint lid = this.startLeaseWithUser(_duration, _storagePricePerSecond, _computePricePerSecond, _provider, msg.sender);
        return lid;
    }
    
    function pauseLease(uint _leaseId) external onlyLeaseOwner(_leaseId) nonReentrant {
        Lease storage lease = leases[_leaseId];
        require(lease.active, "Lease is not active");
        require(!lease.completed, "Lease already completed");
        require(!lease.paused, "Lease is already paused");
        
        // Set the pause time
        lease.lastPausedTime = block.timestamp;
        lease.paused = true;
        
        emit LeasePaused(_leaseId);
    }
    
    function resumeLease(uint _leaseId) external onlyLeaseOwner(_leaseId) nonReentrant {
        Lease storage lease = leases[_leaseId];
        require(lease.active, "Lease is not active");
        require(!lease.completed, "Lease already completed");
        require(lease.paused, "Lease is not paused");
        
        // Calculate the duration of this pause
        uint pauseDuration = block.timestamp - lease.lastPausedTime;
        lease.pausedDuration += pauseDuration;
        lease.lastPausedTime = 0; // Reset last paused time
        lease.paused = false;
        
        emit LeaseResumed(_leaseId);
    }
    
    function calculateActualCost(uint _leaseId) internal view returns (uint actualStorageCost, uint actualComputeCost) {
        Lease storage lease = leases[_leaseId];
        
        // Calculate the effective duration by excluding paused time
        uint effectiveDuration;
        if (lease.paused) {
            // If currently paused, don't count the current pause period in the effective duration 
            effectiveDuration = block.timestamp - lease.startTime - lease.pausedDuration;
        } else {
            // If not paused, use full duration minus pauses
            effectiveDuration = block.timestamp - lease.startTime - lease.pausedDuration;
        }
        
        // Ensure we have a valid duration (cannot be negative)
        if (effectiveDuration > lease.duration) {
            effectiveDuration = lease.duration;
        }
        
        actualStorageCost = effectiveDuration * lease.storagePricePerSecond;
        actualComputeCost = effectiveDuration * lease.computePricePerSecond;
    }
    
    function completeLease(uint _leaseId) external onlyLeaseOwner(_leaseId) nonReentrant {
        Lease storage lease = leases[_leaseId];
        require(lease.active, "Lease is not active");
        require(!lease.completed, "Lease already completed");
        
        uint actualStorageCost;
        uint actualComputeCost;
        (actualStorageCost, actualComputeCost) = calculateActualCost(_leaseId);
        
        // Total cost based on the effective duration
        uint actualTotalCost = actualStorageCost + actualComputeCost; 
        uint refund = lease.totalAmount - actualTotalCost;
        
        // Calculate platform fee from the total actual cost
        uint platformFee = (actualTotalCost * platformFeePercentage) / 100;
        uint providerAmount = actualTotalCost - platformFee;
        
        // Refund unused amount back to user's balance
        if (refund > 0) {
            userBalances[lease.user] = userBalances[lease.user] + refund;
        }
        
        // Transfer actual cost minus platform fee to provider
        require(token.transfer(lease.provider, providerAmount), "Provider transfer failed");
        
        // Collect platform fee 
        if (platformFee > 0) {
            require(token.transfer(deployer, platformFee), "Platform fee transfer failed");
        }
        
        // Unlock the funds from lockedFunds mapping 
        delete lockedFunds[_leaseId];
        
        lease.completed = true;
        lease.active = false;
        
        emit LeaseCompleted(_leaseId, refund, providerAmount);
    }
    
    function cancelLease(uint _leaseId) external onlyLeaseOwner(_leaseId) nonReentrant {
        Lease storage lease = leases[_leaseId];
        require(lease.active, "Lease is not active");
        require(!lease.completed, "Lease already completed");
        
        // Cancel within 5 minutes (300 seconds)
        require(block.timestamp < lease.startTime + 300, "Cannot cancel after 5 minutes");
        
        uint refund = lease.totalAmount;
        
        // Return funds back to user's balance
        userBalances[lease.user] = userBalances[lease.user] + refund;
        
        // Unlock the funds from lockedFunds mapping 
        delete lockedFunds[_leaseId];
        
        lease.active = false;
        lease.completed = true;
        
        emit LeaseCancelled(_leaseId, refund);
    }
    
    function getLeaseStatus(uint _leaseId) external view returns (bool active, bool completed, bool paused, uint startTime, uint duration, uint pausedDuration) {
        Lease storage lease = leases[_leaseId];
        return (lease.active, lease.completed, lease.paused, lease.startTime, lease.duration, lease.pausedDuration);
    }
    
    function getContractBalance() external view returns (uint) {
        return token.balanceOf(address(this));
    }
}
// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./GPULeaseWallet.sol";

interface IGPULeaseReferral {
    function referrerOf(address user) external view returns (address);
    function referralShareBps() external view returns (uint);
}

contract GPULease is Ownable {
    GPULeaseWallet public immutable wallet;
    IGPULeaseReferral public referralManager;
    address public treasury;

     struct Lease {
        address user;
        address provider;
        uint startTime;
        uint duration;
        uint storagePricePerSecond; // Price per second for storage
        uint computePricePerSecond; // Price per second for computation
        uint leaseFeePercentage;
        bool active;
        bool completed;
        bool paused; // Lease can be paused during execution 
        uint pausedAt; // Time when lease was paused
        uint pausedDuration; // Cumulative duration of pauses in seconds
    }

    struct FrozenFundsInfo {
        uint leaseId;
        uint256 amount;
    }

    struct LeaseReferralInfo {
        address referrer;
        uint referralShareBps;
    }

    struct LeaseRequest {
        uint duration;
        uint storagePricePerSecond;
        uint computePricePerSecond;
        address provider;
        address user;
    }

    mapping(address => uint[]) public userActiveLeases;
    mapping(address => uint) public userFeePercentage;
    mapping(address => bool) public hasCustomUserFee;
    mapping(uint => uint256) public frozenFunds;
    mapping(uint => Lease) public leases;
    mapping(uint => LeaseReferralInfo) public leaseReferralInfo;
    uint public leaseCount = 0;

    uint public platformFeePercentage = 10; // 10% default platform fee

    event PlatformFeeUpdated(uint previousFeePercentage, uint newFeePercentage);
    event UserFeeUpdated(address indexed user, uint feePercentage);
    event UserFeeCleared(address indexed user);
    event ReferralManagerUpdated(address indexed previousManager, address indexed newManager);
    event LeaseStarted(uint leaseId, address user, address provider, uint duration, uint amount);
    event LeaseCompleted(uint leaseId);
    event LeasePaused(uint leaseId);
    event LeaseResumed(uint leaseId);

 

    constructor(address wallet_, address treasury_) Ownable(msg.sender) {
        require(wallet_ != address(0), "zero wallet");
        require(treasury_ != address(0), "zero treasury");
        wallet = GPULeaseWallet(wallet_);
        treasury = treasury_;
    }


    //admin stuff
    function setPlatformFee(uint _feePercentage) public onlyOwner {
        require(_feePercentage <= 100, "Fee too high");
        emit PlatformFeeUpdated(platformFeePercentage, _feePercentage);
        platformFeePercentage = _feePercentage;
    }

    function setUserFee(address user, uint _feePercentage) external onlyOwner {
        require(user != address(0), "invalid user");
        require(_feePercentage <= 100, "Fee too high");
        userFeePercentage[user] = _feePercentage;
        hasCustomUserFee[user] = true;
        emit UserFeeUpdated(user, _feePercentage);
    }

    function clearUserFee(address user) external onlyOwner {
        require(user != address(0), "invalid user");
        delete userFeePercentage[user];
        delete hasCustomUserFee[user];
        emit UserFeeCleared(user);
    }

    function setReferralManager(address newReferralManager) external onlyOwner {
        emit ReferralManagerUpdated(address(referralManager), newReferralManager);
        referralManager = IGPULeaseReferral(newReferralManager);
    }

    function setTreasury(address newTreasury) external onlyOwner {
        require(newTreasury != address(0), "zero treasury");
        wallet.moveBalance(treasury, newTreasury);
        treasury = newTreasury;
    }

    function feePercentageForUser(address user) public view returns (uint) {
        if (hasCustomUserFee[user]) {
            return userFeePercentage[user];
        }

        return platformFeePercentage;
    }


    //lease stuff
    function startLease(
        uint _duration,
        uint _storagePricePerSecond,
        uint _computePricePerSecond,
        address _provider,
        address _user
    ) public onlyOwner returns (uint leaseId) {
        return _startLease(LeaseRequest({
            duration: _duration,
            storagePricePerSecond: _storagePricePerSecond,
            computePricePerSecond: _computePricePerSecond,
            provider: _provider,
            user: _user
        }));
    }

    function _startLease(LeaseRequest memory request) internal returns (uint leaseId) {
        require(request.duration > 0, "Duration must be > 0");
        require(request.storagePricePerSecond > 0 || request.computePricePerSecond > 0, "At least one price must be > 0");
        require(request.user != address(0), "invalid user");
        require(request.provider != address(0), "invalid provider");
        
        uint totalAmount;
        uint leaseFeePercentage = feePercentageForUser(request.user);
        address referrer;
        uint referralShareBps;

        {
            // Calculate total amounts for both storage and compute
            totalAmount =
                (request.duration * request.storagePricePerSecond) +
                (request.duration * request.computePricePerSecond);
            uint platformFee = (totalAmount * leaseFeePercentage) / 100;
            totalAmount += platformFee;
            (referrer, referralShareBps) = _referralInfoForLease(request.user);
        }
       
        // Deduct funds from user balance and lock them in lockedFunds mapping by leaseId
        wallet.debitBalance(request.user, totalAmount);
        frozenFunds[leaseCount] = totalAmount;
        
        leaseId = leaseCount;
        leaseCount++;
        
        leases[leaseId] = Lease({
            user: request.user,
            provider: request.provider,
            startTime: block.timestamp - 5 minutes, //so we won't need any cancel function
            duration: request.duration,
            storagePricePerSecond: request.storagePricePerSecond,
            computePricePerSecond: request.computePricePerSecond,
            leaseFeePercentage: leaseFeePercentage,
            active: true,
            completed: false,
            paused: false,
            pausedAt: 0,
            pausedDuration: 0
        });
        leaseReferralInfo[leaseId] = LeaseReferralInfo({
            referrer: referrer,
            referralShareBps: referralShareBps
        });
        userActiveLeases[request.user].push(leaseId);
        emit LeaseStarted(leaseId, leases[leaseId].user, leases[leaseId].provider, leases[leaseId].duration, frozenFunds[leaseId]);
        return leaseId;
    }

     function pauseLease(uint _leaseId) public onlyOwner {
        Lease storage lease = leases[_leaseId];
        require(lease.active, "Lease is not active");
        require(!lease.completed, "Lease already completed");
        require(!lease.paused, "Lease is already paused");
        
        // Set the pause time
        lease.pausedAt = block.timestamp;
        lease.paused = true;
        
        emit LeasePaused(_leaseId);
    }
    
    function resumeLease(uint _leaseId) external onlyOwner {
        Lease storage lease = leases[_leaseId];
        require(lease.active, "Lease is not active");
        require(!lease.completed, "Lease already completed");
        require(lease.paused, "Lease is not paused");
        
        uint pauseDuration = block.timestamp - lease.pausedAt;
        lease.pausedDuration += pauseDuration;
        lease.pausedAt = 0; // Reset last paused time
        lease.paused = false;
        
        emit LeaseResumed(_leaseId);
    }

    function completeLease(uint _leaseId) external onlyOwner  {
        Lease storage lease = leases[_leaseId];
        require(lease.active, "Lease is not active");
        require(!lease.completed, "Lease already completed");
        
        uint actualStorageCost;
        uint actualComputeCost;
        (actualStorageCost, actualComputeCost) = calculateActualCost(_leaseId);
        
        // Total cost based on the effective duration
        uint actualTotalCost = actualStorageCost + actualComputeCost; 
        
        
        // Calculate platform fee from the total actual cost
        uint platformFee = (actualTotalCost * lease.leaseFeePercentage) / 100;
        uint referralAmount = 0;
        LeaseReferralInfo storage referralInfo = leaseReferralInfo[_leaseId];
        if (referralInfo.referrer != address(0) && referralInfo.referralShareBps > 0) {
            referralAmount = (platformFee * referralInfo.referralShareBps) / 10_000;
            wallet.creditBalance(referralInfo.referrer, referralAmount);
        }

        wallet.creditBalance(treasury, platformFee - referralAmount);
        frozenFunds[_leaseId] -= platformFee;

        uint providerAmount = actualTotalCost - platformFee;
        wallet.creditBalance(lease.provider, providerAmount);
        frozenFunds[_leaseId] -= providerAmount;

        wallet.creditBalance(lease.user, frozenFunds[_leaseId]);

        delete frozenFunds[_leaseId];
        
        lease.completed = true;
        lease.active = false;

            address user = leases[_leaseId].user;
    uint[] storage leasesList = userActiveLeases[user];
    for (uint i = 0; i < leasesList.length; i++) {
        if (leasesList[i] == _leaseId) {
            leasesList[i] = leasesList[leasesList.length - 1];
            leasesList.pop();
            break;
        }
    }
        emit LeaseCompleted(_leaseId);
    }


    function calculateActualCost(uint _leaseId)
    internal
    view
    returns (uint actualStorageCost, uint actualComputeCost)
{
    Lease storage lease = leases[_leaseId];

    require(lease.startTime > 0, "Lease not started");

    uint elapsed = block.timestamp - lease.startTime;

    if (elapsed > lease.duration) {
        elapsed = lease.duration;
    }

    uint totalPaused = lease.pausedDuration;

    if (lease.paused) {
        uint currentPause = block.timestamp - lease.pausedAt;
        totalPaused += currentPause;
    }

    if (totalPaused > elapsed) {
        totalPaused = elapsed;
    }

    uint activeDuration = elapsed - totalPaused;

    actualStorageCost = elapsed * lease.storagePricePerSecond;
    actualComputeCost = activeDuration * lease.computePricePerSecond;

    return (actualStorageCost, actualComputeCost);
}

    function _referralInfoForLease(address user)
        internal
        view
        returns (address referrer, uint referralShareBps)
    {
        if (address(referralManager) == address(0)) {
            return (address(0), 0);
        }

        referrer = referralManager.referrerOf(user);
        if (referrer == address(0)) {
            return (address(0), 0);
        }

        referralShareBps = referralManager.referralShareBps();
        require(referralShareBps <= 10_000, "Referral share too high");
    }
    
   function getUserFrozenFunds(address user) 
    external
    view
    returns (FrozenFundsInfo[] memory result) 
{
    uint[] storage leasesList = userActiveLeases[user];
    result = new FrozenFundsInfo[](leasesList.length);

    for (uint i = 0; i < leasesList.length; i++) {
        uint leaseId = leasesList[i];

        result[i] = FrozenFundsInfo({
            leaseId: leaseId,
            amount: frozenFunds[leaseId]
        });
    }

    return result;
}

}

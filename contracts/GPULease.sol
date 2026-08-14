// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./GPULeaseWallet.sol";

interface IGPULeaseReferral {
    function referrerOf(address user) external view returns (address);
    function referralShareBps() external view returns (uint);
}

contract GPULease is Ownable {
    uint public constant SETTLEMENT_INTERVAL = 1 days;
    uint public constant MAX_SETTLEMENT_BATCH_SIZE = 50;

    GPULeaseWallet public immutable wallet;
    IGPULeaseReferral public referralManager;
    address public treasury;
    address public settlementOperator;

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
        uint startTimestamp;
        uint duration;
        uint storagePricePerSecond;
        uint computePricePerSecond;
        address provider;
        address user;
    }

    struct LeaseSettlement {
        uint providerPaid;
        uint feePaid;
        uint referralPaid;
        uint lastSettledAt;
    }

    mapping(address => uint[]) public userActiveLeases;
    mapping(address => uint) public userFeePercentage;
    mapping(address => bool) public hasCustomUserFee;
    mapping(uint => uint256) public frozenFunds;
    mapping(uint => uint256) public frozenCashFunds;
    mapping(uint => uint256) public frozenBonusFunds;
    mapping(uint => Lease) public leases;
    mapping(uint => LeaseReferralInfo) public leaseReferralInfo;
    mapping(uint => LeaseSettlement) public leaseSettlements;
    mapping(uint => uint) public leaseActivationTime;
    uint public leaseCount = 0;

    uint public platformFeePercentage = 10; // 10% default platform fee

    event PlatformFeeUpdated(uint previousFeePercentage, uint newFeePercentage);
    event UserFeeUpdated(address indexed user, uint feePercentage);
    event UserFeeCleared(address indexed user);
    event ReferralManagerUpdated(address indexed previousManager, address indexed newManager);
    event LeaseStarted(
        uint leaseId,
        address user,
        address provider,
        uint startTimestamp,
        uint activationTimestamp,
        uint duration,
        uint amount
    );
    event LeaseCompleted(uint leaseId);
    event LeasePaused(uint leaseId);
    event LeaseResumed(uint leaseId);
    event SettlementOperatorUpdated(address indexed previousOperator, address indexed newOperator);
    event LeaseSettled(
        uint indexed leaseId,
        uint providerAmount,
        uint platformFee,
        uint referralAmount,
        uint settledAt
    );

    modifier onlySettlementOperator() {
        require(msg.sender == owner() || msg.sender == settlementOperator, "not settlement operator");
        _;
    }

    constructor(address wallet_, address treasury_) Ownable(msg.sender) {
        require(wallet_ != address(0), "zero wallet");
        require(treasury_ != address(0), "zero treasury");
        wallet = GPULeaseWallet(wallet_);
        treasury = treasury_;
        settlementOperator = msg.sender;
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


    function setSettlementOperator(address newSettlementOperator) external onlyOwner {
        require(newSettlementOperator != address(0), "zero settlement operator");
        emit SettlementOperatorUpdated(settlementOperator, newSettlementOperator);
        settlementOperator = newSettlementOperator;
    }

    function feePercentageForUser(address user) public view returns (uint) {
        if (hasCustomUserFee[user]) {
            return userFeePercentage[user];
        }

        return platformFeePercentage;
    }


    //lease stuff
    function startLease(
        uint _startTimestamp,
        uint _duration,
        uint _storagePricePerSecond,
        uint _computePricePerSecond,
        address _provider,
        address _user
    ) public onlyOwner returns (uint leaseId) {
        return _startLease(LeaseRequest({
            startTimestamp: _startTimestamp,
            duration: _duration,
            storagePricePerSecond: _storagePricePerSecond,
            computePricePerSecond: _computePricePerSecond,
            provider: _provider,
            user: _user
        }));
    }

    function _startLease(LeaseRequest memory request) internal returns (uint leaseId) {
        require(request.startTimestamp > 0, "invalid start timestamp");
        require(request.startTimestamp <= block.timestamp, "start in future");
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
       
        // Spend bonuses first, then withdrawable funds, and preserve the split for refunds.
        (uint cashUsed, uint bonusUsed) = wallet.debitForLease(request.user, totalAmount);
        frozenFunds[leaseCount] = totalAmount;
        frozenCashFunds[leaseCount] = cashUsed;
        frozenBonusFunds[leaseCount] = bonusUsed;
        
        leaseId = leaseCount;
        leaseCount++;
        
        leases[leaseId] = Lease({
            user: request.user,
            provider: request.provider,
            startTime: request.startTimestamp,
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
        leaseActivationTime[leaseId] = block.timestamp;
        leaseSettlements[leaseId].lastSettledAt = block.timestamp;
        userActiveLeases[request.user].push(leaseId);
        emit LeaseStarted(
            leaseId,
            leases[leaseId].user,
            leases[leaseId].provider,
            leases[leaseId].startTime,
            leaseActivationTime[leaseId],
            leases[leaseId].duration,
            frozenFunds[leaseId]
        );
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
        
        uint leaseEnd = lease.startTime + lease.duration;
        uint pauseEnd = block.timestamp < leaseEnd ? block.timestamp : leaseEnd;
        uint pauseDuration = pauseEnd > lease.pausedAt ? pauseEnd - lease.pausedAt : 0;
        lease.pausedDuration += pauseDuration;
        lease.pausedAt = 0; // Reset last paused time
        lease.paused = false;
        
        emit LeaseResumed(_leaseId);
    }

    function completeLease(uint _leaseId) external onlyOwner  {
        Lease storage lease = leases[_leaseId];
        require(lease.active, "Lease is not active");
        require(!lease.completed, "Lease already completed");

        // Final settlement is always allowed, even if less than one day passed.
        _settleLease(_leaseId);

        wallet.refundLeaseBalance(
            lease.user,
            frozenCashFunds[_leaseId],
            frozenBonusFunds[_leaseId]
        );

        delete frozenFunds[_leaseId];
        delete frozenCashFunds[_leaseId];
        delete frozenBonusFunds[_leaseId];
        
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


    function settleLease(uint _leaseId) external onlySettlementOperator {
        require(isSettlementDue(_leaseId), "settlement not due");
        _settleLease(_leaseId);
    }


    function settleLeases(uint[] calldata leaseIds) external onlySettlementOperator {
        require(leaseIds.length > 0, "empty batch");
        require(leaseIds.length <= MAX_SETTLEMENT_BATCH_SIZE, "batch too large");

        for (uint i = 0; i < leaseIds.length; i++) {
            require(isSettlementDue(leaseIds[i]), "settlement not due");
            _settleLease(leaseIds[i]);
        }
    }


    function isSettlementDue(uint _leaseId) public view returns (bool) {
        Lease storage lease = leases[_leaseId];
        if (!lease.active || lease.completed) {
            return false;
        }

        LeaseSettlement storage settlement = leaseSettlements[_leaseId];
        bool intervalPassed = block.timestamp >= settlement.lastSettledAt + SETTLEMENT_INTERVAL;
        bool leaseEnded = block.timestamp >= lease.startTime + lease.duration;

        if (!intervalPassed && !leaseEnded) {
            return false;
        }

        (uint providerEarned, uint feeEarned, ) = settlementEntitlement(_leaseId);
        return providerEarned > settlement.providerPaid || feeEarned > settlement.feePaid;
    }


    function settlementEntitlement(uint _leaseId)
        public
        view
        returns (uint providerEarned, uint feeEarned, uint referralEarned)
    {
        Lease storage lease = leases[_leaseId];
        require(lease.startTime > 0, "Lease not started");

        (uint actualStorageCost, uint actualComputeCost) = calculateActualCost(_leaseId);
        uint actualTotalCost = actualStorageCost + actualComputeCost;

        feeEarned = (actualTotalCost * lease.leaseFeePercentage) / 100;
        providerEarned = actualTotalCost - feeEarned;

        LeaseReferralInfo storage referralInfo = leaseReferralInfo[_leaseId];
        if (referralInfo.referrer != address(0) && referralInfo.referralShareBps > 0) {
            referralEarned = (feeEarned * referralInfo.referralShareBps) / 10_000;
        }
    }


    function _settleLease(uint _leaseId) internal {
        Lease storage lease = leases[_leaseId];
        require(lease.active, "Lease is not active");
        require(!lease.completed, "Lease already completed");

        LeaseSettlement storage settlement = leaseSettlements[_leaseId];
        (uint providerEarned, uint feeEarned, uint referralEarned) = settlementEntitlement(_leaseId);

        uint providerAmount = providerEarned - settlement.providerPaid;
        uint platformFee = feeEarned - settlement.feePaid;
        uint referralAmount = referralEarned - settlement.referralPaid;
        uint totalAmount = providerAmount + platformFee;

        settlement.providerPaid = providerEarned;
        settlement.feePaid = feeEarned;
        settlement.referralPaid = referralEarned;
        settlement.lastSettledAt = block.timestamp;

        if (totalAmount == 0) {
            return;
        }

        _consumeFrozenFunds(_leaseId, totalAmount);

        LeaseReferralInfo storage referralInfo = leaseReferralInfo[_leaseId];
        if (referralAmount > 0) {
            wallet.creditBalance(referralInfo.referrer, referralAmount);
        }
        if (platformFee > referralAmount) {
            wallet.creditBalance(treasury, platformFee - referralAmount);
        }
        if (providerAmount > 0) {
            wallet.creditBalance(lease.provider, providerAmount);
        }

        emit LeaseSettled(
            _leaseId,
            providerAmount,
            platformFee,
            referralAmount,
            block.timestamp
        );
    }


    function _consumeFrozenFunds(uint _leaseId, uint amount) internal {
        uint bonusAmount = amount < frozenBonusFunds[_leaseId]
            ? amount
            : frozenBonusFunds[_leaseId];
        uint cashAmount = amount - bonusAmount;

        frozenBonusFunds[_leaseId] -= bonusAmount;
        frozenCashFunds[_leaseId] -= cashAmount;
        frozenFunds[_leaseId] -= amount;
    }


    function calculateActualCost(uint _leaseId)
    public
    view
    returns (uint actualStorageCost, uint actualComputeCost)
{
    Lease storage lease = leases[_leaseId];

    require(lease.startTime > 0, "Lease not started");

    uint leaseEnd = lease.startTime + lease.duration;
    uint calculationTime = block.timestamp < leaseEnd ? block.timestamp : leaseEnd;
    uint storageDuration = calculationTime - lease.startTime;
    uint activationTime = leaseActivationTime[_leaseId];
    uint computeDuration = calculationTime > activationTime
        ? calculationTime - activationTime
        : 0;

    uint totalPaused = lease.pausedDuration;

    if (lease.paused) {
        uint pauseEnd = calculationTime;
        if (pauseEnd > lease.pausedAt) {
            totalPaused += pauseEnd - lease.pausedAt;
        }
    }

    if (totalPaused > computeDuration) {
        totalPaused = computeDuration;
    }

    uint activeDuration = computeDuration - totalPaused;

    actualStorageCost = storageDuration * lease.storagePricePerSecond;
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

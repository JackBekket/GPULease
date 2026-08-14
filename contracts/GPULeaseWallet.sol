// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

contract GPULeaseWallet is Ownable {
    using SafeERC20 for IERC20;

    uint256 public constant MAX_BONUS_BATCH_SIZE = 100;

    IERC20 public credit;
    address public leaseManager;

    mapping(address => uint256) public balances;
    mapping(address => uint256) public bonusBalances;
    uint256 public bonusReserve;

    event Deposit(address indexed user, uint256 amount);
    event Withdraw(address indexed user, uint256 amount);
    event LeaseManagerUpdated(address indexed previousManager, address indexed newManager);
    event BonusPoolFunded(address indexed funder, uint256 amount);
    event BonusGranted(address indexed user, uint256 amount);
    event BonusRevoked(address indexed user, uint256 amount);
    event BonusReserveWithdrawn(address indexed recipient, uint256 amount);
    event LeaseBalanceDebited(
        address indexed user,
        uint256 cashAmount,
        uint256 bonusAmount
    );
    event LeaseBalanceRefunded(
        address indexed user,
        uint256 cashAmount,
        uint256 bonusAmount
    );

    modifier onlyLeaseManager() {
        require(msg.sender == leaseManager, "not lease manager");
        _;
    }

    constructor(address credit_) Ownable(msg.sender) {
        credit = IERC20(credit_);
    }

    function setLeaseManager(address newLeaseManager) external onlyOwner {
        require(newLeaseManager != address(0), "zero lease manager");
        emit LeaseManagerUpdated(leaseManager, newLeaseManager);
        leaseManager = newLeaseManager;
    }

    function deposit(uint256 amount) external {
        credit.safeTransferFrom(msg.sender, address(this), amount);
        balances[msg.sender] += amount;
        emit Deposit(msg.sender, amount);
    }

    function depositFor(address beneficiary, uint256 amount) external {
        require(beneficiary != address(0), "zero beneficiary");
        credit.safeTransferFrom(msg.sender, address(this), amount);
        balances[beneficiary] += amount;
        emit Deposit(beneficiary, amount);
    }

    function withdraw(uint256 amount) external {
        require(balances[msg.sender] >= amount, "insufficient balance");
        balances[msg.sender] -= amount;
        credit.safeTransfer(msg.sender, amount);
        emit Withdraw(msg.sender, amount);
    }

    function userBalance(address user) public view returns (uint256) {
        return spendableBalance(user);
    }

    function withdrawableBalance(address user) public view returns (uint256) {
        return balances[user];
    }

    function bonusBalance(address user) public view returns (uint256) {
        return bonusBalances[user];
    }

    function spendableBalance(address user) public view returns (uint256) {
        return balances[user] + bonusBalances[user];
    }

    function fundBonusPool(uint256 amount) external {
        require(amount > 0, "zero amount");
        credit.safeTransferFrom(msg.sender, address(this), amount);
        bonusReserve += amount;
        emit BonusPoolFunded(msg.sender, amount);
    }

    function grantBonus(address user, uint256 amount) external onlyOwner {
        _grantBonus(user, amount);
    }

    function grantBonuses(
        address[] calldata users,
        uint256[] calldata amounts
    ) external onlyOwner {
        require(users.length > 0, "empty batch");
        require(users.length == amounts.length, "length mismatch");
        require(users.length <= MAX_BONUS_BATCH_SIZE, "batch too large");

        uint256 totalAmount;
        for (uint256 i = 0; i < users.length; i++) {
            require(users[i] != address(0), "zero user");
            require(amounts[i] > 0, "zero amount");
            totalAmount += amounts[i];
        }
        require(bonusReserve >= totalAmount, "insufficient bonus reserve");

        bonusReserve -= totalAmount;
        for (uint256 i = 0; i < users.length; i++) {
            bonusBalances[users[i]] += amounts[i];
            emit BonusGranted(users[i], amounts[i]);
        }
    }

    function revokeBonus(address user, uint256 amount) external onlyOwner {
        require(user != address(0), "zero user");
        require(amount > 0, "zero amount");
        require(bonusBalances[user] >= amount, "insufficient bonus balance");

        bonusBalances[user] -= amount;
        bonusReserve += amount;
        emit BonusRevoked(user, amount);
    }

    function withdrawBonusReserve(
        address recipient,
        uint256 amount
    ) external onlyOwner {
        require(recipient != address(0), "zero recipient");
        require(amount > 0, "zero amount");
        require(bonusReserve >= amount, "insufficient bonus reserve");

        bonusReserve -= amount;
        credit.safeTransfer(recipient, amount);
        emit BonusReserveWithdrawn(recipient, amount);
    }

    function debitForLease(
        address user,
        uint256 amount
    ) external onlyLeaseManager returns (uint256 cashUsed, uint256 bonusUsed) {
        require(spendableBalance(user) >= amount, "Insufficient token balance");

        bonusUsed = amount < bonusBalances[user]
            ? amount
            : bonusBalances[user];
        cashUsed = amount - bonusUsed;

        bonusBalances[user] -= bonusUsed;
        balances[user] -= cashUsed;
        emit LeaseBalanceDebited(user, cashUsed, bonusUsed);
    }

    function refundLeaseBalance(
        address user,
        uint256 cashAmount,
        uint256 bonusAmount
    ) external onlyLeaseManager {
        balances[user] += cashAmount;
        bonusBalances[user] += bonusAmount;
        emit LeaseBalanceRefunded(user, cashAmount, bonusAmount);
    }

    function debitBalance(address user, uint256 amount) external onlyLeaseManager {
        require(balances[user] >= amount, "Insufficient token balance");
        balances[user] -= amount;
    }

    function creditBalance(address user, uint256 amount) external onlyLeaseManager {
        balances[user] += amount;
    }

    function moveBalance(address from, address to) external onlyLeaseManager returns (uint256 amount) {
        amount = balances[from];
        balances[to] += amount;
        balances[from] = 0;
    }

    function _grantBonus(address user, uint256 amount) internal {
        require(user != address(0), "zero user");
        require(amount > 0, "zero amount");
        require(bonusReserve >= amount, "insufficient bonus reserve");

        bonusReserve -= amount;
        bonusBalances[user] += amount;
        emit BonusGranted(user, amount);
    }
}

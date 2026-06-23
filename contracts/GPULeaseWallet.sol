// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

contract GPULeaseWallet is Ownable {
    using SafeERC20 for IERC20;

    IERC20 public credit;
    address public leaseManager;

    mapping(address => uint256) public balances;

    event Deposit(address indexed user, uint256 amount);
    event Withdraw(address indexed user, uint256 amount);
    event LeaseManagerUpdated(address indexed previousManager, address indexed newManager);

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
        return balances[user];
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
}

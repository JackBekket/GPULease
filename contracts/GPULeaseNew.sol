// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract GPULease is Ownable {
    using SafeERC20 for IERC20;
    IERC20 public credit; 
    address public treasury;

    mapping(address => uint256) public balances;


    event Deposit(address indexed user, uint256 amount);
    event Withdraw(address indexed user, uint256 amount);
    

    constructor(address credit_, address treasury_) Ownable(msg.sender) {
        credit = IERC20(credit_);
        treasury = treasury_;
    }

   //
    function deposit(uint256 amount) external {
        credit.safeTransferFrom(msg.sender, address(this), amount);
        balances[msg.sender] += amount;
        emit Deposit(msg.sender, amount);
    }

    function withdraw(uint256 amount) external {
        require(balances[msg.sender] >= amount, "insufficient balance");
        credit.safeTransfer(msg.sender, amount);
        balances[msg.sender] -= amount;
        emit Withdraw(msg.sender, amount);
    }
   
    function setTreasury(address newTreasury) external onlyOwner {
        require(newTreasury != address(0), "zero treasury");
        treasury = newTreasury;
    }

    function userBalance(address user) public view returns (uint256) {
    return balances[user];
    }
}
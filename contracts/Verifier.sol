// SPDX-License-Identifier: MIT
pragma solidity ^0.8.31;

contract Verifier {
    event WalletVerified(address indexed wallet, string userId);

    function verify(string calldata userId) external {
        emit WalletVerified(msg.sender, userId);
    }
}

// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/access/Ownable.sol";

contract GPULeaseReferral is Ownable {
    uint private constant BPS = 10_000;

    mapping(address => address) public referrerOf;

    uint public referralShareBps = 5_000;

    event ReferrerUpdated(address indexed user, address indexed referrer);
    event ReferrerCleared(address indexed user, address indexed previousReferrer);
    event ReferralShareUpdated(uint previousShareBps, uint newShareBps);

    constructor() Ownable(msg.sender) {}

    function setReferrer(address user, address referrer) external onlyOwner {
        require(user != address(0), "invalid user");
        require(referrer != address(0), "invalid referrer");
        require(user != referrer, "self referral");

        referrerOf[user] = referrer;
        emit ReferrerUpdated(user, referrer);
    }

    function clearReferrer(address user) external onlyOwner {
        require(user != address(0), "invalid user");

        address previousReferrer = referrerOf[user];
        delete referrerOf[user];

        emit ReferrerCleared(user, previousReferrer);
    }

    function setReferralShareBps(uint _shareBps) external onlyOwner {
        require(_shareBps <= BPS, "Share too high");

        emit ReferralShareUpdated(referralShareBps, _shareBps);
        referralShareBps = _shareBps;
    }
}

# GPULease - Ethereum GPU Leasing Platform

Current Base addresses (updated 2026-08-14T11:10:21Z):

- GPULease: `0x1350d6D31dc4c8B1314aF51d99e61cF0E3da938f`
- GPULeaseWallet: `0xf6d56d64938b65c6Ad58cFD447Cd1d74b39eEeF2`
- CampaignFactory: `0x40C99f349A8cB30d452c9cdc7808221D57427851`
- Campaign implementation: `0x0b085C432C4dDAC1C8EF2D4b3C259F5ace69cbd4`

A smart contract for leasing GPU computing resources with flexible pricing and pause/resume functionality.

## Overview

GPULease is a decentralized smart contract that enables users to lease GPU compute and storage resources from providers. The system handles payment processing, platform fee collection, and provides mechanisms for pausing and resuming leases.

## Features

- **Flexible Pricing**: Separate pricing for storage and computation per second
- **Platform Fee System**: 10% default platform fee (configurable globally and per user)
- **Daily Settlement**: Providers receive accrued income without waiting for lease completion
- **Bonus USDC Balance**: Owner grants backed, non-withdrawable credits usable for leases
- **Crowdfunding Fee**: Campaign targets include a fixed 5% fee on top; creators receive their full requested amount
- **Pause/Resume Functionality**: Leases can be paused during execution with accurate cost calculation
- **Lease Cancellation**: Within 5 minutes of creation
- **Refund Mechanism**: Automatic refunds for unused time
- **Access Control**: Role-based permissions using OpenZeppelin AccessControl

## Contract Structure

### Core Data Structures

**Lease Struct**
```solidity
struct Lease {
    address user;                    // The lease requester
    address provider;                // The GPU provider
    uint startTime;                  // Supplied start of storage billing
    uint duration;                   // Total lease duration in seconds
    uint storagePricePerSecond;      // Price per second for storage
    uint computePricePerSecond;      // Price per second for computation
    uint totalAmount;                // Total amount to be paid
    bool active;                     // Whether the lease is currently active
    bool completed;                  // Whether the lease has been completed
    bool paused;                     // Whether the lease is currently paused
    uint lastPausedTime;             // Time when lease was last paused
    uint pausedDuration;             // Cumulative duration of pauses in seconds
}
```

### Key Functions

**Deposit & Withdraw**
- `GPULeaseWallet.deposit(amount)`: Add tokens to the withdrawable balance
- `GPULeaseWallet.withdraw(amount)`: Withdraw cash; bonuses cannot be withdrawn
- `GPULeaseWallet.spendableBalance(user)`: Cash plus bonus available for leases

**Lease Management**
- `startLease(startTimestamp, duration, storagePricePerSecond, computePricePerSecond, provider, user)`: Start a lease; storage is billed from startTimestamp and compute from on-chain activation
- `pauseLease(leaseId)`: Pause an active lease (only the user or provider can call)
- `resumeLease(leaseId)`: Resume a paused lease (only the user or provider can call)
- `completeLease(leaseId)`: Complete an active lease and settle payments
- `cancelLease(leaseId)`: Cancel a lease within 5 minutes
- `settleLease(leaseId)` / `settleLeases(leaseIds)`: Credit accrued daily income

**Bonus Management (wallet owner)**
- `fundBonusPool(amount)`: Back the bonus reserve with USDC
- `grantBonus(user, amount)` / `grantBonuses(users, amounts)`: Grant credits
- `revokeBonus(user, amount)`: Return unused credits to the reserve

**Utility Functions**
- `getLeaseStatus(leaseId)`: Get detailed information about a specific lease

## Security Features

- **ReentrancyGuard**: Prevents reentrant calls during critical operations
- **Access Control**: Role-based permissions for administrative functions
- **Platform Fee Management**: Admin can update platform fee percentage

## Events

The contract emits the following events:
- `LeaseStarted`: When a lease is created
- `LeaseCompleted`: When a lease is completed successfully 
- `LeaseCancelled`: When a lease is cancelled
- `PaymentReceived`: When payment is received for a lease
- `PlatformFeeCollected`: When platform fee is collected
- `UserDeposited`: When a user deposits tokens
- `UserWithdrawn`: When a user withdraws tokens
- `LeasePaused`: When a lease is paused
- `LeaseResumed`: When a lease is resumed

## Usage Flow

1. **Fund Balance**: Users deposit USDC, receive a bonus, or use both
2. **Start Lease**: Choose duration and pricing, then start the lease  
3. **Pause/Resume** (Optional): During execution, pause or resume the lease
4. **Complete/Cancellation**: Either complete the lease or cancel it within 5 minutes

## Contract Parameters

- `platformFeePercentage`: Default 10% platform fee (modifiable by owner)
- `SETTLEMENT_INTERVAL`: 1 day
- All prices are expressed in smallest token units (wei)

## Access Control

The contracts use OpenZeppelin `Ownable` and `ReentrancyGuard`:
- Owner: starts leases and manages fees, referrals, operators, and wallet bonuses
- Settlement operator: can execute only daily settlements
- Lease participants: Users and providers can manage their own leases

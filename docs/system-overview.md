# GPULease System Overview

Current deployment updated at 2026-08-14T11:10:21Z. Base `GPULease`:
`0x1350d6D31dc4c8B1314aF51d99e61cF0E3da938f`, bonus wallet:
`0xf6d56d64938b65c6Ad58cFD447Cd1d74b39eEeF2`, fee-enabled CampaignFactory:
`0x40C99f349A8cB30d452c9cdc7808221D57427851`.

This document describes the current GPULease architecture, the money flows, the upgrade model, and the public/external contract functions.

## High-Level Architecture

The system is split into three main GPULease modules:

- `GPULeaseWallet`: permanent wallet/storage contract for ERC20 funds and user balances.
- `GPULease`: replaceable lease/business-logic contract.
- `GPULeaseReferral`: replaceable referral-rules contract.

Campaign contracts interact with `GPULeaseWallet` directly for deposits. `GPULease` never exposes user deposit, withdraw, balance, or token getters; all user money logic lives in `GPULeaseWallet`.

```text
User / Campaign
  -> GPULeaseWallet.deposit / depositFor / withdraw
  -> GPULeaseWallet stores tokens and balances

Owner / backend
  -> GPULease.startLease / pauseLease / resumeLease / completeLease
  -> GPULeaseWallet debit/credit/move balances
  -> GPULeaseReferral lookup referrer and referral split
```

## Upgrade Model

`GPULeaseWallet` owns the token balances and exposes manager-only accounting functions.

To replace `GPULease`:

1. Deploy a new `GPULease(existingWallet, treasury)`.
2. Call `wallet.setLeaseManager(newGPULease)`.
3. Optional: call `newGPULease.setReferralManager(existingOrNewReferral)`.

User balances remain in that `GPULeaseWallet`.

Important limitation: a wallet supports only one `leaseManager`, while active leases and `frozenFunds` live in a specific `GPULease`. With active leases, deploy a parallel wallet/lease stack or implement migration first. The current Base deployment uses the parallel-stack approach.

To replace referral logic:

1. Deploy a new referral manager contract implementing:
   - `referrerOf(address) view returns (address)`
   - `referralShareBps() view returns (uint)`
2. Call `gpuLease.setReferralManager(newReferralManager)`.

Existing leases keep the referral data captured when they were started.

## GPULeaseWallet

File: `contracts/GPULeaseWallet.sol`

Purpose:

- Holds the ERC20 `credit` token.
- Stores withdrawable balances in `balances` and non-withdrawable credits in `bonusBalances`.
- Allows normal users to deposit and withdraw directly.
- Allows the current `leaseManager` to mutate balances for lease settlement.

State:

- `credit`: ERC20 token used for payments.
- `leaseManager`: authorized lease contract.
- `balances[user]`: internal user balance.
- `bonusBalances[user]`: non-withdrawable balance usable for leases.
- `bonusReserve`: USDC-backed reserve available for bonus grants.

Events:

- `Deposit(user, amount)`
- `Withdraw(user, amount)`
- `LeaseManagerUpdated(previousManager, newManager)`
- `BonusPoolFunded`, `BonusGranted`, `BonusRevoked`
- `LeaseBalanceDebited`, `LeaseBalanceRefunded`

Functions:

- `constructor(address credit_)`
  - Sets the ERC20 token.

- `setLeaseManager(address newLeaseManager) onlyOwner`
  - Sets the only contract allowed to call manager-only balance functions.
  - Rejects zero address.

- `deposit(uint256 amount)`
  - Transfers tokens from caller into the wallet.
  - Credits caller's internal balance.

- `depositFor(address beneficiary, uint256 amount)`
  - Transfers tokens from caller into the wallet.
  - Credits `beneficiary`.
  - This is useful for campaigns or backend flows that fund another user's balance.

- `withdraw(uint256 amount)`
  - Debits only caller's withdrawable balance; bonuses cannot be withdrawn.
  - Transfers tokens from wallet to caller.

- `userBalance(address user) view returns (uint256)`
  - Returns total spendable balance: withdrawable plus bonus.

- `withdrawableBalance`, `bonusBalance`, `spendableBalance`
  - Return separate and combined balances.

- `fundBonusPool`, `grantBonus`, `grantBonuses`, `revokeBonus`
  - Manage the USDC-backed bonus reserve and user bonus balances.

- `debitForLease` / `refundLeaseBalance`
  - Spend bonuses before cash and preserve the balance type for unused refunds.

- `debitBalance(address user, uint256 amount) onlyLeaseManager`
  - Debits internal balance without transferring tokens out.
  - Used when starting a lease and freezing funds.

- `creditBalance(address user, uint256 amount) onlyLeaseManager`
  - Credits internal balance.
  - Used when settling provider, treasury, referrer, and refunds.

- `moveBalance(address from, address to) onlyLeaseManager returns (uint256 amount)`
  - Moves full internal balance from one address to another.
  - Used when changing treasury.

## GPULeaseReferral

File: `contracts/GPULeaseReferral.sol`

Purpose:

- Stores user -> referrer mapping.
- Stores the referrer share of the platform fee.
- Can be replaced independently from wallet and lease logic.

State:

- `referrerOf[user]`: referrer address for a user.
- `referralShareBps`: referrer share in basis points, default `5000`, meaning 50% of the platform fee.

Events:

- `ReferrerUpdated(user, referrer)`
- `ReferrerCleared(user, previousReferrer)`
- `ReferralShareUpdated(previousShareBps, newShareBps)`

Functions:

- `constructor()`
  - Sets owner.

- `setReferrer(address user, address referrer) onlyOwner`
  - Assigns a referrer.
  - Rejects zero user, zero referrer, and self-referral.

- `clearReferrer(address user) onlyOwner`
  - Removes a user's referrer.

- `setReferralShareBps(uint shareBps) onlyOwner`
  - Sets the referrer share of the platform fee.
  - `10000` means 100% of the fee.
  - Default is `5000`, meaning fee is split 50/50 between referrer and treasury.

## GPULease

File: `contracts/GPULease.sol`

Purpose:

- Main lease lifecycle contract.
- Delegates all wallet balance storage to `GPULeaseWallet`.
- Reads referral rules from `GPULeaseReferral`.
- Keeps lease records, frozen funds, per-user fee overrides, and active lease indexes.

State:

- `wallet`: immutable `GPULeaseWallet`.
- `referralManager`: replaceable referral manager.
- `treasury`: platform treasury balance address.
- `platformFeePercentage`: default user fee, currently `10`.
- `userFeePercentage[user]`: custom user fee value.
- `hasCustomUserFee[user]`: whether a custom fee is enabled.
- `leases[leaseId]`: lease data.
- `leaseReferralInfo[leaseId]`: referrer and referral share captured at lease start.
- `frozenFunds[leaseId]`: funds reserved for a lease.
- `userActiveLeases[user]`: active lease IDs for each user.

Events:

- `PlatformFeeUpdated(previousFeePercentage, newFeePercentage)`
- `UserFeeUpdated(user, feePercentage)`
- `UserFeeCleared(user)`
- `ReferralManagerUpdated(previousManager, newManager)`
- `LeaseStarted(leaseId, user, provider, startTimestamp, activationTimestamp, duration, amount)`
- `LeaseCompleted(leaseId)`
- `LeasePaused(leaseId)`
- `LeaseResumed(leaseId)`

Admin functions:

- `setPlatformFee(uint feePercentage) onlyOwner`
  - Sets default platform fee for users without a custom fee.
  - Rejects values above 100.

- `setUserFee(address user, uint feePercentage) onlyOwner`
  - Sets a custom fee for a specific user.
  - Rejects zero user and values above 100.

- `clearUserFee(address user) onlyOwner`
  - Removes a user's custom fee so they fall back to default fee.

- `setReferralManager(address newReferralManager) onlyOwner`
  - Updates referral manager address.
  - May be set to zero to disable referral lookup for future leases.

- `setTreasury(address newTreasury) onlyOwner`
  - Moves the old treasury's internal balance to the new treasury address.
  - Updates `treasury`.

Fee functions:

- `feePercentageForUser(address user) view returns (uint)`
  - Returns custom fee when set.
  - Otherwise returns default `platformFeePercentage`.

Lease lifecycle functions:

- `startLease(startTimestamp, duration, storagePricePerSecond, computePricePerSecond, provider, user) onlyOwner returns (leaseId)`
  - Requires a positive `startTimestamp` that is not in the future.
  - Requires positive duration.
  - Requires at least one price to be positive.
  - Requires non-zero user and provider.
  - Calculates max cost:
    - `base = duration * storagePricePerSecond + duration * computePricePerSecond`
    - `fee = base * feePercentageForUser(user) / 100`
    - `frozen = base + fee`
  - Captures current referrer and referral share for the lease.
  - Debits user wallet balance and stores frozen funds.
  - Charges storage only from `startTimestamp` until on-chain activation.
  - Stores activation time in `leaseActivationTime[leaseId]`.
  - Charges storage and compute after activation; pauses reduce compute time.
  - Counts total `duration` from `startTimestamp`.
  - Creates an active lease.

- `pauseLease(uint leaseId) onlyOwner`
  - Marks active lease as paused and stores pause start timestamp.

- `resumeLease(uint leaseId) onlyOwner`
  - Ends pause.
  - Adds pause duration to cumulative `pausedDuration`.

- `completeLease(uint leaseId) onlyOwner`
  - Calculates actual storage and compute costs.
  - Storage cost uses elapsed time.
  - Compute cost uses active time excluding pauses.
  - Caps elapsed time at declared duration.
  - Calculates platform fee from actual cost and lease's captured fee percentage.
  - If referral exists:
    - `referralAmount = platformFee * referralShareBps / 10000`
    - referrer receives `referralAmount`
    - treasury receives `platformFee - referralAmount`
  - Provider receives `actualTotalCost - platformFee`.
  - User receives any unused frozen funds.
  - Lease is marked completed and removed from user's active lease list.

- `getUserFrozenFunds(address user) view returns (FrozenFundsInfo[])`
  - Returns active lease IDs and frozen amounts for a user.

Internal functions:

- `_startLease(LeaseRequest request)`
  - Internal implementation for `startLease`.

- `calculateActualCost(uint leaseId) public view`
  - Calculates actual storage and compute costs.

- `_referralInfoForLease(address user) internal view`
  - Reads referral manager and returns referrer/share for a new lease.

## Fee And Referral Examples

Without referrer:

```text
actualTotalCost = 1000
user fee = 10%
platformFee = 100
treasury = 100
referrer = 0
provider = 900
```

With referrer and default 50/50 referral split:

```text
actualTotalCost = 1000
user fee = 10%
platformFee = 100
referrer = 50
treasury = 50
provider = 900
```

With referrer, custom user fee 7%, and default 50/50 referral split:

```text
actualTotalCost = 1000
user fee = 7%
platformFee = 70
referrer = 35
treasury = 35
provider = 930
```

The user fee and referral info are captured when a lease starts. Later changes do not affect existing leases.

## Campaign System

Files:

- `contracts/Campaign.sol`
- `contracts/CampaignFactory.sol`
- `contracts/CampaignMetadataRenderer.sol`

Campaigns currently use this wallet interface:

```solidity
interface IGPULeaseWallet {
    function depositFor(address beneficiary, uint256 amount) external;
}
```

When a campaign succeeds:

1. A fixed 5% fee is added on top of the creator amount; a 100 USDC request displays a 105 USDC target.
2. Campaign approves the net creator amount for `GPULeaseWallet`.
3. Campaign calls `gpuLeaseWallet.depositFor(owner(), creatorTargetAmount)`.
4. The withdrawable wallet balance is credited to the campaign owner.
5. `feeAmount` is transferred directly to the factory's immutable `feeRecipient`.

The fee is charged only on `SUCCESS`. Failed campaigns refund complete donations. Gross-target overfunding is rejected, so the creator receives exactly the requested amount and treasury receives exactly 5% on top.

NFT reward metadata is rendered through `CampaignMetadataRenderer`. The campaign NFT `tokenURI()` returns base64 JSON with:

- `name`: campaign name plus backer level, for example `Medical LLM crowdfunding - Lead Backer`.
- `description`: reward description.
- `image`: base64-encoded SVG text card, not a bitmap image. This lets wallets render visible text in the NFT preview without IPFS or external image hosting.
- `attributes`: campaign name, backer level, and campaign ID.

### LLMFundraising

Main public/external functions:

- `deadline()`
- `isExpired()`
- `isTargetReached()`
- `checkConditions()`
- `donorShareBps(address donor)`
- `gradeForDonation(address donor)`
- `donorsCount()`
- `donorAt(uint index)`
- `donors()`
- `donorsSlice(uint offset, uint limit)`
- `donorInfo(address donor)`
- `claimBackerReward()`
- `mintBackerRewards(uint offset, uint limit)`
- `tokenURI(uint tokenId)`
- `donate(uint amount)`
- `checkState()`
- `refund()`

Main flow:

- Donors send USDC with `donate`.
- Donations are tracked and donor grade is updated.
- `creatorTargetAmount` is the creator's net amount, `feeAmount` is 5%, and `targetAmount` is their displayed gross sum.
- When the gross target is reached, campaign state becomes `SUCCESS`.
- A successful campaign credits the creator through `GPULeaseWallet.depositFor` and transfers the fee to treasury.
- Failed campaign allows donors to refund.
- Successful donors can mint backer reward NFTs.

### LLMFundraisingFactory

Main public/external functions:

- `createCampaign(targetAmount, duration, startTimestamp, templateId, campaignName)`; input `targetAmount` is the creator's net amount
- `campaignsCount()`
- `campaignsByCreator(address creator)`
- `registerParticipant(address participant)`
- `campaignsByParticipant(address participant)`
- `participantCampaignsCount(address participant)`
- `participantCampaignAt(address participant, uint index)`
- `hasParticipatedInCampaign(address participant, address campaign)`

Main flow:

- Factory deploys one `LLMFundraising` implementation in its constructor.
- Each new campaign is an OpenZeppelin `Clones` minimal proxy pointing to that implementation.
- After cloning, factory calls `initialize(...)` on the campaign clone.
- The implementation contract is locked against direct initialization, and campaign clones expose constant ERC721 `name()` and `symbol()`.
- Tracks campaigns by ID and creator.
- Campaign contracts register participants back into the factory.

Constructor arguments:

- `usdc`
- `gpuLeaseWallet`
- `metadataRenderer`
- `feeRecipient`

### CampaignMetadataRenderer

Main public/external functions:

- `tokenURI(string campaignName, uint256 campaignId, uint8 grade)`

Main flow:

- Converts backer grade to a human-readable level.
- Builds ERC721 JSON metadata.
- Adds a base64-encoded SVG text card in `image` with campaign name and backer level.
- Returns a `data:application/json;base64,...` URI.

## Verifier

File: `contracts/Verifier.sol`

Functions:

- `verify(string userId)`
  - Emits `WalletVerified(msg.sender, userId)`.

## MockERC20

File: `contracts/MockERC20.sol`

Test helper only.

Functions:

- `mint(address to, uint256 amount)`
  - Mints tokens for tests.

## Deployment Flow

Current `scripts/deploy.ts` deploys:

1. `GPULeaseWallet(USDC)`
2. `GPULeaseReferral()`
3. `GPULease(wallet, treasury)`
4. `wallet.setLeaseManager(gpuLease)`
5. `gpuLease.setReferralManager(referral)`

Current `scripts/deploy-campaign-factory.ts` deploys:

1. `CampaignMetadataRenderer()`
2. `LLMFundraisingFactory(USDC, GPULeaseWallet, metadataRenderer)`

## Test Status

Verified commands:

```bash
/home/m0rs/.nvm/versions/node/v22.22.1/bin/node node_modules/hardhat/dist/src/cli.js test test/GPULease.ts 'test/GPULease (EdgeCases).ts' test/GasCosts.ts
```

Result:

```text
33 passing
```

Additional NFT metadata renderer test:

```bash
/home/m0rs/.nvm/versions/node/v22.22.1/bin/node node_modules/hardhat/dist/src/cli.js test test/CampaignMetadataRenderer.ts
```

Result:

```text
1 passing
```

Coverage currently included in GPULease tests:

- Wallet deposit and withdraw.
- `depositFor` and zero beneficiary rejection.
- Lease start and frozen funds.
- Pause and resume.
- Completion and settlement.
- Duration edge cases with and without pause.
- Withdraw above balance rejection.
- Default fee change.
- Per-user custom fee set, clear, and owner-only guard.
- Custom fee settlement.
- Fee captured at lease start.
- Referral split of platform fee.
- Referral split with custom user fee 7%.
- Referral rules captured at lease start.
- Referral manager replacement.
- Custom referral share.
- Referral admin guard checks.
- Treasury balance migration.
- Wallet balance persistence after replacing `GPULease`.
- Gas snapshots for key operations.
- Campaign NFT metadata renderer outputs JSON metadata and a base64-encoded SVG text card.

Full test command:

```bash
/home/m0rs/.nvm/versions/node/v22.22.1/bin/node node_modules/hardhat/dist/src/cli.js test
```

Current result:

```text
39 passing
```

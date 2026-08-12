# Backend Integration Guide

Короткий гайд для backend-интеграции с текущей архитектурой GPULease.

Актуальный Base deployment на 2026-08-12T11:45:55Z:

- `GPULease`: `0xCCD732200366886e04F508D12F561ee94Eb03110`
- `GPULeaseWallet`: `0xD4352D14Ba7928f6066dd7ec6031C7c0CCF13340`
- `GPULeaseReferral`: `0x2695d98bF8b233539f5a1Fb823298AA055f2a143`

Предыдущий `GPULease` `0xB6E47Eb260160BD6A18A246CC0b27D9240706401`
выведен из эксплуатации 2026-08-12T11:45:55Z. Новые аренды нужно создавать
только на актуальном адресе.

## Что Хранить В Конфиге

Backend должен хранить адреса:

- `creditToken`: ERC20 токен оплаты, например USDC.
- `gpuLeaseWallet`: контракт денег и балансов.
- `gpuLease`: контракт аренды.
- `gpuLeaseReferral`: контракт рефералок.
- `campaignMetadataRenderer`: renderer NFT metadata для кампаний.
- `campaignFactory`: factory кампаний, если backend работает с crowdfunding.
- `treasury`: адрес treasury-баланса внутри wallet.

Главное правило: **деньги живут только в `GPULeaseWallet`**. Не дергать `GPULease` для deposit/withdraw/balance.

## Главные Изменения С Прошлой Версии

- `GPULease.deposit`, `GPULease.withdraw`, `GPULease.userBalance`, `GPULease.depositFor` удалены.
- `deposit`, `depositFor`, `withdraw`, `userBalance` теперь вызываются только на `GPULeaseWallet`.
- `Campaign` и `CampaignFactory` теперь работают с `GPULeaseWallet`, а не с `GPULease`, для депозитов.
- `CampaignFactory` constructor теперь принимает `metadataRenderer`.
- `CampaignFactory` теперь создает кампании через OpenZeppelin `Clones`: один `LLMFundraising` implementation деплоится при создании factory, а каждая кампания это minimal proxy + `initialize(...)`.
- Reward NFT metadata теперь рендерится через `CampaignMetadataRenderer`.
- NFT metadata содержит `image` как base64-encoded SVG-карточку с текстом: название кампании и уровень бэкерства.
- Реферальная выплата теперь считается как доля от комиссии, а не отдельный процент от стоимости аренды.

## Пользовательские Балансы

Контракт: `GPULeaseWallet`

### Пополнить Свой Баланс

1. User делает approve на `gpuLeaseWallet`.
2. User вызывает `deposit(amount)`.

```ts
await creditToken.connect(user).approve(gpuLeaseWalletAddress, amount);
await gpuLeaseWallet.connect(user).deposit(amount);
```

### Пополнить Баланс Другого Пользователя

Использовать `depositFor(beneficiary, amount)`.

```ts
await creditToken.connect(payer).approve(gpuLeaseWalletAddress, amount);
await gpuLeaseWallet.connect(payer).depositFor(userAddress, amount);
```

Это нужно для backend/campaign flow, когда платит один адрес, а баланс нужно начислить другому.

### Вывести Средства

```ts
await gpuLeaseWallet.connect(user).withdraw(amount);
```

### Проверить Баланс

```ts
const balance = await gpuLeaseWallet.userBalance(userAddress);
```

Также есть public mapping getter:

```ts
const balance = await gpuLeaseWallet.balances(userAddress);
```

## Аренды

Контракт: `GPULease`

Аренды запускает owner/backend-адрес. Пользователь заранее должен иметь баланс в `GPULeaseWallet`.

### Запустить Аренду

```ts
await gpuLease["startLease(uint256,uint256,uint256,uint256,address,address)"](
  startTimestamp,
  duration,
  storagePricePerSecond,
  computePricePerSecond,
  providerAddress,
  userAddress
);
```

`startTimestamp` должен быть положительным Unix timestamp и не может быть в
будущем. От `startTimestamp` до timestamp транзакции `startLease` начисляется
только storage. Timestamp транзакции сохраняется в
`leaseActivationTime[leaseId]`; после него начисляются storage и compute.
`duration` отсчитывается от `startTimestamp`.

Что происходит:

- контракт считает максимальную стоимость аренды;
- добавляет комиссию пользователя;
- списывает сумму с баланса пользователя в `GPULeaseWallet`;
- кладет сумму в `frozenFunds[leaseId]`;
- фиксирует текущую комиссию пользователя;
- фиксирует текущие referral-правила для этого lease.

### Поставить На Паузу

```ts
await gpuLease.pauseLease(leaseId);
```

### Снять С Паузы

```ts
await gpuLease.resumeLease(leaseId);
```

### Завершить Аренду

```ts
await gpuLease.completeLease(leaseId);
```

Что происходит при завершении:

- считается фактическая стоимость storage и compute;
- compute-время уменьшается на паузы;
- storage считается по elapsed time;
- elapsed time не может быть больше duration;
- provider получает `actualTotalCost - platformFee`;
- treasury получает свою часть комиссии;
- referrer получает свою часть комиссии, если он есть;
- пользователь получает неиспользованный остаток frozen funds.

### Получить Frozen Funds Пользователя

```ts
const frozen = await gpuLease.getUserFrozenFunds(userAddress);
```

Возвращает массив:

```ts
[
  { leaseId, amount },
  ...
]
```

## Комиссии

Контракт: `GPULease`

Дефолтная комиссия сейчас `10%`.

### Посмотреть Комиссию Пользователя

```ts
const fee = await gpuLease.feePercentageForUser(userAddress);
```

### Изменить Дефолтную Комиссию

Только owner:

```ts
await gpuLease.setPlatformFee(10);
```

### Выставить Персональную Комиссию Пользователю

Только owner:

```ts
await gpuLease.setUserFee(userAddress, 7);
```

Это значит, что новые аренды пользователя будут использовать `7%`.

Важно: комиссия фиксируется при `startLease`. Если после старта поменять комиссию, уже созданный lease не пересчитается.

### Убрать Персональную Комиссию

```ts
await gpuLease.clearUserFee(userAddress);
```

После этого пользователь снова использует дефолтную комиссию.

## Рефералка

Контракт: `GPULeaseReferral`

Текущая модель:

- у пользователя может быть `referrer`;
- комиссия пользователя делится между treasury и referrer;
- по умолчанию `referralShareBps = 5000`, то есть referrer получает 50% от комиссии.

Пример:

```text
actualTotalCost = 1000
user fee = 7%
platformFee = 70
referralShareBps = 5000
referrer = 35
treasury = 35
```

### Назначить Referrer

Только owner:

```ts
await gpuLeaseReferral.setReferrer(userAddress, referrerAddress);
```

Нельзя:

- zero user;
- zero referrer;
- referrer == user.

### Убрать Referrer

```ts
await gpuLeaseReferral.clearReferrer(userAddress);
```

### Изменить Долю Referrer

```ts
await gpuLeaseReferral.setReferralShareBps(5000);
```

Шкала:

```text
10000 = 100%
5000  = 50%
1000  = 10%
100   = 1%
1     = 0.01%
```

Важно: referrer и referral share фиксируются при `startLease`. Если поменять referrer/share после старта, старый lease не изменится.

## Кампании

Контракты:

- `LLMFundraisingFactory`
- `LLMFundraising`
- `CampaignMetadataRenderer`

### Деплой Factory

Constructor:

```solidity
LLMFundraisingFactory(
  address usdc,
  address gpuLeaseWallet,
  address metadataRenderer
)
```

Backend/deploy script должен сначала задеплоить `CampaignMetadataRenderer`, потом передать его адрес в factory.

Factory внутри constructor деплоит один `LLMFundraising` implementation. При `createCampaign(...)` она создает minimal proxy clone и вызывает на нем `initialize(...)`. Implementation заблокирован от прямого initialize, а clones возвращают ERC721 `name()` и `symbol()` через константный override. Для backend flow почти не меняется: адрес кампании по-прежнему берется из `campaignById(campaignId)` или события `CampaignCreated`.

### Создать Кампанию

```ts
await campaignFactory.createCampaign(
  targetAmount,
  duration,
  startTimestamp,
  templateId,
  campaignName
);
```

Factory вернет/запишет campaign address:

```ts
const campaign = await campaignFactory.campaignById(campaignId);
```

### Donation Flow

Donor:

```ts
await usdc.connect(donor).approve(campaignAddress, amount);
await campaign.connect(donor).donate(amount);
```

Когда target достигнут:

- campaign становится `SUCCESS`;
- campaign переводит собранные средства в `GPULeaseWallet.deposit`;
- баланс начисляется на адрес campaign-контракта.

Если нужно начислять средства не campaign-контракту, а owner/user, нужно менять campaign integration на `GPULeaseWallet.depositFor(beneficiary, amount)`.

## Reward NFT

Контракт campaign является ERC721.

### Mint Reward

После успешной кампании donor может вызвать:

```ts
await campaign.connect(donor).claimBackerReward();
```

Или backend может батчево минтить:

```ts
await campaign.mintBackerRewards(offset, limit);
```

### Metadata

```ts
const uri = await campaign.tokenURI(tokenId);
```

`tokenURI()` возвращает:

```text
data:application/json;base64,...
```

Внутри JSON:

- `name`: например `Medical LLM crowdfunding - Lead Backer`;
- `description`;
- `image`: `data:image/svg+xml;base64,...`;
- `attributes`: campaign, backer level, campaign ID.

`image` это не PNG/JPEG и не пиксельная картинка. Это base64-encoded SVG-карточка с текстом, чтобы кошелек мог показать preview.

## События, Которые Стоит Индексировать

### GPULeaseWallet

- `Deposit(user, amount)`
- `Withdraw(user, amount)`
- `LeaseManagerUpdated(previousManager, newManager)`

### GPULease

- `LeaseStarted(leaseId, user, provider, startTimestamp, activationTimestamp, duration, amount)`
- `LeaseCompleted(leaseId)`
- `LeasePaused(leaseId)`
- `LeaseResumed(leaseId)`
- `PlatformFeeUpdated(previousFeePercentage, newFeePercentage)`
- `UserFeeUpdated(user, feePercentage)`
- `UserFeeCleared(user)`
- `ReferralManagerUpdated(previousManager, newManager)`

### GPULeaseReferral

- `ReferrerUpdated(user, referrer)`
- `ReferrerCleared(user, previousReferrer)`
- `ReferralShareUpdated(previousShareBps, newShareBps)`

### Campaign

- `Donated(donor, amount, totalDonated, grade)`
- `BackerGradeUpdated(donor, previousGrade, newGrade, totalDonated, targetShareBps)`
- `CampaignSucceeded(totalRaised)`
- `CampaignFailed(totalRaised)`
- `Refunded(donor, amount)`
- `FundsTransferred(wallet, amount)`
- `BackerRewardMinted(donor, tokenId, campaignName, grade)`

### CampaignFactory

- `CampaignCreated(campaignId, campaign, creator, targetAmount, startTimestamp, duration, templateId, campaignName)`
- `CampaignParticipantRegistered(participant, campaign)`

## Типичные Backend Flows

### 1. User Deposit

```text
user approve GPULeaseWallet
user GPULeaseWallet.deposit
backend indexes Deposit
backend reads userBalance
```

### 2. Backend Starts Lease

```text
backend checks GPULeaseWallet.userBalance(user)
backend GPULease.startLease(...)
backend indexes LeaseStarted
```

### 3. Backend Completes Lease

```text
backend GPULease.completeLease(leaseId)
backend indexes LeaseCompleted
backend reads balances for user/provider/treasury/referrer if needed
```

### 4. Set Referral

```text
backend/owner GPULeaseReferral.setReferrer(user, referrer)
backend indexes ReferrerUpdated
future leases capture this referrer
```

### 5. Campaign Success NFT

```text
donors donate
campaign reaches target
campaign deposits funds to GPULeaseWallet
donor claimBackerReward
wallet displays NFT image from tokenURI metadata
```

## Важные Ограничения

- `GPULeaseWallet` должен быть долгоживущим контрактом. Не redeploy без миграции балансов.
- `GPULease` можно заменить, но активные leases живут в конкретном экземпляре `GPULease`.
- Перед заменой `GPULease` лучше завершить активные аренды или реализовать lease migration.
- `CampaignFactory` использует clone pattern, чтобы не зашивать полный creation bytecode `LLMFundraising` в factory.

## Минимальный ABI, Который Нужен Backend

### GPULeaseWallet

```solidity
function credit() view returns (address);
function leaseManager() view returns (address);
function balances(address user) view returns (uint256);
function userBalance(address user) view returns (uint256);
function deposit(uint256 amount);
function depositFor(address beneficiary, uint256 amount);
function withdraw(uint256 amount);
```

### GPULease

```solidity
function platformFeePercentage() view returns (uint256);
function feePercentageForUser(address user) view returns (uint256);
function userFeePercentage(address user) view returns (uint256);
function hasCustomUserFee(address user) view returns (bool);
function treasury() view returns (address);
function referralManager() view returns (address);
function frozenFunds(uint256 leaseId) view returns (uint256);
function leases(uint256 leaseId) view returns (...);
function leaseReferralInfo(uint256 leaseId) view returns (address referrer, uint256 referralShareBps);
function getUserFrozenFunds(address user) view returns (...);
function setPlatformFee(uint256 feePercentage);
function setUserFee(address user, uint256 feePercentage);
function clearUserFee(address user);
function setReferralManager(address newReferralManager);
function setTreasury(address newTreasury);
function startLease(uint256 startTimestamp, uint256 duration, uint256 storagePricePerSecond, uint256 computePricePerSecond, address provider, address user);
function leaseActivationTime(uint256 leaseId) view returns (uint256);
function calculateActualCost(uint256 leaseId) view returns (uint256 storageCost, uint256 computeCost);
function pauseLease(uint256 leaseId);
function resumeLease(uint256 leaseId);
function completeLease(uint256 leaseId);
```

### GPULeaseReferral

```solidity
function referrerOf(address user) view returns (address);
function referralShareBps() view returns (uint256);
function setReferrer(address user, address referrer);
function clearReferrer(address user);
function setReferralShareBps(uint256 shareBps);
```

### CampaignFactory

```solidity
function createCampaign(uint256 targetAmount, uint256 duration, uint256 startTimestamp, uint256 templateId, string campaignName);
function campaignById(uint256 campaignId) view returns (address);
function campaignsCount() view returns (uint256);
function campaignsByCreator(address creator) view returns (address[] memory);
function campaignsByParticipant(address participant) view returns (address[] memory);
```

### Campaign

```solidity
function donate(uint256 amount);
function refund();
function claimBackerReward() returns (uint256 tokenId);
function mintBackerRewards(uint256 offset, uint256 limit) returns (uint256 minted);
function tokenURI(uint256 tokenId) view returns (string memory);
function donorInfo(address donor) view returns (...);
function donorsCount() view returns (uint256);
function donorAt(uint256 index) view returns (address);
function donorsSlice(uint256 offset, uint256 limit) view returns (address[] memory);
```

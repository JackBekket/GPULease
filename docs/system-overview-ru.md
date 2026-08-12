# Обзор системы GPULease

Актуальный Base `GPULease`: `0xCCD732200366886e04F508D12F561ee94Eb03110`
(обновлен 2026-08-12T11:45:55Z). Постоянный `GPULeaseWallet`:
`0xD4352D14Ba7928f6066dd7ec6031C7c0CCF13340`.

Этот документ описывает текущую архитектуру GPULease, денежные потоки, модель обновлений и публичные/внешние функции контрактов.

## Общая Архитектура

Система разделена на три основных модуля GPULease:

- `GPULeaseWallet`: постоянный контракт-кошелек для ERC20-средств и балансов пользователей.
- `GPULease`: заменяемый контракт бизнес-логики аренды.
- `GPULeaseReferral`: заменяемый контракт правил реферальной системы.

Контракты кампаний работают с `GPULeaseWallet` напрямую для депозитов. `GPULease` не содержит пользовательские deposit, withdraw, balance или token getter функции; вся денежная логика пользователей живет в `GPULeaseWallet`.

```text
Пользователь / Кампания
  -> GPULeaseWallet.deposit / depositFor / withdraw
  -> GPULeaseWallet хранит токены и балансы

Owner / backend
  -> GPULease.startLease / pauseLease / resumeLease / completeLease
  -> GPULeaseWallet списывает, начисляет и переносит балансы
  -> GPULeaseReferral возвращает referrer и долю реферальной комиссии
```

## Модель Обновлений

`GPULeaseWallet` рассчитан на долгую жизнь. Он хранит токены и балансы, а также предоставляет manager-only функции для учета.

Чтобы заменить `GPULease`:

1. Задеплоить новый `GPULease(existingWallet, treasury)`.
2. Вызвать `wallet.setLeaseManager(newGPULease)`.
3. При необходимости вызвать `newGPULease.setReferralManager(existingOrNewReferral)`.

Балансы пользователей остаются в `GPULeaseWallet`.

Важное ограничение: активные аренды и `frozenFunds` живут в конкретном экземпляре `GPULease`. Заменять `GPULease` безопаснее, когда нет активных аренд, либо нужно заранее добавить явный механизм миграции активных lease.

Чтобы заменить реферальную логику:

1. Задеплоить новый referral manager, который реализует:
   - `referrerOf(address) view returns (address)`
   - `referralShareBps() view returns (uint)`
2. Вызвать `gpuLease.setReferralManager(newReferralManager)`.

Уже созданные lease сохраняют реферальные данные, зафиксированные на момент старта.

## GPULeaseWallet

Файл: `contracts/GPULeaseWallet.sol`

Назначение:

- Хранит ERC20-токен `credit`.
- Хранит пользовательские балансы в `balances`.
- Позволяет обычным пользователям напрямую делать deposit и withdraw.
- Позволяет текущему `leaseManager` менять балансы при расчетах по арендам.

Состояние:

- `credit`: ERC20-токен для платежей.
- `leaseManager`: авторизованный lease-контракт.
- `balances[user]`: внутренний баланс пользователя.

События:

- `Deposit(user, amount)`
- `Withdraw(user, amount)`
- `LeaseManagerUpdated(previousManager, newManager)`

Функции:

- `constructor(address credit_)`
  - Устанавливает ERC20-токен.

- `setLeaseManager(address newLeaseManager) onlyOwner`
  - Устанавливает единственный контракт, которому разрешены manager-only функции балансов.
  - Запрещает zero address.

- `deposit(uint256 amount)`
  - Переводит токены от caller в кошелек.
  - Начисляет внутренний баланс caller.

- `depositFor(address beneficiary, uint256 amount)`
  - Переводит токены от caller в кошелек.
  - Начисляет баланс `beneficiary`.
  - Полезно для кампаний или backend-потоков, где нужно пополнить баланс другого пользователя.

- `withdraw(uint256 amount)`
  - Списывает внутренний баланс caller.
  - Переводит токены из кошелька caller.

- `userBalance(address user) view returns (uint256)`
  - Возвращает внутренний баланс пользователя.

- `debitBalance(address user, uint256 amount) onlyLeaseManager`
  - Списывает внутренний баланс без вывода токенов наружу.
  - Используется при старте lease и заморозке средств.

- `creditBalance(address user, uint256 amount) onlyLeaseManager`
  - Начисляет внутренний баланс.
  - Используется при расчетах с provider, treasury, referrer и при возврате неиспользованных средств.

- `moveBalance(address from, address to) onlyLeaseManager returns (uint256 amount)`
  - Переносит весь внутренний баланс с одного адреса на другой.
  - Используется при смене treasury.

## GPULeaseReferral

Файл: `contracts/GPULeaseReferral.sol`

Назначение:

- Хранит связь пользователь -> referrer.
- Хранит долю referrer от платформенной комиссии.
- Может заменяться независимо от кошелька и lease-логики.

Состояние:

- `referrerOf[user]`: адрес referrer для пользователя.
- `referralShareBps`: доля referrer в basis points. По умолчанию `5000`, то есть 50% от платформенной комиссии.

События:

- `ReferrerUpdated(user, referrer)`
- `ReferrerCleared(user, previousReferrer)`
- `ReferralShareUpdated(previousShareBps, newShareBps)`

Функции:

- `constructor()`
  - Устанавливает owner.

- `setReferrer(address user, address referrer) onlyOwner`
  - Назначает referrer.
  - Запрещает zero user, zero referrer и self-referral.

- `clearReferrer(address user) onlyOwner`
  - Удаляет referrer пользователя.

- `setReferralShareBps(uint shareBps) onlyOwner`
  - Устанавливает долю referrer от платформенной комиссии.
  - `10000` означает 100% комиссии.
  - По умолчанию `5000`, то есть комиссия делится 50/50 между referrer и treasury.

## GPULease

Файл: `contracts/GPULease.sol`

Назначение:

- Основной контракт жизненного цикла аренды.
- Делегирует хранение балансов в `GPULeaseWallet`.
- Читает реферальные правила из `GPULeaseReferral`.
- Хранит lease-записи, замороженные средства, персональные fee overrides и индексы активных lease.

Состояние:

- `wallet`: immutable `GPULeaseWallet`.
- `referralManager`: заменяемый referral manager.
- `treasury`: адрес внутреннего баланса treasury.
- `platformFeePercentage`: дефолтная комиссия пользователя, сейчас `10`.
- `userFeePercentage[user]`: персональная комиссия пользователя.
- `hasCustomUserFee[user]`: включена ли персональная комиссия.
- `leases[leaseId]`: данные lease.
- `leaseReferralInfo[leaseId]`: referrer и referral share, зафиксированные при старте lease.
- `frozenFunds[leaseId]`: средства, зарезервированные под lease.
- `userActiveLeases[user]`: активные lease IDs пользователя.

События:

- `PlatformFeeUpdated(previousFeePercentage, newFeePercentage)`
- `UserFeeUpdated(user, feePercentage)`
- `UserFeeCleared(user)`
- `ReferralManagerUpdated(previousManager, newManager)`
- `LeaseStarted(leaseId, user, provider, startTimestamp, activationTimestamp, duration, amount)`
- `LeaseCompleted(leaseId)`
- `LeasePaused(leaseId)`
- `LeaseResumed(leaseId)`

Admin-функции:

- `setPlatformFee(uint feePercentage) onlyOwner`
  - Устанавливает дефолтную платформенную комиссию для пользователей без персональной комиссии.
  - Запрещает значения выше 100.

- `setUserFee(address user, uint feePercentage) onlyOwner`
  - Устанавливает персональную комиссию конкретному пользователю.
  - Запрещает zero user и значения выше 100.

- `clearUserFee(address user) onlyOwner`
  - Убирает персональную комиссию пользователя, возвращая его на дефолтную.

- `setReferralManager(address newReferralManager) onlyOwner`
  - Обновляет адрес referral manager.
  - Можно поставить zero address, чтобы отключить реферальный lookup для будущих lease.

- `setTreasury(address newTreasury) onlyOwner`
  - Переносит внутренний баланс старого treasury на новый адрес.
  - Обновляет `treasury`.

Функции комиссий:

- `feePercentageForUser(address user) view returns (uint)`
  - Возвращает персональную комиссию, если она задана.
  - Иначе возвращает дефолтный `platformFeePercentage`.

Функции жизненного цикла lease:

- `startLease(startTimestamp, duration, storagePricePerSecond, computePricePerSecond, provider, user) onlyOwner returns (leaseId)`
  - Требует положительный `startTimestamp`, который не находится в будущем.
  - Требует положительный duration.
  - Требует, чтобы хотя бы одна цена была больше нуля.
  - Требует non-zero user и provider.
  - Считает максимальную стоимость:
    - `base = duration * storagePricePerSecond + duration * computePricePerSecond`
    - `fee = base * feePercentageForUser(user) / 100`
    - `frozen = base + fee`
  - Фиксирует текущий referrer и referral share для lease.
  - Списывает баланс пользователя и сохраняет frozen funds.
  - От `startTimestamp` до on-chain активации начисляет только storage.
  - Сохраняет timestamp активации в `leaseActivationTime[leaseId]`.
  - После активации начисляет storage и compute; паузы вычитаются из compute.
  - Общий `duration` отсчитывается от `startTimestamp`.
  - Создает активный lease.

- `pauseLease(uint leaseId) onlyOwner`
  - Помечает активный lease как paused и сохраняет timestamp паузы.

- `resumeLease(uint leaseId) onlyOwner`
  - Завершает паузу.
  - Добавляет длительность паузы в накопленный `pausedDuration`.

- `completeLease(uint leaseId) onlyOwner`
  - Считает фактическую стоимость storage и compute.
  - Storage cost считается по elapsed time.
  - Compute cost считается по active time без пауз.
  - Elapsed time ограничивается заявленным duration.
  - Считает платформенную комиссию от actual cost и fee, зафиксированной в lease.
  - Если есть referral:
    - `referralAmount = platformFee * referralShareBps / 10000`
    - referrer получает `referralAmount`
    - treasury получает `platformFee - referralAmount`
  - Provider получает `actualTotalCost - platformFee`.
  - User получает неиспользованный остаток frozen funds.
  - Lease помечается completed и удаляется из списка активных lease пользователя.

- `getUserFrozenFunds(address user) view returns (FrozenFundsInfo[])`
  - Возвращает active lease IDs и frozen amounts пользователя.

Internal-функции:

- `_startLease(LeaseRequest request)`
  - Internal implementation для `startLease`.

- `calculateActualCost(uint leaseId) public view`
  - Считает фактическую стоимость storage и compute.

- `_referralInfoForLease(address user) internal view`
  - Читает referral manager и возвращает referrer/share для нового lease.

## Примеры Комиссий И Рефералки

Без referrer:

```text
actualTotalCost = 1000
user fee = 10%
platformFee = 100
treasury = 100
referrer = 0
provider = 900
```

С referrer и дефолтным referral split 50/50:

```text
actualTotalCost = 1000
user fee = 10%
platformFee = 100
referrer = 50
treasury = 50
provider = 900
```

С referrer, персональной комиссией пользователя 7% и дефолтным referral split 50/50:

```text
actualTotalCost = 1000
user fee = 7%
platformFee = 70
referrer = 35
treasury = 35
provider = 930
```

Комиссия пользователя и referral info фиксируются при старте lease. Более поздние изменения не влияют на уже существующие lease.

## Система Кампаний

Файлы:

- `contracts/Campaign.sol`
- `contracts/CampaignFactory.sol`
- `contracts/CampaignMetadataRenderer.sol`

Кампании сейчас используют wallet-интерфейс:

```solidity
interface IGPULeaseWallet {
    function deposit(uint256 amount) external;
}
```

Когда кампания успешно завершается:

1. Campaign делает approve на `GPULeaseWallet`.
2. Campaign вызывает `gpuLeaseWallet.deposit(amount)`.
3. `GPULeaseWallet` переводит токены в себя.
4. Внутренний баланс начисляется на адрес campaign-контракта.

Если средства должны начисляться owner кампании или другому пользователю, интеграцию кампании стоит перевести на `GPULeaseWallet.depositFor(beneficiary, amount)`.

Metadata для reward NFT рендерится через `CampaignMetadataRenderer`. `tokenURI()` кампании возвращает base64 JSON с:

- `name`: название кампании плюс уровень бэкерства, например `Medical LLM crowdfunding - Lead Backer`.
- `description`: описание reward.
- `image`: base64-encoded SVG-карточка с текстом, не bitmap-картинка. Это позволяет кошелькам показывать текст в превью NFT без IPFS и без внешнего хостинга картинок.
- `attributes`: название кампании, уровень бэкерства и campaign ID.

### LLMFundraising

Основные public/external функции:

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

Основной поток:

- Donors отправляют USDC через `donate`.
- Donations учитываются, grade donor обновляется.
- Когда target достигнут, состояние кампании становится `SUCCESS`.
- Успешная кампания переводит средства в `GPULeaseWallet.deposit`.
- Failed campaign позволяет donors сделать refund.
- Donors успешной кампании могут mint backer reward NFT.

### LLMFundraisingFactory

Основные public/external функции:

- `createCampaign(targetAmount, duration, startTimestamp, templateId, campaignName)`
- `campaignsCount()`
- `campaignsByCreator(address creator)`
- `registerParticipant(address participant)`
- `campaignsByParticipant(address participant)`
- `participantCampaignsCount(address participant)`
- `participantCampaignAt(address participant, uint index)`
- `hasParticipatedInCampaign(address participant, address campaign)`

Основной поток:

- Factory деплоит один `LLMFundraising` implementation в своем constructor.
- Каждая новая кампания создается как OpenZeppelin `Clones` minimal proxy, указывающий на этот implementation.
- После clone factory вызывает `initialize(...)` на campaign clone.
- Implementation-контракт заблокирован от прямого initialize, а campaign clones возвращают константные ERC721 `name()` и `symbol()`.
- Отслеживает кампании по ID и creator.
- Campaign-контракты регистрируют participants обратно в factory.

Аргументы конструктора:

- `usdc`
- `gpuLeaseWallet`
- `metadataRenderer`

### CampaignMetadataRenderer

Основные public/external функции:

- `tokenURI(string campaignName, uint256 campaignId, uint8 grade)`

Основной поток:

- Конвертирует backer grade в человекочитаемый уровень.
- Собирает ERC721 JSON metadata.
- Добавляет в `image` base64-encoded SVG-карточку с названием кампании и уровнем бэкерства.
- Возвращает `data:application/json;base64,...` URI.

## Verifier

Файл: `contracts/Verifier.sol`

Функции:

- `verify(string userId)`
  - Эмитит `WalletVerified(msg.sender, userId)`.

## MockERC20

Файл: `contracts/MockERC20.sol`

Только test helper.

Функции:

- `mint(address to, uint256 amount)`
  - Минтит токены для тестов.

## Deployment Flow

Текущий `scripts/deploy.ts` деплоит:

1. `GPULeaseWallet(USDC)`
2. `GPULeaseReferral()`
3. `GPULease(wallet, treasury)`
4. `wallet.setLeaseManager(gpuLease)`
5. `gpuLease.setReferralManager(referral)`

Текущий `scripts/deploy-campaign-factory.ts` деплоит:

1. `CampaignMetadataRenderer()`
2. `LLMFundraisingFactory(USDC, GPULeaseWallet, metadataRenderer)`

## Статус Тестов

Проверенная команда:

```bash
/home/m0rs/.nvm/versions/node/v22.22.1/bin/node node_modules/hardhat/dist/src/cli.js test test/GPULease.ts 'test/GPULease (EdgeCases).ts' test/GasCosts.ts
```

Результат:

```text
33 passing
```

Дополнительный тест NFT metadata renderer:

```bash
/home/m0rs/.nvm/versions/node/v22.22.1/bin/node node_modules/hardhat/dist/src/cli.js test test/CampaignMetadataRenderer.ts
```

Результат:

```text
1 passing
```

Что сейчас покрыто в GPULease-тестах:

- Wallet deposit и withdraw.
- `depositFor` и запрет zero beneficiary.
- Start lease и frozen funds.
- Pause и resume.
- Complete lease и settlement.
- Duration edge cases с паузой и без паузы.
- Запрет withdraw больше баланса.
- Изменение дефолтной комиссии.
- Установка/очистка персональной комиссии и owner-only guard.
- Settlement с персональной комиссией.
- Фиксация fee при старте lease.
- Referral split платформенной комиссии.
- Referral split с персональной комиссией пользователя 7%.
- Фиксация referral rules при старте lease.
- Замена referral manager.
- Custom referral share.
- Admin guard checks для referral.
- Миграция treasury balance.
- Сохранение wallet balances после замены `GPULease`.
- Gas snapshots для ключевых операций.
- Campaign NFT metadata renderer возвращает JSON metadata и base64-encoded SVG-карточку с текстом.

Полная команда тестов:

```bash
/home/m0rs/.nvm/versions/node/v22.22.1/bin/node node_modules/hardhat/dist/src/cli.js test
```

Текущий результат:

```text
39 passing
```

# Ежедневные выплаты по арендам

Актуальный Base `GPULease`: `0x1350d6D31dc4c8B1314aF51d99e61cF0E3da938f`
(обновлен 2026-08-13T13:50:35Z). Wallet:
`0xf6d56d64938b65c6Ad58cFD447Cd1d74b39eEeF2`.

`GPULease` начисляет провайдеру, treasury и referrer накопившийся доход не чаще
одного раза в сутки. Начисление выполняется во внутренние балансы
`GPULeaseWallet`; провайдер по-прежнему самостоятельно вызывает `withdraw`.

## Подготовка контракта

После деплоя новый `GPULease` назначает deployer-адрес оператором по умолчанию.
Для backend лучше создать отдельный горячий кошелек и назначить его с холодного
owner-кошелька:

```ts
await gpuLease.setSettlementOperator(operatorAddress);
```

Owner также всегда может вызывать функции settlement. Operator может только
выполнять выплаты и не получает доступ к старту, паузе или завершению аренд.

Основные функции:

```solidity
function isSettlementDue(uint256 leaseId) view returns (bool);
function settleLease(uint256 leaseId);
function settleLeases(uint256[] leaseIds);
function leaseSettlements(uint256 leaseId) view returns (
    uint256 providerPaid,
    uint256 feePaid,
    uint256 referralPaid,
    uint256 lastSettledAt
);
```

Batch принимает от 1 до 50 ID. На backend рекомендуется использовать batch по
20 ID, чтобы оставить запас по gas limit.

## Worker на ethers v6

Backend должен сохранять активные `leaseId` из собственных операций
`startLease` или из событий `LeaseStarted`, а после `LeaseCompleted` удалять их
из активного списка.

```ts
import { Contract, JsonRpcProvider, Wallet } from "ethers";
import gpuleaseAbi from "./GPULease.abi" with { type: "json" };

const rpc = new JsonRpcProvider(process.env.BASE_RPC_URL!);
const signer = new Wallet(
  process.env.SETTLEMENT_OPERATOR_PRIVATE_KEY!,
  rpc,
);
const gpuLease = new Contract(
  process.env.GPULEASE_ADDRESS!,
  gpuleaseAbi,
  signer,
);

export async function runDailySettlements(activeLeaseIds: bigint[]) {
  const dueChecks = await Promise.all(
    activeLeaseIds.map(async (leaseId) => ({
      leaseId,
      due: await gpuLease.isSettlementDue(leaseId),
    })),
  );

  const dueIds = dueChecks
    .filter(({ due }) => due)
    .map(({ leaseId }) => leaseId);

  for (let offset = 0; offset < dueIds.length; offset += 20) {
    const batch = dueIds.slice(offset, offset + 20);
    const tx = await gpuLease.settleLeases(batch);
    await tx.wait(1);
  }
}
```

Рекомендуется запускать worker каждые 10–30 минут: контракт сам проверяет
суточный интервал. Запуск cron ровно раз в сутки также допустим, но выплата
сдвинется при задержке/сбое cron и восстановится только на следующем запуске.
Если worker пропустил день, следующая выплата начислит весь накопившийся доход.
Одновременно должен работать только один settlement worker:
иначе второй worker может отправить уже устаревший batch и получить revert.

Перед отправкой batch полезно повторно выполнить `isSettlementDue` либо
`estimateGas`, поскольку аренду могли завершить между чтением и транзакцией.

## Завершение аренды

`completeLease` выполняет финальный settlement без ожидания суток, возвращает
пользователю остаток frozen funds и закрывает аренду. Уже выплаченные суммы
повторно не начисляются.

## Текущая Схема Перехода

Старый `GPULease` нельзя обновить на месте, а `GPULeaseWallet` поддерживает один
`leaseManager`. Поэтому новый бонусный стек развёрнут с отдельным wallet.

До завершения legacy-аренды `#9` backend должен маршрутизировать:

1. Все новые депозиты, бонусы и аренды — в актуальные wallet/GPULease выше.
2. Завершение аренды `#9` — через legacy `GPULease`
   `0xCCD732200366886e04F508D12F561ee94Eb03110`.
3. Старый wallet `0xD4352D14Ba7928f6066dd7ec6031C7c0CCF13340` не переключать
   на новый manager, пока legacy-обязательства не закрыты.

# Streams

Streams хранят записи и consumer groups отслеживают pending/ack. `XADD`, `XREADGROUP`, `XACK` дают повторное чтение, но delivery всё равно at-least-once: consumer обязан быть idempotent. Pending без ACK — сигнал retry/claim.

## Где это применяется в реальном backend

1. **Order processing queue** — entries сохраняются для позднего worker; без ACK остаются pending.
2. **Consumer recovery** — claim переносит stale pending; прежний consumer может ещё завершить side effect.
3. **Audit feed** — stream id задаёт position; retention trimming может удалить нужную историю.

## Глубокое погружение

Stream хранит ordered entries, group хранит last-delivered и PEL. `XACK` меняет pending metadata, не удаляет запись. Delivery at-least-once: crash after side effect before ACK даёт duplicate. Costs: stream/PEL memory, polling/blocking connections. Доказывай `XPENDING`, crash point, claim и idempotency record в уже пройденном store.

## Мини-проект

### Результат

Построй поверх order store Redis Stream processor: order event меняет read model, crash до ACK воспроизводит duplicate, повтор безопасен.

### Разрешённые знания

Все Redis темы и предыдущий PostgreSQL store/concurrency; Kafka не нужна.

### Проверка

CLI/Go integration: XADD, group read, pending, restart/claim, ACK; проверить одно итоговое состояние при duplicate.

### Критерии приёмки

- [ ] pending/ACK/reclaim доказаны;
- [ ] side effect idempotent через store;
- [ ] retention и Pub/Sub difference объяснены.

### Усложнение после первой версии

Добавить bounded retry metadata и отдельный поток failed events.

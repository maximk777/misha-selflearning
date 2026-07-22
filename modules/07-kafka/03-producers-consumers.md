# Producers и consumers

Producer подтверждает запись по требуемому acks; consumer читает и коммитит offset после безопасной обработки. Manual commit после side effect даёт at-least-once. Commit до обработки даёт at-most-once и возможную потерю.

## Где это применяется в реальном backend

1. **Publish order-created** — producer acks определяют подтверждение; success API раньше ack создаёт loss window.
2. **Update read model** — consumer делает side effect, затем commit; crash между ними даёт duplicate.
3. **Batch consumption** — повышает throughput, но увеличивает latency и объём replay.

## Глубокое погружение

Producer batching/idempotent mode влияют ordering/retry; broker ack не означает business processing. Consumer poll loop обязан продолжать heartbeat/poll в рамках limits. Offset commit отмечает next position, не результат side effect. Failures: ambiguous produce, slow processing, duplicate. Доказывай crash points и committed/current offsets.

## Мини-проект

### Результат

Собери order event processor: producer публикует created events, consumer обновляет уже пройденный store и коммитит только после успешной обработки.

### Разрешённые знания

Kafka topics/partitions/offsets, PostgreSQL store, context/concurrency; groups/rebalance детали впереди.

### Проверка

Integration test success, handler error, crash after side effect/before commit; inspect committed offsets and store.

### Критерии приёмки

- [ ] broker ack и processing completion различены;
- [ ] failed side effect не коммитит progress;
- [ ] duplicate риск виден и назван.

### Усложнение после первой версии

Добавить controlled producer retry и проверить порядок records одного key.

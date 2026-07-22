# Delivery semantics

At-most-once: commit до обработки. At-least-once: обработка до commit, значит duplicates. Exactly-once — граница конкретной transactional системы, не обещание обычного HTTP/database side effect. Проектируй consumer idempotent.

## Где это применяется в реальном backend

1. **Order projection** — at-least-once сохраняет event, но требует idempotent write.
2. **Notification** — at-most-once может потерять письмо; допустимость решает продукт.
3. **Kafka transaction** — атомарен consume-transform-produce в Kafka, но не обычный HTTP/DB side effect.

## Глубокое погружение

Semantics возникают из порядка side effect/commit и failure windows. Exactly-once ограничен ресурсами одной transactional boundary; external side effect остаётся uncertain. Costs: dedup state, transaction latency, retries. Edge cases: crash before/after commit, ambiguous response. Доказывай fault injection в каждой точке и invariant итогового store.

## Мини-проект

### Результат

Для order processor воспроизведи at-most-once и at-least-once crash windows, затем выбери режим и защити business invariant idempotency.

### Разрешённые знания

Kafka producer/consumer/groups и предыдущие stores; retries/DLQ ещё не обязательны.

### Проверка

Fault-injection integration tests до side effect, после side effect, до/после commit; сверить events, offsets и store.

### Критерии приёмки

- [ ] loss/duplicate windows показаны evidence;
- [ ] выбранная semantic соответствует сценарию;
- [ ] exactly-once boundary сформулирована честно.

### Усложнение после первой версии

Добавить consume-transform-produce Kafka transaction и отдельно проверить внешний DB side effect.

# Stampede и locks

На одновременном miss много запросов бьют в source. Ограничивай singleflight/lock, jitter TTL и stale-while-revalidate. Наивный distributed lock опасен при pause и expiry: lock не превращает операцию в транзакцию без fencing/idempotency.

## Где это применяется в реальном backend

1. **Hot order после expiry** — concurrent misses бьют PostgreSQL; singleflight объединяет работу только внутри process.
2. **TTL wave** — одинаковые TTL создают нагрузочный пик; jitter размывает его.
3. **Distributed rebuild lock** — lease может истечь во время work; lock без fencing/idempotency не гарантирует correctness.

## Глубокое погружение

Stampede — queueing/amplification problem. Local singleflight даёт leader/waiters; Redis lease имеет token/expiry и требует compare-and-delete, но pause оставляет старого owner. Costs: waiter latency, lock contention, stale serving. Доказывай concurrent counter, bounded wait, failure leader и race test.

## Мини-проект

### Результат

Добавь к cached order read защиту от одновременного rebuild одного key и измерь число обращений к store под burst.

### Разрешённые знания

Cache-aside, TTL, concurrency/worker pool/context; Kafka не нужна.

### Проверка

`go test -race`; barrier запускает N misses, store counter ожидаемо bounded; tests leader failure/cancel.

### Критерии приёмки

- [ ] burst не создаёт N одинаковых DB reads;
- [ ] wait bounded и error не кэшируется бесконечно;
- [ ] local/distributed boundary и lease risk объяснены.

### Усложнение после первой версии

Добавить stale-while-revalidate и сравнить latency/staleness.

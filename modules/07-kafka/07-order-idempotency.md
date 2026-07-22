# Ordering и idempotency

Key сохраняет порядок одной сущности в одной partition. Idempotency key хранится с результатом side effect; duplicate возвращает уже принятый результат. Нельзя гарантировать порядок между partitions — меняй модель данных или компенсируй.

## Где это применяется в реальном backend

1. **Order state events** — key обеспечивает partition order; duplicate не должен повторить transition.
2. **Payment-like side effect** — event/idempotency key возвращает прежний result; payload mismatch с тем же key — конфликт.
3. **CDC replay** — snapshot/restart повторяет change; offset alone не защищает DB после crash.

## Глубокое погружение

Idempotency — business invariant, а не только dedup cache. Durable key/result записываются атомарно с side effect в уже пройденной БД. Ordering version/sequence выявляет stale/gap; partitions не дают cross-entity order. Costs: dedup storage/retention, contention on hot entity. Доказывай parallel duplicates, restart и payload conflict.

## Мини-проект

### Результат

Заверши order event processor durable idempotency: duplicates/restarts дают один side effect, stale/out-of-order version имеет явную policy.

### Разрешённые знания

Все Kafka/Redis/PostgreSQL/HTTP/concurrency темы; outbox/architecture можно упомянуть как следующий этап, но не требовать.

### Проверка

Integration tests duplicate concurrently, crash/restart, same key different payload, stale/gap version; `go test -race`.

### Критерии приёмки

- [ ] idempotency record и side effect атомарны;
- [ ] duplicate возвращает стабильный outcome;
- [ ] key scope/retention и ordering boundary защищены.

### Усложнение после первой версии

Добавить DLQ replay и доказать, что повтор не меняет итог второй раз.

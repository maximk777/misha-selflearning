# Groups и rebalance

В group partition назначается одному active consumer. Join/leave вызывает rebalance и временно останавливает consumption; обработчик должен переживать повтор записи после rebalance. Число consumer больше partitions не ускоряет topic.

## Где это применяется в реальном backend

1. **Масштабирование order processor** — partitions распределяются между instances; consumers сверх count не ускоряют работу.
2. **Deploy instance** — leave/join вызывает rebalance и паузу; in-flight record может быть повторён.
3. **Slow consumer** — превышение poll interval отзывает partitions; heartbeat alone не спасает stalled processing.

## Глубокое погружение

Coordinator управляет membership/generation и assignment. Commit старой generation может fail; revoke требует прекратить/завершить owned work. Eager/cooperative strategies отличаются объёмом перемещения. Costs: stop time, cache warmup, duplicates. Доказывай assignment logs, controlled join/leave и record accounting.

## Мини-проект

### Результат

Запусти несколько instances order processor, зафиксируй assignments, добавь/убери consumer во время in-flight processing и сохрани корректный итог.

### Разрешённые знания

Предыдущие Kafka темы, context/graceful shutdown, PostgreSQL idempotent store; delivery semantics формально дальше.

### Проверка

Integration scenario join/leave; logs partition ownership/generation, final store и duplicate count; `go test -race`.

### Критерии приёмки

- [ ] partition одновременно не обрабатывается двумя active owners штатно;
- [ ] rebalance не теряет event, duplicate допустим/видим;
- [ ] revoke/shutdown ownership объяснён.

### Усложнение после первой версии

Сравнить eager и cooperative assignment по числу перемещённых partitions.

# Pool, migrations, vacuum

Pool ограничивает соединения; слишком большой pool перегружает БД. Долгая транзакция держит locks и старые MVCC-версии. Миграции версионируют схему и должны иметь понятный rollback/forward strategy. Vacuum освобождает место для повторного использования, analyze обновляет статистику.

## Где это применяется в реальном backend

1. **Order API pool** — ограничивает DB concurrency; oversized pool перегружает БД.
2. **Zero-downtime schema change** — expand/contract поддерживает старый/new code; destructive one-shot ломает rollout.
3. **Order updates и vacuum** — dead tuples требуют cleanup; long tx удерживает horizon.

## Глубокое погружение

Pool — очередь finite connections; leaked rows/tx исчерпывают её. Migration DDL имеет lock/compatibility cost. Vacuum помечает space reusable и поддерживает visibility map, ANALYZE обновляет stats. Edge cases: idle-in-tx, pool timeout, autovacuum lag. Доказывай pool metrics, migration compatibility test и `pg_stat_*`.

## Мини-проект

### Результат

Добавь order store bounded pool и backward-compatible migration, затем воспроизведи wait при занятом pool и отсутствие долгой transaction.

### Разрешённые знания

Все предыдущие PostgreSQL/HTTP/context темы; replication не нужна.

### Проверка

Integration test старой/new schema contract, pool timeout; inspect `pg_stat_activity` и dead tuples после controlled updates.

### Критерии приёмки

- [ ] pool size/budget обоснованы, resources закрыты;
- [ ] migration совместима в rollout window;
- [ ] long tx/vacuum impact объяснён evidence.

### Усложнение после первой версии

Сделать второй contract-шаг, удаляющий старое поле только после доказанной совместимости.

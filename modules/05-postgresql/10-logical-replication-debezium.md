# Logical replication и CDC

Publication выбирает таблицы/операции, subscription применяет логические изменения. `REPLICA IDENTITY` определяет, как найти старую строку для UPDATE/DELETE. Debezium читает log/WAL через connector, превращает изменения в события и часто публикует их в Kafka; это CDC, а не замена бизнес-outbox во всех случаях.

## Где это применяется в реальном backend

1. **Order changes subset** — publication передаёт выбранные tables; schema changes не всегда реплицируются автоматически.
2. **UPDATE/DELETE identity** — replica identity позволяет найти old row; плохой key делает события дорогими/неполными.
3. **CDC integration** — Debezium превращает WAL changes в events; duplicate/reorder требуют idempotent downstream, Kafka изучается позже.

## Глубокое погружение

Logical decoding интерпретирует WAL через slot; confirmed/restart LSN влияют retention. Initial snapshot и streaming должны состыковаться. Ownership schema contract разделено producer/consumer. Failures: slot lag fills disk, missing key, connector restart duplicates. Доказывай publication tables, slot LSN/lag и повторным применением event.

## Мини-проект

### Результат

Настрой lab-поток изменений order store в логическую subscription и докажи INSERT/UPDATE/DELETE, включая поведение выбранной replica identity.

### Разрешённые знания

Все темы PostgreSQL; Debezium/Kafka можно наблюдать в готовом lab, но собственный Kafka consumer не требуется.

### Проверка

SQL проверки publication/subscription/slot, controlled changes и сравнение source/target rows; повторный запуск.

### Критерии приёмки

- [ ] выбранные operations реплицируются, лишние tables нет;
- [ ] UPDATE/DELETE identity объяснена;
- [ ] lag/duplicate/schema boundaries зафиксированы.

### Усложнение после первой версии

Добавить новую nullable колонку и проверить совместимость потока.

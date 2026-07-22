# Partitions и offsets

Ordering гарантировано только внутри partition. Один key должен стабильно выбирать одну partition; offset — позиция чтения consumer group, не глобальный ID сообщения. Больше partitions повышают параллелизм, но усложняют ordering.

## Где это применяется в реальном backend

1. **Порядок одного order** — order id как key направляет events в partition; глобального порядка между orders нет.
2. **Parallel processing** — partitions ограничивают concurrency group; лишние consumers простаивают.
3. **Replay** — offset позволяет перечитать; это position, не business event id.

## Глубокое погружение

Partitioner отображает key на partition; увеличение partition count может изменить mapping новых records. Offset монотонен только внутри partition. Committed position принадлежит group. Costs: больше files/connections/rebalance overhead. Edge cases: null key, skew/hot partition, offset out of retention. Доказывай mapping большого набора keys и чтение partition/offset metadata.

## Мини-проект

### Результат

Опубликуй последовательности изменений нескольких orders и докажи per-order ordering при параллельном чтении, не обещая global order.

### Разрешённые знания

Kafka basics и вся ранняя concurrency/data-store база; commit strategy следующей темы не требуется.

### Проверка

Логировать key/partition/offset; repeated run проверяет порядок каждого order и показывает interleaving между ними.

### Критерии приёмки

- [ ] один order сохраняет порядок;
- [ ] global ordering не заявлен;
- [ ] skew и partition-count change объяснены.

### Усложнение после первой версии

Добавить hot key workload и измерить imbalance.

# Log, brokers, topics

Topic — лог записей; broker хранит partition replicas. Retention ограничивает историю независимо от потребителя. Создай topic CLI, опиши key/value/header и не путай topic с очередью одного consumer.

## Где это применяется в реальном backend

1. **Order event history** — topic хранит факты, consumers читают независимо; это не mutable row store.
2. **Broker failure** — replicas сохраняют partition availability/durability; replication factor не заменяет backup/producer ack.
3. **Retention** — replay нового processor возможен в окне хранения; просроченные records исчезнут независимо от offset.

## Глубокое погружение

Partition — append-only ordered log segments; leader принимает writes, followers replicate. Retention удаляет segments по time/size, consumer position отдельно. Costs: disk, network replication, page cache. Failures: under-replicated partitions, leader election, oversized messages. Доказывай describe topic, offsets, broker stop/recovery и end-to-end record count.

## Мини-проект

### Результат

Создай topic order events и небольшой наблюдатель, который публикует/читает events, фиксируя key/value/header и retention boundary.

### Разрешённые знания

Пройденные Go/HTTP/PostgreSQL/Redis темы и текущие Kafka basics; delivery semantics позже.

### Проверка

CLI produce/consume from beginning, describe topic, сравнить keys/headers/count до и после controlled restart.

### Критерии приёмки

- [ ] topic/partition/replica roles объяснены;
- [ ] payload contract не смешан с mutable store;
- [ ] retention boundary зафиксирована.

### Усложнение после первой версии

Изменить retention в lab и доказать исчезновение старого segment/event.

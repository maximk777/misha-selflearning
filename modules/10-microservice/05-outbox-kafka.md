# Outbox и Kafka

> Уровень: `production` · Время: 60 минут · Практика: `postgres/outbox.go`, `kafka/`

Результат: найди `FOR UPDATE SKIP LOCKED` в `ClaimOutboxSQL`. **Один вопрос:** какую dual-write ошибку закрывает outbox? Временно убери `SKIP LOCKED` в копии запроса, предскажи конфликт workers и верни. Перескажи: event ID плюс durable idempotency сохраняют consumer side effect при повторной доставке.

## Где это применяется в реальном backend

1. **Публикация task.created/task.completed** — outbox связывает committed PostgreSQL state и eventual Kafka event. Прямой publish после commit теряет событие при crash между действиями.
2. **Параллельные publisher workers** — `SKIP LOCKED` делит backlog без ожидания одной строки. Большой batch/долгая transaction удерживает locks и pool; network publish внутри lock требует явного lease/design.
3. **Повторная доставка consumer** — event ID и durable processed record защищают side effect. Commit offset до side effect теряет работу, после side effect — допускает duplicate delivery.

## Глубокое погружение

Outbox lifecycle включает pending, claimed/lease, published и cleanup; Kafka ordering существует только внутри partition, поэтому key выбирается по aggregate. Crash matrix: до claim, после claim, после publish, до mark, после side effect, до offset commit. Инвариант exactly-once business effect строится локальной transaction consumer, не обещанием transport. Costs — DB polling/index bloat, broker batching, duplicate storage и backlog recovery. Edge cases: poison event, rebalance mid-processing, schema version, partition hotspot, producer timeout с неизвестным outcome. Доказывай fault injection каждой crash point, concurrent claim, duplicate consumer test и lag/oldest-age metrics.

## Мини-проект

### Результат

Бизнес-сценарий: аналитический side effect должен получить каждое committed task event и не применить duplicate дважды. Доведи `project/backend-lab` до потока PostgreSQL row → bounded publisher → Kafka topic keyed by task ID → consumer с durable idempotency; расширь cumulative `project/backend-lab/compose.yaml` service Kafka. Не добавляй новый business service; observable side effect может быть отдельной таблицей проекта.

### Разрешённые знания

Все предыдущие checkpoints, PostgreSQL/outbox, Kafka partitions/groups/delivery/retry/DLQ/idempotency, context/concurrency и metrics. Kubernetes не нужен для локальной проверки.

### Проверка

Из корня репозитория выполни `cd project/backend-lab && go test ./... -count=1 -race`, `docker compose -f project/backend-lab/compose.yaml up -d postgres kafka` и созданный на этом checkpoint `make -C project/backend-lab integration-kafka`. Target создаёт task, ждёт side effect, перезапускает publisher/consumer в crash points и повторно доставляет event ID; проверяются backlog age, duplicates и единственный side effect.

### Критерии приёмки

- [ ] dual write закрыт одной DB transaction, событие не публикуется для rolled-back task;
- [ ] два publishers не теряют/не обрабатывают claim одновременно без объяснённого duplicate path;
- [ ] duplicate delivery не дублирует durable side effect;
- [ ] partition key, retry/DLQ и cleanup/backlog operations защищены trade-offs.

### Усложнение после первой версии

Добавь poison event с bounded retries и DLQ, затем реализуй безопасный ручной replay с сохранением исходного event ID.

# Retries и DLQ

Transient error: ограниченный retry с backoff; poison message: причина + metadata в DLQ. Не ретраить бесконечно и не коммитить исходный offset до маршрутизации в DLQ. Наблюдай retry count и возраст сообщения.

## Где это применяется в реальном backend

1. **Temporary DB outage** — bounded backoff даёт восстановиться; immediate retry создаёт storm.
2. **Invalid order event** — poison message не должен блокировать partition навсегда; DLQ сохраняет payload/reason.
3. **Dependency throttling** — retry-after/budget ограничивают попытки; бесконечный retry превышает SLA/retention.

## Глубокое погружение

Retry меняет ordering и load; exponential backoff требует cap/jitter. До реализации нужна явная per-key policy: блокировать следующие события ключа до исхода, разрешить overtaking с последующей проверкой version либо остановить partition — выбор фиксируется как trade-off, без требования будущей архитектурной темы. DLQ routing должно завершиться до commit source offset. Crash после успешного DLQ publish, но до offset commit, повторит source record и может создать duplicate в DLQ; downstream/replay обязан учитывать original topic/partition/offset. Metadata хранит origin, attempts и error class. Доказывай controlled failures в обеих сторонах crash window, retry count, age и DLQ inspection.

## Мини-проект

### Результат

Расширь order processor классификацией transient/permanent errors, bounded retry и DLQ с достаточной диагностикой.

### Разрешённые знания

Все Kafka темы до retries, context/time, PostgreSQL store; architecture outbox не требуется.

### Проверка

Integration tests transient success on N, permanent/exhausted path, DLQ publish failure и crash после DLQ publish до source offset commit; inspect metadata, duplicate DLQ consequence и commit.

### Критерии приёмки

- [ ] attempts bounded, backoff/cancel работают;
- [ ] source commit только после success или durable DLQ;
- [ ] duplicate DLQ после crash распознаётся по origin metadata;
- [ ] per-key ordering/overtaking policy сформулирована, poison event и причина наблюдаемы.

### Усложнение после первой версии

Добавить replay tool из DLQ с dry-run и новым idempotency boundary.

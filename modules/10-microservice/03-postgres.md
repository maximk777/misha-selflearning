# PostgreSQL как source of truth

> Уровень: `production` · Время: 60 минут · Практика: `migrations/001_init.up.sql`

Результат: объясни tasks, unique idempotency key и outbox index. **Один вопрос:** почему cache не источник истины? Временно убери `UNIQUE` в копии migration, назови риск, верни constraint. Перед реальным adapter добавь одну transaction boundary: task и outbox должны сохраняться вместе.

## Где это применяется в реальном backend

1. **Durable Task API** — PostgreSQL хранит source of truth после restart replicas. In-memory success без committed row нельзя считать завершённой операцией.
2. **Idempotent create при конкурентном retry** — UNIQUE constraint решает race на уровне owner данных. Схема «сначала SELECT, потом INSERT» без constraint оставляет окно двойной записи.
3. **Task + outbox atomicity** — одна transaction фиксирует состояние и намерение event. Network publish внутри transaction держит connection/locks и всё равно не даёт atomicity с broker.

## Глубокое погружение

Transaction boundary принадлежит use case, repository выполняет queries на переданном transaction handle. Constraints — последняя линия correctness при конкуренции; ошибки SQL переводятся в domain outcomes без утечки driver text. Pool — ограниченный ресурс, и каждый долгий call блокирует других. Index outbox должен поддерживать claim predicate/order, но увеличивает write/vacuum cost. Edge cases: commit outcome unknown после connection loss, deadlock, serialization failure, migration rollback, nullable scan, context timeout. Доказывай integration tests на реальной PostgreSQL: concurrent same key, rollback after injected failure, query plan/index и clean migrations up/down.

## Мини-проект

### Результат

Бизнес-сценарий: tasks должны переживать restart API, а `task.created` не должен потеряться между DB и broker. Подключи PostgreSQL adapter к `project/backend-lab`: migrations создают tasks/outbox, create/get/complete сохраняют domain semantics, task и event атомарны. In-memory adapter остаётся fast test double; реальная DB запускается через cumulative `project/backend-lab/compose.yaml`, созданный в checkpoint `09-infrastructure/02-compose-networking`. Redis/Kafka пока не подключай.

### Разрешённые знания

Предыдущие contract/domain/HTTP checkpoints, PostgreSQL DDL/DML/transactions/MVCC/locks/indexes/pool/migrations и архитектурный outbox. Redis/Kafka runtime пока не нужен.

### Проверка

Из корня репозитория выполни `cd project/backend-lab && go test ./... -count=1 -race`, затем `docker compose -f project/backend-lab/compose.yaml up -d postgres` и созданный на этом checkpoint `make -C project/backend-lab integration-postgres`; target проверяет migrations, concurrent idempotency, rollback task+outbox, persistence после API restart и domain errors. Завершение — `docker compose -f project/backend-lab/compose.yaml stop postgres`.

### Критерии приёмки

- [ ] schema constraints защищают invariants при concurrency, а не только application pre-check;
- [ ] task/outbox commit и rollback атомарны;
- [ ] pool/timeouts настроены и connections/rows закрываются;
- [ ] unit и integration tests явно доказывают разные свойства, trade-offs объяснены.

### Усложнение после первой версии

Внедри один retriable deadlock/serialization failure и добавь bounded transaction retry только для безопасного целого use case.

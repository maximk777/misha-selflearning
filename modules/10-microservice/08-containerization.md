# Containerization

> Уровень: `production` · Время: 45 минут · Практика: `project/backend-lab/Dockerfile`

Результат: объясни два stage Dockerfile и `USER nonroot`. **Один вопрос:** почему builder не должен быть runtime image? Временно сравни отсутствие `USER` в копии и назови риск, верни non-root. Перед Docker запуском проверь image/tag и не удаляй volumes без явного подтверждения.

## Где это применяется в реальном backend

1. **Воспроизводимый release artifact** — один pinned image проходит tests и запускается в Compose/deployment. Пересборка из mutable base после проверки создаёт уже другой artifact.
2. **Локальный production-like stack** — Compose поднимает API, DB, Redis, Kafka и proxy с health/dependency checks. Это интеграционное доказательство, но не доказывает cluster scheduling или реальную production capacity.
3. **Операционная приёмка перед delivery** — health, logs, metrics, graceful stop, migrations и restart проверяются командами runbook. «Контейнеры Up» не подтверждает бизнес-flow и отсутствие backlog.

## Глубокое погружение

Финальная упаковка связывает immutable image digest, external config/secrets, persistent volumes и startup/shutdown order. Runtime не должен содержать compiler/source/credentials; PID 1 и non-root влияют на signals/permissions. Compose health dependency — лишь локальная orchestration, приложение всё равно обязано переживать restart/outage. Costs — image size/pull, build cache, multi-service resources, migrations и retained volumes. Edge cases: stale schema volume, wrong architecture, certificate/CA absence, Kafka advertised address, secret in layer, shutdown быстрее broker flush. Доказательство — clean build, vulnerability/config inspection где доступно, full Compose scenario, restart/outage experiments и ops checklist.

## Мини-проект

### Результат

Бизнес-сценарий: release Task API должен воспроизводимо подняться с нуля и пройти business-flow и outage runbook до передачи в delivery. Заверши существующий `project/backend-lab`, не создавая новый business service: используй cumulative `project/backend-lab/compose.yaml`, который создаётся в `09-infrastructure/02-compose-networking` и расширяется последующими checkpoints. Если файла ещё нет, сначала заверши тот checkpoint; до команд проверки также заверши воспроизводимые Makefile targets `integration` из testing checkpoint и `ops-check` для последовательных dependency/recovery/SIGTERM checks. Затем production-like image и stack запускают HTTP→PostgreSQL→Redis→outbox/Kafka→consumer через Nginx, дают logs/metrics/health и корректно завершаются. Итог — tests, Compose evidence, ops checks и защита trade-offs.

### Разрешённые знания

Только изученный syllabus модулей 01–09 и checkpoints финального модуля: Go/HTTP/PostgreSQL/Redis/Kafka/outbox/workers/testing, Docker/Compose/Nginx/LB/TLS/timeouts и Kubernetes manifests как delivery context. Не добавляй service mesh, cloud-specific platform или новый framework.

### Проверка

1. Из корня репозитория: `cd project/backend-lab && go test ./... -count=1 -race`.
2. `docker compose -f project/backend-lab/compose.yaml config`, затем `docker compose -f project/backend-lab/compose.yaml build --pull` и `docker compose -f project/backend-lab/compose.yaml up -d`; дождись healthy по `docker compose -f project/backend-lab/compose.yaml ps`, без ручных sleep.
3. `make -C project/backend-lab integration` воспроизводит smoke create→replay→read→complete→event/side effect через public proxy.
4. `make -C project/backend-lab ops-check` последовательно останавливает/возвращает Redis и Kafka, масштабирует API `2→1→2`, фиксирует degraded behavior, recovery и backlog, затем отправляет SIGTERM под нагрузкой и проверяет bounded shutdown.
5. Сверь `docker compose -f project/backend-lab/compose.yaml logs`, RED metrics, health и отсутствие secret/PII. Заверши `docker compose -f project/backend-lab/compose.yaml down` без `-v`.

### Критерии приёмки

- [ ] все unit/integration tests проходят свежим запуском, Compose business-flow воспроизводим с нуля;
- [ ] image pinned/multi-stage/non-root, secrets не находятся в repository/image history, volumes не удаляются неявно;
- [ ] dependency outages, retry/idempotency, backlog recovery и graceful shutdown подтверждены ops checks;
- [ ] OpenAPI, migrations, cache, events и telemetry согласованы, нет скрытого второго source of truth;
- [ ] Миша защищает trade-offs: boundaries, consistency, delivery semantics, timeout budgets, capacity и что всё ещё не доказано локально.

### Усложнение после первой версии

Проведи canary v2 на одной replica с backward-compatible schema, собери сравнительные error/latency metrics и выполни rollback по заранее объявленному критерию.

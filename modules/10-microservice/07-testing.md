# Тесты без Docker

> Уровень: `production` · Время: 45 минут · Практика: `internal/task/service_test.go`, `tests/integration_test.go`

Результат: из `project/backend-lab` команда `go test ./...` проходит без PostgreSQL/Redis/Kafka. **Один вопрос:** что именно доказывает in-memory test, а чего не доказывает? Сломай second completion, дождись conflict failure и верни инвариант. Compose integration tests включай отдельно после healthy dependencies и cleanup rows.

## Где это применяется в реальном backend

1. **Fast domain regression suite** — unit tests проверяют invariants и error mapping без external systems. Fake repository не доказывает SQL constraints, transaction isolation или driver behavior.
2. **Integration proof adapters** — реальные PostgreSQL/Redis/Kafka обнаруживают migration, serialization и delivery ошибки. Один happy path создаёт ложную уверенность без concurrency/failure scenarios.
3. **Contract и end-to-end smoke** — HTTP через proxy подтверждает wiring и public semantics. End-to-end test медленнее и хуже локализует дефект, поэтому не заменяет нижние уровни.

## Глубокое погружение

Test pyramid здесь разделяет свойства: pure/domain, service with fakes, adapter integration, contract, Compose smoke. Determinism требует управляемых clock/IDs/timeouts и isolation данных; `sleep` — источник flaky tests. Parallel tests владеют уникальными records/topics/DB schema или последовательным fixture. Costs — runtime, environment startup, maintenance и diagnostic quality. Edge cases: race only under load, test passes against stale container, cleanup hides failure, retry masks bug, generated contract drift. Доказывай свежими командами, `-race`, repeated count, targeted fault injection и явной таблицей «какой тест что не доказывает».

## Мини-проект

### Результат

Бизнес-сценарий: команда должна быстро ловить domain regression и отдельно доказывать реальное wiring PostgreSQL/Redis/Kafka перед release. Собери test suite `project/backend-lab`: domain/HTTP tests работают без Docker, adapter tests — с реальными dependencies, Compose smoke идёт через public proxy и event flow. Добавь failure matrix и воспроизводимый Makefile target `integration`, который использует cumulative `project/backend-lab/compose.yaml`; тесты не должны удалять чужие volumes или данные.

### Разрешённые знания

Весь пройденный курс и checkpoints 01–06, `testing`, `httptest`, race detector, PostgreSQL/Redis/Kafka/Compose. Новые test frameworks не обязательны.

### Проверка

Из корня репозитория выполни `cd project/backend-lab && go test ./... -count=1 -race`, повтор критичных tests точной командой `cd project/backend-lab && go test ./internal/task/... -count=20`, затем `docker compose -f project/backend-lab/compose.yaml up -d --build` и `make -C project/backend-lab integration`. Target выполняет smoke create→read→complete→event через public proxy и очищает только собственные fixtures; зафиксируй image digests и health dependencies.

### Критерии приёмки

- [ ] каждый бизнес-инвариант имеет быстрый test, каждый adapter — реальный integration proof;
- [ ] suite покрывает invalid input, conflicts, dependency outage, duplicate delivery и shutdown;
- [ ] tests не зависят от случайного порядка/sleep и очищают только собственные данные;
- [ ] Миша объясняет границы unit/integration/contract/E2E и оставшиеся недоказанные риски.

### Усложнение после первой версии

Внедри deterministic clock и случайный seed в один flaky-prone сценарий, затем докажи воспроизводимость одного и того же failure.

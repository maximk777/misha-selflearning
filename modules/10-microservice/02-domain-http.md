# Domain и HTTP

> Уровень: `production` · Время: 60 минут · Практика: `internal/task`, `internal/platform/httpserver`

Результат: из `project/backend-lab` команда `go test ./...` проверяет idempotent create и conflict completion без Docker. **Один вопрос:** где должна жить проверка пустого title — в handler или service? Временно сделай `Create` принимать пустой title, увидь failure и почини. Перескажи путь request ID и стабильного error envelope.

## Где это применяется в реальном backend

1. **Единый инвариант независимо от transport** — service запрещает пустой title и повторное completion для HTTP, worker и тестов. Проверка только в handler обходится другим caller.
2. **Преобразование domain errors в HTTP** — adapter стабильно отображает not found/conflict/validation, не сравнивая строки ошибок. Domain не должен знать status code, иначе смена transport заражает core.
3. **Cancellation и request metadata** — handler передаёт context и trusted request ID вниз. Хранение context внутри struct или использование после запроса создаёт leaks и неверное ownership.

## Глубокое погружение

Handler владеет decode/encode и transport validation, service — use-case orchestration/invariants, repository port — persistence contract. Dependency direction идёт внутрь; interfaces определяются потребителем только там, где дают реальную substitution boundary. Costs чрезмерных layers — boilerplate и потерянный stack context; отсутствие boundary — смешанные тесты и случайный coupling. Edge cases: duplicate JSON field, unknown field policy, empty/whitespace title, cancellation после mutation, panic, concurrent complete. Доказывай table-driven service tests, `httptest` для mapping, race detector и request ID/error envelope assertions.

## Мини-проект

### Результат

Бизнес-сценарий: Task API должен одинаково защищать domain invariants при вызове через HTTP и будущего worker. Реализуй contract checkpoint в `project/backend-lab`: domain/service корректно выполняют create/get/complete, HTTP adapter сохраняет envelope/request ID и не содержит storage-specific логики. До кода представь ownership validation/errors; оставь persistence и messaging на следующих checkpoints.

### Разрешённые знания

Контракт предыдущего checkpoint, Go domain/interfaces/errors, HTTP/middleware/context/tests и все ранние темы. PostgreSQL/Redis/Kafka checkpoints пока не требуются для working version: используй existing in-memory ports.

### Проверка

Из корня репозитория выполни `cd project/backend-lab && go test ./... -count=1 -race`; `httptest` или локальные HTTP requests проверяют happy path, invalid JSON/title, not found, conflict, replay и cancellation.

### Критерии приёмки

- [ ] domain invariants выполняются вне HTTP и покрыты service tests;
- [ ] handler отвечает по OpenAPI и не раскрывает internal errors;
- [ ] request ID проходит в response/log, context не сохраняется после call;
- [ ] Миша объясняет dependency direction и цену выбранного числа layers.

### Усложнение после первой версии

Добавь optimistic version к completion и воспроизведи конкурентный conflict двумя calls без PostgreSQL-specific кода в domain.

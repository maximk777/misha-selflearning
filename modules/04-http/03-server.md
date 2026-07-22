# Server

`http.Server` отделяет конфигурацию от handler. JSON service обязан задать `Content-Type`, использовать понятные статусы и не возвращать внутренние ошибки клиенту. Ошибки кодируй единым envelope. Проверяй server через `httptest`, не через реальный порт в unit-test.

## Где это применяется в реальном backend

1. **Task API boundary** — server владеет listener/lifecycle, handlers — запросами; запуск listener в handler смешивает ownership.
2. **JSON contract** — create/get task имеют стабильную схему; internal errors не должны раскрывать реализацию.
3. **Slow clients** — timeouts защищают sockets/goroutines; слишком короткий budget рвёт допустимый request.

## Глубокое погружение

Handlers выполняются конкурентно: shared state требует синхронизации. Headers предшествуют body; ошибка encode после status уже не позволяет заменить ответ. Costs: connections, goroutines, buffers. Edge cases: empty/trailing JSON, disconnect, panic, partial response. Доказывай через test server, race detector и assertions.

## Мини-проект

### Результат

Доведи Task API до запуска через явно настроенный `http.Server`, сохранив единый JSON contract и отделив lifecycle от handlers.

### Разрешённые знания

`net/http`, HTTP client и ранний Go/concurrency; middleware и PostgreSQL не нужны.

### Проверка

`go test ./...`, `go test -race ./...` и `curl` для valid, empty и malformed JSON.

### Критерии приёмки

- [ ] timeouts обоснованы, contract согласован;
- [ ] internal errors скрыты, shared state race-free;
- [ ] lifecycle/handler ownership объяснён.

### Усложнение после первой версии

Добавить read-header timeout и показать эффект тестом.

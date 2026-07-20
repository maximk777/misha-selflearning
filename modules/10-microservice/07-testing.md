# Тесты без Docker

> Уровень: `production` · Время: 45 минут · Практика: `internal/task/service_test.go`, `tests/integration_test.go`

Результат: `go test ./...` проходит без PostgreSQL/Redis/Kafka. **Один вопрос:** что именно доказывает in-memory test, а чего не доказывает? Сломай second completion, дождись conflict failure и верни инвариант. Compose integration tests включай отдельно после healthy dependencies и cleanup rows.

# Domain и HTTP

> Уровень: `production` · Время: 60 минут · Практика: `internal/task`, `internal/platform/httpserver`

Результат: `go test ./...` проверяет idempotent create и conflict completion без Docker. **Один вопрос:** где должна жить проверка пустого title — в handler или service? Временно сделай `Create` принимать пустой title, увидь failure и почини. Перескажи путь request ID и стабильного error envelope.

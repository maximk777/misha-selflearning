# Workers, context и shutdown

> Уровень: `production` · Время: 45 минут · Практика: `cmd/api/main.go`

Результат: server завершает `Shutdown` по SIGINT/SIGTERM с bounded timeout. **Один вопрос:** почему shutdown context не равен request context? Запусти API, отправь Ctrl-C и объясни порядок stop intake → wait → timeout. Worker pool добавляется вокруг outbox claim с bounded concurrency и cancellation.

# Redis cache-aside

> Уровень: `production` · Время: 45 минут · Практика: `internal/platform/redis/cache.go`

Результат: прочитай in-memory cache port с TTL. **Один вопрос:** что происходит при miss? Сымитируй отсутствие entry, затем объясни read-through from repository и Set. Redis outage не ломает корректный read: это degraded latency, а не повод потерять source of truth.

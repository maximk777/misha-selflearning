# PostgreSQL как source of truth

> Уровень: `production` · Время: 60 минут · Практика: `migrations/001_init.up.sql`

Результат: объясни tasks, unique idempotency key и outbox index. **Один вопрос:** почему cache не источник истины? Временно убери `UNIQUE` в копии migration, назови риск, верни constraint. Перед реальным adapter добавь одну transaction boundary: task и outbox должны сохраняться вместе.

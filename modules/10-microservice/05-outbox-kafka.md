# Outbox и Kafka

> Уровень: `production` · Время: 60 минут · Практика: `postgres/outbox.go`, `kafka/`

Результат: найди `FOR UPDATE SKIP LOCKED` в `ClaimOutboxSQL`. **Один вопрос:** какую dual-write ошибку закрывает outbox? Временно убери `SKIP LOCKED` в копии запроса, предскажи конфликт workers и верни. Перескажи: event ID плюс durable idempotency сохраняют consumer side effect при повторной доставке.

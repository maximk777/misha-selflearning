# Middleware и JSON

Middleware оборачивает `next`: создаёт/принимает request ID, кладёт его в header/context и вызывает `next`. JSON decoder не должен принимать неизвестные поля без решения команды; validation возвращает `400` с `{\"error\": ...}`. Нельзя писать два разных ответа после ошибки.

## Где это применяется в реальном backend

1. **Correlation ID** — middleware связывает логи Task API; входной id требует лимита формата/длины.
2. **Error envelope** — клиенты получают стабильный JSON; поздняя обёртка не исправит уже записанный body.
3. **Create validation** — strict decoder ловит unknown fields; permissive режим принимает опечатку.

## Глубокое погружение

Порядок middleware меняет поведение. Context values только request-scoped. Инварианты: `next` не более раза, error завершает path, status один. Buffering/logging стоят памяти и latency. Доказывай table tests порядка, id, unknown/trailing JSON и отсутствия двойной записи.

## Мини-проект

### Результат

Добавь Task API request-ID middleware и строгий JSON input/error contract; сначала опиши порядок обёрток.

### Разрешённые знания

HTTP server/client, context и ранний Go/concurrency; frameworks не нужны.

### Проверка

`go test ./...`; table tests входного/созданного id, unknown field, trailing JSON и validation error.

### Критерии приёмки

- [ ] `next` вызван ровно раз, id ограничен/возвращён;
- [ ] malformed/invalid input различимы;
- [ ] после ошибки success body не пишется, порядок объяснён.

### Усложнение после первой версии

Добавить recovery middleware с безопасным `500` и request ID.

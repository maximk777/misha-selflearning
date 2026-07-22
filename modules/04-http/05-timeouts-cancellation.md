# Timeouts и cancellation

У server и client разные timeout-границы. Handler передаёт `r.Context()` в долгую работу и прекращает её при disconnect/deadline. Не создавай background context внутри request path: это ломает отмену. Проверяй timeout детерминированно с коротким test deadline.

## Где это применяется в реальном backend

1. **Отмена Task operation** — disconnect отменяет downstream work; background context оставляет сироту.
2. **Budget цепочки** — upstream получает часть SLA; независимые timeouts могут его превысить.
3. **Slow clients** — network timeout защищает connection, но не отменяет бизнес-операцию.

## Глубокое погружение

Cancellation закрывает `Done`, но код обязан его наблюдать. Deadline абсолютен; `cancel` освобождает timer. Context передаётся вниз и не хранится в struct. Failures: timer/goroutine leak, work after response, partial side effect. Доказывай controlled dependency и channel signals без sleeps.

## Мини-проект

### Результат

Добавь Task API долгую операцию, прекращающуюся по request cancel/deadline и дающую согласованный error до записи ответа.

### Разрешённые знания

Пройденные HTTP темы, context, channels, select, goroutines, sync; БД не нужна.

### Проверка

`go test ./...`, `go test -race ./...`; доказать остановку work после cancel и отсутствие результата после deadline.

### Критерии приёмки

- [ ] `r.Context()` доходит до work, ресурс освобождён;
- [ ] timeout/internal error различимы;
- [ ] объяснены budget и partial side effect.

### Усложнение после первой версии

Дать downstream более короткий budget и защитить разницу тестом.

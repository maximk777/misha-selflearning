# HTTP graceful shutdown

На сигнал создай context с deadline и вызови `server.Shutdown(ctx)`: listener закрывается, новые connections не принимаются, активным дают закончить до deadline. Обрабатывай `http.ErrServerClosed` как штатный результат. В тесте вызывай Shutdown напрямую — не посылай настоящий OS signal.

## Где это применяется в реальном backend

1. **Deploy Task API** — instance прекращает intake и завершает active requests; `Close` создаёт ошибки.
2. **Остановка workers** — intake/work требуют порядка; иначе появляется новая работа.
3. **Deadline exhaustion** — зависший handler не блокирует процесс; слишком короткий budget рвёт request.

## Глубокое погружение

`Shutdown` закрывает listeners и ждёт connections, но не arbitrary goroutines. Lifecycle у composition root. Порядок: stop intake, cancel/wait owned work, close dependencies; повтор безопасен. Failures: неверно трактовать `ErrServerClosed`, потерять error, зависнуть без deadline. Доказывай blocking handler test.

## Мини-проект

### Результат

Заверши Task API lifecycle: active request получает шанс закончить, новый intake прекращается, зависшая работа ограничена deadline.

### Разрешённые знания

Все HTTP/concurrency темы; signal только в main, тест вызывает shutdown напрямую.

### Проверка

`go test ./...`, `go test -race ./...`; normal shutdown, deadline exhaustion, `ErrServerClosed`.

### Критерии приёмки

- [ ] shutdown bounded и repeat-safe;
- [ ] intake прекращён, active request завершён/явно оборван;
- [ ] порядок и `Shutdown` vs `Close` объяснены.

### Усложнение после первой версии

Добавить owned worker и доказать порядок его остановки.

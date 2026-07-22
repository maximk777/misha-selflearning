# Client

Клиент имеет timeout или request context. После успешного `Do` всегда `defer resp.Body.Close()`, даже при не-2xx статусе. Сначала проверь status, затем ограниченно декодируй JSON. Timeout защищает ожидание; cancellation сообщает, что результат больше не нужен.

## Где это применяется в реальном backend

1. **Вызов Task API из worker** — client получает задачу или отправляет результат; без timeout зависший upstream удержит goroutine и connection.
2. **Агрегация нескольких HTTP API** — конкурентные запросы уменьшают latency; без лимита параллелизма сервис создаёт нагрузочный всплеск на себя и партнёра.
3. **Чтение error response** — status проверяется до декодирования успешной модели; для reuse соединения body обычно нужно дочитать до EOF с разумным limit и закрыть. Для большого/неожиданного non-2xx допустимо сразу закрыть body, осознанно потеряв reuse этого соединения.

## Глубокое погружение

`http.Client` безопасен для повторного использования и владеет transport/pool, а response body передаёт caller'у обязанность обработки и закрытия. Одного `Close` недостаточно для гарантированного reuse: нормальный путь дочитывает ограниченное body до EOF и закрывает его; error path явно выбирает bounded drain+close либо немедленный close с потерей reuse. Общий timeout ограничивает полный обмен, request context связывает отмену с вызывающей операцией. Инварианты dispatcher: in-flight не больше лимита, каждый успешный `Do` заканчивается выбранной body policy, отменённая работа не публикует ложный успех, сбор результатов не имеет data race. Цена — sockets, goroutines, память bodies и очередь; failures — partial read, oversized/malformed body, slow headers/body и исчерпание pool. Доказывай test server'ом, который считает новые connections, выдаёт bounded и oversized non-2xx bodies, а также проверкой `go test -race` и отмены.

## Мини-проект

### Результат

Расширь Task API клиентом и cumulative HTTP dispatcher: набор запросов обрабатывается конкурентно, с worker pool, semaphore, общим `context`, безопасным сбором результатов и ошибок. До кода предложи модель ownership и остановки; готовая архитектура не задана.

### Разрешённые знания

HTTP client и все ранее пройденные темы: goroutines, channels, `select`, `context`, sync, race/leak diagnosis, worker pool и semaphore. Следующие HTTP темы не обязательны.

### Проверка

`go test ./...`, `go test -race ./...`; test server наблюдает предел concurrency, отмену, bounded/oversized non-2xx, reuse соединения после drain до EOF и осознанную потерю reuse при раннем close.

### Критерии приёмки

- [ ] каждый response body закрывается, bounded body дочитывается до EOF для reuse, а oversized/non-2xx path имеет явно выбранную drain/close policy;
- [ ] worker pool и semaphore реально ограничивают in-flight requests;
- [ ] cancellation прекращает ожидание, сбор race-free и ошибки не теряются;
- [ ] тесты защищают invariants, а trade-offs объяснены.

### Усложнение после первой версии

Добавить отдельный budget на один запрос внутри общего deadline и сравнить поведение.

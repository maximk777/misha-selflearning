# select

`select` ждёт готовый case; `default` делает операцию неблокирующей, но способен создать busy loop. `case <-ctx.Done()` — путь отмены. При нескольких готовых case выбор псевдослучаен: не полагайся на порядок. Таймаут — отдельный канал/контекст, а не бесконечное ожидание.

## Где это применяется в реальном backend

1. **Результат или timeout** — `select` ограничивает ожидание job; timeout не останавливает работу автоматически.
2. **Несколько очередей** — worker реагирует на jobs и stop; готовые cases не имеют priority guarantee.
3. **Optional output** — nil channel отключает case; `default` в tight loop сжигает CPU.

## Глубокое погружение

Select проверяет готовность коммуникаций и блокируется, если ready cases нет и отсутствует `default`; несколько ready выбираются псевдослучайно. Nil case никогда не ready, closed receive всегда ready и способен вызвать spin, если его не отключить. Timer нужно создавать/останавливать осознанно; `time.After` в длинном цикле создаёт ресурсы до срабатывания. Инвариант — каждый блокирующий путь имеет нужный exit. Докажи deterministic tests с управляемыми channels/time budget и CPU profile для busy loop.

## Мини-проект

### Результат

Добавь runner ожидание результата job с timeout и отдельным stop-channel, не оставляя заблокированного sender.

### Разрешённые знания

Предыдущие concurrency темы, `select`, channels и timers; без context/HTTP.

### Проверка

Из корня: `cd project/concurrency-runner && go test -race -count=20 ./...` для результата, timeout, stop и одновременно готовых cases.

### Критерии приёмки

- [ ] нет предположения о priority select;
- [ ] closed/nil channel обработаны без spin;
- [ ] timeout path не оставляет goroutine ждать send.

### Усложнение после первой версии

После закрытия одного input отключить его через nil и продолжить второй.

# Races, deadlocks, leaks

Race — несинхронизированный одновременный доступ, где есть запись: ищи `go test -race ./...`. Deadlock — все ждут: используй отдельный demo, который runtime завершает с `fatal error`, либо `go test -timeout`; исправляй порядок lock. Leak — работа не узнаёт об отмене: добавь `ctx.Done()` и дождись worker. Не лечи проблему `Sleep`.

## Где это применяется в реальном backend

1. **Параллельная статистика** — unsynchronized map/counter даёт race и неверный итог; тест без `-race` может пройти.
2. **Несогласованный lock order** — две jobs взаимно ждут locks; timeout диагностирует, но не исправляет invariant.
3. **Отменённый batch** — worker ждёт send/receive после ухода collector; leak постепенно съедает память и scheduler time.

## Глубокое погружение

Data race — нарушение memory model, не просто «не то число». Deadlock бывает глобальным runtime panic или локальным зависанием части системы. Leak — живая, но навсегда ненужная goroutine со ссылками на память. Invariants: единый synchronization owner, фиксированный lock order, cancel-aware blocking operations, обязательное ожидание spawned work. Race detector динамический и покрывает лишь исполненные paths; goroutine profile/stack dump и bounded tests дополняют его. Докажи намеренной поломкой, `go test -race`, встроенным test timeout или runtime deadlock demo и повторным сравнением goroutine profiles.

## Мини-проект

### Результат

Создай три изолированных regression-сценария runner: race и leak проверяются обычными tests, а намеренный deadlock — существующим отдельным demo `go -C labs/concurrency run ./deadlock/starter -deadlock`, который Go runtime завершает с `fatal error: all goroutines are asleep - deadlock!`. Каждый сценарий сначала воспроизводит дефект, затем проходит после локальной починки.

### Разрешённые знания

Все предыдущие concurrency темы; без worker-pool abstraction и HTTP.

### Проверка

Из корня: `cd project/concurrency-runner && go test -race -count=50 ./...` для race/leak. Затем `go -C labs/concurrency run ./deadlock/starter -deadlock` с ожидаемым runtime fatal error и обычный `go -C labs/concurrency run ./deadlock/starter` после исправления. Альтернатива для нового test — `go test -timeout=2s`.

### Критерии приёмки

- [ ] причины трёх классов дефектов не смешаны;
- [ ] deadlock воспроизводится изолированно и не «лечится» Sleep/большим timeout;
- [ ] leak test доказывает завершение созданных goroutines.

### Усложнение после первой версии

Добавить cancellation в момент blocked result send и подтвердить отсутствие leak.

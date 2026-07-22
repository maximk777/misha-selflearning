# Channels

Unbuffered channel синхронизирует sender и receiver; buffered разрешает до capacity отправок без receiver. Закрывает только владелец отправки, ровно один раз. Receive после close возвращает zero value и `ok=false`; send after close паникует. Выполни `go test ./labs/concurrency/...` после предсказания блокировки.

## Где это применяется в реальном backend

1. **Очередь jobs** — channel передаёт ownership задания worker; маленький buffer создаёт backpressure, безлимитную очередь он не заменяет.
2. **Сбор результатов** — workers отправляют значения collector; закрытие не тем участником вызывает panic или потерю результатов.
3. **Сигнал завершения** — close broadcast-ит receivers; отправка sentinel смешивает данные и протокол.

## Глубокое погружение

Концептуально channel содержит lock, buffer state и очереди ожидающих senders/receivers; детали runtime не API. Unbuffered send завершается при rendezvous, buffered send — пока есть место; состояния buffer: empty, partially filled, full. Nil channel блокирует send/receive навсегда и полезен позже для отключения `select` case. Closed: receive немедленно выдаёт buffered values, затем zero/`ok=false`; send и повторный close panic. Close принадлежит единственному sender-owner, receiver обычно не закрывает. Успешная передача и close создают happens-before для соответствующих наблюдений, но channel не делает произвольную общую память безопасной. Costs — contention, blocking, memory buffer. Production failures: forgotten receiver, unmatched send, double close, goroutine leak. Докажи timeout-guarded tests, `-race`, `len/cap` только как наблюдение, goroutine profile.

## Мини-проект

### Результат

Собери минимальный поток: один producer отправляет конечный набор jobs и закрывает channel, а `main` как consumer читает через `range` и вызывает инъецируемую функцию job. HTTP dispatcher — будущая адаптация после модуля HTTP.

### Разрешённые знания

GMP, goroutines, channels и весь Go Core; без WaitGroup, `select`, context и HTTP.

### Проверка

Первый implementation-capable шаг начинается в scaffold: `cd project/concurrency-runner && go test -race -count=20 ./...`; tests для unbuffered/buffered, empty input, всех полученных значений и закрытия producer после всех отправок.

### Критерии приёмки

- [ ] один явный owner закрывает каждый channel;
- [ ] single producer закрывает channel, `main` consumer дочитывает его без WaitGroup и `Sleep`;
- [ ] ученик объясняет wait queues, buffer states, nil/closed и happens-before.

### Усложнение после первой версии

Изменить capacity и объяснить backpressure без требования конкретного порядка результатов.

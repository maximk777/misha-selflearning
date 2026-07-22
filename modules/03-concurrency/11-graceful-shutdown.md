# Graceful shutdown

Shutdown: перестать принимать новую работу, отменить context, закрыть вход, дождаться worker с deadline, затем освободить ресурсы. Сигнал — только триггер; он не доказывает, что работа дренирована. Тестируй, что функция возвращается до deadline и не теряет уже принятую работу.

## Где это применяется в реальном backend

1. **Deploy/restart runner** — вход закрывается, принятые jobs завершаются в budget; мгновенный exit теряет работу.
2. **Аварийная зависшая job** — deadline ограничивает ожидание; Go не может безопасно принудительно убить некооперативную goroutine.
3. **Освобождение ресурсов** — close выполняется после остановки producers/workers; неверный порядок даёт send-after-close.

## Глубокое погружение

Shutdown — state transition: accepting → draining/cancelling → stopped, но конкретные owners и close order ученик сначала предлагает сам. Наставник проверяет модель и раскрывает по одному нарушенному инварианту после попытки. Job contract должен быть cooperative: job принимает context и обязуется возвращаться после cancellation. Если runner всё же допускает некооперативную job, deadline разрешает вернуть управление caller, но такая goroutine не считается остановленной: её судьба и невозможность принудительного kill явно отражаются в результате/документации. Jobs, подтверждённые как accepted до cutoff, получают один outcome по выбранной политике; новые после cutoff отклоняются. Один OS signal лишь входное событие. Deadline имеет единый budget. Докажи injected blocking cooperative job, bounded deadline, accepted-ID outcomes и `-race`; HTTP server shutdown появится позже.

## Мини-проект

### Результат

До кода предложи state transitions, cutoff принятия, owners и close order для Submit/Shutdown. Реализуй выбранный контракт так, чтобы cooperative jobs завершались в budget; если scope допускает stuck job, возврат по deadline отдельно сообщает, что принадлежащая ей goroutine осталась незавершённой.

### Разрешённые знания

Весь Go Core и весь модуль Concurrency; инъецируемые jobs, без HTTP.

### Проверка

Из корня: `cd project/concurrency-runner && go test -race -count=50 ./...` на normal drain, cancel, deadline, concurrent Submit и повторный Shutdown.

### Критерии приёмки

- [ ] state transitions и ownership закрытий явны;
- [ ] accepted jobs не исчезают вне выбранной политики;
- [ ] test доказывает deadline и отсутствие leaks для cooperative jobs; для допустимой stuck job результат честно фиксирует незавершённую goroutine;
- [ ] ученик защищает порядок остановки и trade-off drain/cancel.

### Усложнение после первой версии

Только после рабочей первой версии наставник просит второй shutdown-вызов; ученик сам выбирает и документирует idempotent semantics.

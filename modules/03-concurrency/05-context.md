# context

Context образует дерево: отмена родителя отменяет потомков. Передавай `context.Context` первым параметром, не храни в struct и всегда вызывай `cancel`. В цикле worker проверяй `ctx.Done()`. Context переносит deadline/cancellation, но не заменяет явное владение каналами.

## Где это применяется в реальном backend

1. **Отмена batch** — parent cancel останавливает workers; игнорирование `Done` продолжает тратить CPU после ухода caller.
2. **Deadline job** — бюджет распространяется вниз; каждый слой не должен самовольно добавлять больший timeout.
3. **Graceful stop** — process context сообщает циклам завершение; context не закрывает принадлежащие runner channels.

## Глубокое погружение

Derived contexts образуют дерево cancellation: cancel/deadline закрывает `Done`, `Err` объясняет причину. `cancel` освобождает timer и parent-child references, поэтому вызывается даже при раннем успехе. Context immutable для caller и передаётся первым; values — только request-scoped metadata, не config/бизнес-параметры. Ownership jobs/results остаётся явным. Race-safe close `Done` даёт happens-before для наблюдения отмены, но job должна кооперативно проверять context. Failures: lost cancel, detached `Background`, blocked send after cancel, deadline budget amplification. Докажи fake/injected blocking job, bounded test, `-race` и goroutine count/profile.

## Мини-проект

### Результат

Runner принимает `context.Context` и инъецируемую `func(context.Context, Job) (Result, error)`; отмена прекращает выдачу и ожидание jobs.

### Разрешённые знания

Предыдущие темы, context, select, channels, goroutines; без HTTP.

### Проверка

Из корня: `cd project/concurrency-runner && go test -race -count=20 ./...` на success, parent cancel, deadline и job, ожидающую `Done`.

### Критерии приёмки

- [ ] cancel вызывается владельцем derived context;
- [ ] ни один send/receive не остаётся заблокирован после cancel;
- [ ] ученик объясняет tree, ownership, budget и почему context не business bag.

### Усложнение после первой версии

Вернуть частичные результаты вместе с причиной отмены по явно выбранному контракту.

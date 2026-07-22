# Транзакции и ACID

Транзакция объединяет действия в одну единицу: всё фиксируется `COMMIT` или откатывается `ROLLBACK`. Atomicity — всё или ничего; Consistency — соблюдение правил; Isolation — наблюдаемое взаимодействие конкурентов; Durability — подтверждённое переживает сбой. Практика: `labs/postgres/02-transactions/`.

## Где это применяется в реальном backend

1. **Order + items** — все rows появляются вместе; отдельные commits оставляют ползаказа.
2. **Резервирование** — read-modify-write в tx сохраняет invariant; внешний HTTP call внутри tx держит resources.
3. **Смена статуса с audit** — state и audit атомарны; rollback не отменяет уже отправленное письмо.

## Глубокое погружение

Transaction живёт на одном connection до commit/rollback. Atomicity не означает изоляцию, consistency задаётся constraints/application, durability зависит от commit/WAL settings. Costs: connection, locks, MVCC versions. Edge cases: commit uncertainty, panic, context cancel. Доказывай forced error между statements и проверкой отсутствия partial data.

## Мини-проект

### Результат

Расширь order store атомарным созданием order с items и audit record, включая rollback при одном контролируемом failure.

### Разрешённые знания

DDL/DML, Go errors/context и пройденный материал; MVCC/locks можно упомянуть, но не требовать настройки.

### Проверка

Integration test happy path и failure после первого statement; проверить row counts и `go test ./...`.

### Критерии приёмки

- [ ] commit даёт полный набор, failure — ни одной части;
- [ ] rollback вызывается на всех error paths;
- [ ] объяснено, что tx не компенсирует внешний side effect.

### Усложнение после первой версии

После MVP смоделировать неопределённый результат commit и определить безопасную реакцию.

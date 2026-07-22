# Блокировки

Row locks защищают выбранные строки, table locks — операции над отношением, advisory locks — договорённость приложения. `SELECT ... FOR UPDATE` нужен для read-modify-write, но держи транзакцию короткой и бери locks в стабильном порядке.

## Где это применяется в реальном backend

1. **Order state transition** — row lock сериализует competing updates; lock без повторной проверки state недостаточен.
2. **Worker claim** — `SKIP LOCKED` делит jobs; он может голодать старую запись.
3. **Единичная maintenance job** — advisory lock координирует instances; БД не знает, какой invariant он защищает.

## Глубокое погружение

Locks принадлежат transaction/connection и образуют wait graph. Row lock не блокирует обычный MVCC read, table lock modes конфликтуют по матрице. Costs: wait time, held connections, contention. Edge cases: context timeout, lock upgrade, forgotten tx. Доказывай двумя sessions и запросом к `pg_locks`.

## Мини-проект

### Результат

Добавь order store конкурентный переход статуса, который сохраняет разрешённый state invariant под contention.

### Разрешённые знания

Предыдущие PostgreSQL темы, context/concurrency; deadlock prevention подробно в следующей теме.

### Проверка

Parallel integration test competing transitions, timeout ожидания и проверка final state/`pg_locks`.

### Критерии приёмки

- [ ] invalid двойной переход невозможен;
- [ ] transaction короткая, error/timeout обработаны;
- [ ] выбранный lock и его граница объяснены.

### Усложнение после первой версии

Реализовать безопасный claim нескольких pending orders без двойной выдачи.

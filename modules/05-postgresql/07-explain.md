# EXPLAIN

Читай план изнутри наружу: node type, estimated/actual rows, loops, time, buffers. `EXPLAIN (ANALYZE, BUFFERS)` реально выполняет запрос. Seq Scan нормален для маленькой таблицы или низкой selectivity; наличие индекса не обязывает planner его использовать.

## Где это применяется в реальном backend

1. **Slow order list** — actual rows/loops показывают настоящий объём; один total time без nodes не объясняет причину.
2. **Неверная cardinality** — stale stats ведут к плохому join/scan choice.
3. **После индекса** — planner может оставить Seq Scan для большой доли table, и это нормально.

## Глубокое погружение

Planner оценивает costs, executor даёт actuals; `ANALYZE` реально выполняет DML. Loops умножают работу узла, buffers показывают cache/I/O, estimation error направляет диагностику. Edge cases: warm cache, parameter skew, tiny tables. Доказывай одинаковым dataset и repeated measurements, не одной цифрой.

## Мини-проект

### Результат

Диагностируй два order-store query: один реально медленный, другой с нормальным Seq Scan; сформулируй evidence-based изменение только для первого.

### Разрешённые знания

Все PostgreSQL темы до EXPLAIN, benchmarks/statistics basics.

### Проверка

`EXPLAIN (ANALYZE, BUFFERS)` на representative data до/после; сохранить plans и проверить result equivalence.

### Критерии приёмки

- [ ] plan прочитан по nodes/rows/loops/buffers;
- [ ] estimate vs actual расхождение объяснено;
- [ ] optimization не меняет результат и доказана повторно.

### Усложнение после первой версии

Проверить parameter skew для редкого и частого customer/status.

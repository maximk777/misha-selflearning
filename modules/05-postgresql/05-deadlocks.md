# Deadlock

Deadlock: T1 держит A и ждёт B, T2 держит B и ждёт A. PostgreSQL обнаруживает цикл и отменяет одну транзакцию. Лечение: единый порядок locks, короткие транзакции, retry всей транзакции, а не последнего statement.

## Где это применяется в реальном backend

1. **Обновление двух orders** — разный порядок lock создаёт цикл; retry скрывает симптом, но не устраняет частоту.
2. **Order и account** — разные code paths берут resources наоборот.
3. **Batch update** — несортированные ids дают непредсказуемый order locks.

## Глубокое погружение

PostgreSQL строит wait-for graph и после `deadlock_timeout` отменяет victim. Вся tx становится aborted. Costs: latency, wasted work/WAL, retries. Edge cases: lock timeout не равен deadlock; retry storm. Доказывай controlled barrier двумя connections и проверкой SQLSTATE `40P01`.

## Мини-проект

### Результат

Воспроизведи deadlock order store, затем измени только порядок захвата так, чтобы тест стабильно проходил; retry оставь bounded safety net.

### Разрешённые знания

Locks, transactions, MVCC и Go concurrency/context.

### Проверка

Integration test с barriers, capture `40P01`, затем `-count=10` для исправленного порядка.

### Критерии приёмки

- [ ] исходный цикл доказан, не угадан;
- [ ] единый ordering предотвращает его;
- [ ] retry всей tx bounded, причина объяснена.

### Усложнение после первой версии

Добавить метрику retries и искусственно проверить alert threshold.

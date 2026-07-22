# Isolation и MVCC

MVCC хранит версии строк, поэтому читатель обычно не блокирует писателя. `READ COMMITTED` даёт новый snapshot каждому statement; `REPEATABLE READ` — транзакции; `SERIALIZABLE` может отменить конфликтующую транзакцию. Важно уметь назвать dirty/non-repeatable/phantom/read-write anomaly.

## Где это применяется в реальном backend

1. **Два оператора читают order** — snapshot определяет видимость изменений; stale decision возможен без проверки версии.
2. **Подсчёт суммы заказов** — повторный SELECT при READ COMMITTED может дать другое значение.
3. **Конкурентный лимит** — SERIALIZABLE защищает predicate invariant, но требует retry всей tx.

## Глубокое погружение

MVCC создаёт tuple versions, snapshot выбирает видимые; старые versions удерживаются активными tx. Isolation задаёт допустимые anomalies, не скорость. Costs: retries, bloat, longer snapshots. Edge cases: serialization failure, write skew, lost update. Доказывай двумя управляемыми connections/barriers, а не sleep.

## Мини-проект

### Результат

Проведи воспроизводимый эксперимент order store с двумя transactions и зафиксируй различие READ COMMITTED, REPEATABLE READ и SERIALIZABLE для одного invariant.

### Разрешённые знания

DDL/DML, transactions, SQL и Go concurrency для orchestration; locks следующей темы не обязательны.

### Проверка

Integration tests с двумя connections и каналами-барьерами; лог наблюдений и итоговые rows.

### Критерии приёмки

- [ ] anomaly воспроизводится детерминированно;
- [ ] retry охватывает всю SERIALIZABLE tx;
- [ ] snapshot/visibility/trade-off объяснены.

### Усложнение после первой версии

Добавить optimistic version check и сравнить с выбранной isolation.

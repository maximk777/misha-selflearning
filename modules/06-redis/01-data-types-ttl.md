# Types и TTL

Strings — значения/счётчики, hashes — поля объекта, sets — уникальные элементы. TTL нужен данным, которые допустимо потерять по времени; `TTL key`, `GET`, `HGETALL`, `SMEMBERS` — первая CLI-проверка. Expiry не гарантирует мгновенное физическое удаление.

## Где это применяется в реальном backend

1. **Order cache** — hash/string хранит read model; Redis не становится source of truth без отдельного решения.
2. **Idempotency key** — string с TTL ограничивает окно повторов; слишком короткий TTL пропускает late duplicate.
3. **Active order ids** — set даёт uniqueness; хранение больших payload в set тратит память.

## Глубокое погружение

Redis исполняет command атомарно в одном event loop, но цепочка commands не атомарна. TTL metadata расходует память; expiry бывает passive/active и не обещает мгновенное удаление. Тип определяет operations/encoding. Failures: eviction, hot key, oversized values. Доказывай `TYPE`, `TTL/PTTL`, memory usage и controlled expiry.

## Мини-проект

### Результат

Добавь поверх order store временный read model трёх orders с осознанно выбранными Redis types/TTL; PostgreSQL остаётся truth.

### Разрешённые знания

Пройденные HTTP/PostgreSQL темы и текущий Redis CLI/data types; cache-aside ещё не требуется.

### Проверка

CLI/Go tests create/read/expiry/type mismatch; сравнить store после исчезновения Redis key.

### Критерии приёмки

- [ ] types соответствуют операциям;
- [ ] TTL boundary и source of truth названы;
- [ ] expiry/eviction не объявлены мгновенной гарантией.

### Усложнение после первой версии

Провести expiry-boundary experiment: сравнить значение `PTTL` непосредственно до границы, логическую недоступность после неё и момент фактического освобождения памяти без обещания мгновенного удаления.

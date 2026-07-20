# Банк: PostgreSQL

Использовать после PostgreSQL I/II. Практический checkpoint обязателен: две SQL-сессии, `EXPLAIN (ANALYZE, BUFFERS)` или контролируемый deadlock. Оценивание: [RUBRIC.md](RUBRIC.md).

| ID и теги | Тип | Вопрос экзаменатора | Follow-up и практическое доказательство |
|---|---|---|---|
| PG-1 · `R1 R7` | Концептуальный | Что ACID обещает, а что не обещает приложению при внешнем side effect? | Приведи пример, где tx в БД не откатывает отправленное сообщение. |
| PG-2 · `R1` | Чтение SQL | Чем отличаются `SELECT ... FOR UPDATE` и обычный `SELECT` в конкурентном worker? | Запусти две сессии и покажи blocking/ожидание. |
| PG-3 · `R7 R21` | Практический | Реализуй получение jobs через `FOR UPDATE SKIP LOCKED` и обозначь transaction boundary. | Докажи, что два worker не взяли одну job. |
| PG-4 · `R1 R7` | Misconception probe | «Индекс всегда ускоряет SELECT». Что решает planner? | Сравни `EXPLAIN (ANALYZE, BUFFERS)` до/после индекса на селективном и неселективном запросе. |
| PG-5 · `R21` | Чтение/debugging | Две транзакции держат разные row locks и ждут друг друга. Почему PostgreSQL отменяет одну и что исправлять? | Введи одинаковый порядок locks и воспроизведи отсутствие deadlock. |
| PG-6 · `R7` | Концептуальный | Как MVCC связан с долгими транзакциями, vacuum и connection pool? | Назови метрику/симптом и безопасный способ расследования. |
| PG-7 · `surprise-old` | Интервью | Когда пригодятся GIN, partial или composite index? | По одному predicate и объяснение порядка колонок. |

<!-- ТОЛЬКО ЭКЗАМЕНАТОРУ
PG-1: ACID внутри transaction boundary; external message needs outbox/idempotency/compensation.
PG-2: FOR UPDATE locks selected rows until tx ends; ordinary read follows isolation/MVCC and may not block. Discuss scope.
PG-3: SELECT candidate + lock/update/commit with bounded tx, ensure status predicate and idempotent worker.
PG-4: planner evaluates cost/statistics/selectivity; index has write/storage cost, Seq Scan good for many rows.
PG-5: detector aborts one tx; consistent lock order/short tx/retry with care.
PG-6: old snapshots retain dead tuples, block cleanup visibility; pool exhaustion from long held connections. Never claim vacuum locks all reads.
PG-7: GIN for suitable multi-valued/full-text operators, partial for predicate subset, composite leftmost/order guided by queries.
-->

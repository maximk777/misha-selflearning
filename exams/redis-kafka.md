# Банк: Redis и Kafka

Использовать после недели 10. Тестировать не термин, а поведение при повторе, потере, rebalance и недоступности зависимости. Оценивание: [RUBRIC.md](RUBRIC.md).

| ID и теги | Тип | Вопрос экзаменатора | Follow-up и практическое доказательство |
|---|---|---|---|
| RK-1 · `R1 R7` | Концептуальный | Опиши cache-aside: read miss, write и invalidation. Где source of truth? | Нарисуй порядок для stale cache и назови допустимый stale window. |
| RK-2 · `R1` | Misconception probe | «Redis lock делает операцию точно один раз». Почему это неверно? | Назови fencing/idempotency/expiry risk и безопасную альтернативу для сценария. |
| RK-3 · `R7` | Чтение/debugging | Subscriber подключился после `PUBLISH`. Почему сообщения нет и когда выбрать Streams? | Воспроизведи потерю и назови difference retention/consumer groups. |
| RK-4 · `R1 R21` | Концептуальный | Что partition, offset и consumer group гарантируют про порядок? | Объясни, что произойдёт с ключами и order при нескольких partitions. |
| RK-5 · `R7` | Практический | Consumer обработал событие, side effect прошёл, но offset не committed. Как построить идемпотентность? | Покажи key/dedup storage и повторную доставку без двойного эффекта. |
| RK-6 · `R21` | Чтение/debugging | Poison message постоянно retry. Как устроить backoff, лимит попыток и DLQ? | Назови поля, нужные для исследования в DLQ. |
| RK-7 · `surprise-old` | Misconception probe | «Kafka exactly-once гарантирует exactly-once для любой базы и HTTP API». Разбери утверждение. | Соотнеси producer semantics, consumer side effect и transaction boundary. |

<!-- ТОЛЬКО ЭКЗАМЕНАТОРУ
RK-1: DB/source system remains truth; cache miss load/set TTL, writes update truth then invalidate/update by chosen policy. Discuss races.
RK-2: lock can expire/owner crash/network partition; exactly-once is end-to-end problem. Require idempotency/fencing where applicable.
RK-3: Pub/Sub no retention for absent subscriber; Streams retains entries and supports groups/ack/pending.
RK-4: order only within partition; group assigns partitions, rebalances move ownership, key selects partition.
RK-5: idempotency key/event ID persisted atomically with effect when possible; at-least-once expected.
RK-6: exponential bounded backoff, retry count/time, DLQ payload metadata/error/topic/partition/offset/correlation id; manual policy.
RK-7: Kafka feature scope does not cover arbitrary external effects; no magical global transaction.
-->

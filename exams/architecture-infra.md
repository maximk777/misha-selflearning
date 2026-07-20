# Банк: Архитектура и инфраструктура

Использовать после недели 11 и в итоговой защите. Практика — схема, конфигурация или запуск нескольких реплик. Оценивание: [RUBRIC.md](RUBRIC.md).

| ID и теги | Тип | Вопрос экзаменатора | Follow-up и практическое доказательство |
|---|---|---|---|
| AI-1 · `R1 R7` | Концептуальный | Разбери CAP на конкретном сетевом разделении между двумя репликами. | Какие операции ты ограничишь или отложишь и почему? |
| AI-2 · `R1` | Практический | Есть PostgreSQL update и Kafka event. Почему нужен transactional outbox и как его доставить? | Покажи transaction boundary, polling/CDC и идемпотентный publisher. |
| AI-3 · `R7` | Misconception probe | «Saga — распределённая транзакция с automatic rollback». Что неверно? | Сравни compensation с rollback и orchestration/choreography. |
| AI-4 · `R1 R21` | Чтение конфигурации | В Dockerfile один stage и в image попадает исходный код/инструменты. Что улучшит multi-stage build? | Объясни layers, minimal runtime и non-root concern. |
| AI-5 · `R7` | Чтение/debugging | Nginx retry делает запрос к перегруженным репликам. Как появляется retry storm? | Предложи bounded timeout/retry, health checks и backpressure. |
| AI-6 · `R21` | Концептуальный | Чем readiness probe отличается от liveness probe и что произойдёт при путанице? | Приведи пример rollout, где readiness защищает трафик. |
| AI-7 · `surprise-old` | Интервью | Когда оставить модульный монолит вместо микросервисов? | Назови стоимость распределённой системы и критерий выделения границы. |

<!-- ТОЛЬКО ЭКЗАМЕНАТОРУ
AI-1: CAP only under partition; clarify consistency meaning. Expect contextual availability/consistency decision, not acronym recital.
AI-2: write domain+outbox same DB tx, relay publishes/retries, consumers idempotent; outbox does not guarantee external exactly-once.
AI-3: compensation is new business action, may fail/not perfectly undo; no global ACID.
AI-4: builder separate from slim runtime, reproducible deps, fewer tools/attack surface; preserve debug/observability needs.
AI-5: retries amplify load; budgets, jitter/backoff/circuit breaking/concurrency bounds, idempotency.
AI-6: readiness gates traffic; liveness restarts stuck process. Incorrect liveness can cause restart loop; readiness failure should not necessarily restart.
AI-7: operational overhead, latency, distributed failure/data consistency; split on real domain/team/scale need.
-->

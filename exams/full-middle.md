# Четыре накопительных mock interview

Каждый набор проводится как разговор: один вопрос за сообщение, без ключа до ответа. Ведущий фиксирует ответ по [RUBRIC.md](RUBRIC.md), добавляет один follow-up и требует практическое доказательство там, где оно указано. Наборы не заменяют мини-проверки: они проверяют, умеет ли Миша связать темы.

## Mock 1 · Go Core · неделя 4

**Сценарий:** маленький сервис читает конфигурацию, нормализует данные и возвращает ошибку вызывающему коду.

1. `M1-1 R1` — Спроектируй маленький interface для чтения конфигурации. Почему не передавать конкретный client везде?
2. `M1-2 R7` — Прочитай код с typed-nil в interface и предскажи ветку `if err != nil`.
3. `M1-3 R1` — Сделай wrapped error, который caller проверяет через `errors.Is`; покажи test.
4. `M1-4 R21` — Нужны ли здесь generics? Сравни с interface и назови критерий.
5. `M1-5 surprise-old` — В профиле видно allocations. Как сформулируешь baseline и следующий измеримый шаг?
6. `M1-6 R7` — Предскажи последствия передачи struct со slice-полем по значению; что копируется, а что остаётся общим?
7. `M1-7 R21` — Как fuzzing дополняет table-driven tests и почему найденный input должен стать регрессией?

**Практический checkpoint:** за 15 минут добавить table-driven test на three boundary cases; один случай должен ломать старую реализацию.

<!-- ТОЛЬКО ЭКЗАМЕНАТОРУ: M1 ожидает consumer-side small interface, typed nil model, %w/Is test, no forced generics, measured optimization. Практика без теста не засчитывается. -->

## Mock 2 · Конкурентность · неделя 6

**Сценарий:** API получает batch задач, ограничивает параллелизм и должен корректно завершиться по SIGTERM.

1. `M2-1 R1` — Где появляется `context`, как он проходит до worker и что отменяет?
2. `M2-2 R7` — Кто владеет закрытием jobs channel при нескольких producer?
3. `M2-3 R1` — Race detector показывает shared map. Выбери mutex или single owner и обоснуй.
4. `M2-4 R21` — Очередь растёт быстрее workers. Где поставить backpressure/semaphore и какие метрики смотреть?
5. `M2-5 surprise-old` — Опиши порядок graceful shutdown и bounded deadline.
6. `M2-6 R7` — В чём разница buffered/unbuffered channel для этого API и почему buffer не является исправлением утечки?
7. `M2-7 R21` — Счётчик обновляется через CAS. Какой составной инвариант подсказывает вернуться к mutex?

**Практический checkpoint:** запустить `go test -race`, показать cancel без goroutine leak и объяснить один deliberately broken вариант.

<!-- ТОЛЬКО ЭКЗАМЕНАТОРУ: проверять ownership, propagate cancellation, WaitGroup/drain policy, race not merely hidden. M2-4 accepts bounded queue/reject/rate limit with trade-off. -->

## Mock 3 · PostgreSQL · неделя 9

**Сценарий:** два worker обрабатывают заказы, а endpoint списка стал медленным после роста таблицы.

1. `M3-1 R1` — Предложи SQL-модель получения jobs без двойной обработки.
2. `M3-2 R7` — Две tx deadlock. Что увидит приложение, что изменишь и когда retry допустим?
3. `M3-3 R1` — Прочитай `EXPLAIN (ANALYZE, BUFFERS)` с Seq Scan. Почему это может быть хорошим планом?
4. `M3-4 R21` — Выбери индекс для `WHERE tenant_id = ? AND status = ? ORDER BY created_at` и объясни, что нужно проверить.
5. `M3-5 surprise-old` — Долгая tx в pool: как связаны MVCC, vacuum и saturation connections?
6. `M3-6 R7` — Когда row lock хуже advisory lock и почему оба не заменяют бизнес-инвариант в схеме?
7. `M3-7 R21` — Данные надо передать в event stream: сравни transactional outbox и CDC без обещаний exactly-once.

**Практический checkpoint:** показать два сеанса `FOR UPDATE SKIP LOCKED`, один deadlock с bounded cleanup и один plan до/после осмысленного индекса.

<!-- ТОЛЬКО ЭКЗАМЕНАТОРУ: expect short tx, correct status transition/idempotency, lock order, planner cost/selectivity, composite index justified by query. Не требовать index если data distribution оправдывает Seq Scan. -->

## Mock 4 · Full middle · неделя 12

**Сценарий:** сервис заказов: PostgreSQL — source of truth; Redis — cache; Kafka получает `OrderCreated`; три реплики за Nginx; развёртывание описано Kubernetes manifests.

1. `M4-1 R1` — Нарисуй request path и назови timeout/cancellation boundaries от Nginx до PostgreSQL.
2. `M4-2 R7` — Спроектируй cache-aside для чтения заказа и handling stale/invalidation после update.
3. `M4-3 R1` — Как избежать dual write PostgreSQL+Kafka? Опиши outbox, relay и idempotent consumer.
4. `M4-4 R21` — Consumer получил одно событие дважды и один раз после rebalance. Как гарантировать безопасный side effect?
5. `M4-5 R7` — Какие error/health/readiness endpoints нужны и что не надо отдавать клиенту?
6. `M4-6 R1` — Защити `requests/limits`, readiness/liveness и rolling update в manifest.
7. `M4-7 surprise-old` — При деградации upstream retries усугубили нагрузку. Как поменять budgets, backoff и admission?
8. `M4-8 R21` — Почему сейчас оставить modular monolith или почему выделить сервис? Назови измеримый критерий.

**Практический checkpoint:** запустить integration flow Compose, показать один HTTP test, один SQL evidence, один повтор Kafka-like event без двойного эффекта и объяснить Nginx/Kubernetes configuration. Защита должна назвать минимум два trade-off и одну границу знания.

<!-- ТОЛЬКО ЭКЗАМЕНАТОРУ: M4 оценивает связность: context/time budgets; cache remains non-authoritative; tx outbox plus relay; durable idempotency key; stable error + correlation id; probes not interchangeable; retry amplification; architecture based on bounded evidence. Не ожидать локального Kubernetes cluster. -->

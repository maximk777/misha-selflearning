# Силлабус: темы, практика и экзамены

Каждая строка — логическая единица. Теория должна жить в указанном пути `modules/…`; если есть практика, ей соответствует связка `labs/<track>/<lab>/LAB.md`, `CHECK.md` и `starter/`. Пути лабораторий здесь намеренно записаны как код, а не как ссылки: учебные файлы создаются следующими задачами и текущая проверка не должна считать будущий материал битой ссылкой.

Уровни: `start` — первая рабочая модель, `core` — обязательный базис, `interview` — объяснение на собеседовании, `advanced` — необязательное углубление, `production` — решение с наблюдаемыми эксплуатационными последствиями. Минимальный путь включает всё, кроме `advanced`.

## 01 · Go Start

| Код и теория | Тема | Уровень | Предпосылки | Время | Практика | Пересказ | Экзамен |
|---|---|---|---|---:|---|---|---|
| `modules/01-go-start/01-syntax.md` | Переменные, условия, циклы | start | Go установлен | 1 ч | `labs/go/01-syntax` | Почему тип известен до запуска? | [go-start](../exams/go-start.md) |
| `modules/01-go-start/02-packages-modules.md` | `go env`, модули, пакеты, запуск, тест | start | синтаксис | 1 ч | `labs/go/01-syntax` | Чем модуль отличается от пакета? | [go-start](../exams/go-start.md) |
| `modules/01-go-start/03-functions.md` | Функции, несколько return | start | синтаксис | 1 ч | `labs/go/01-syntax` | Почему error обычно последний return? | [go-start](../exams/go-start.md) |
| `modules/01-go-start/04-arrays-strings.md` | Массивы, строки, bytes/runes | core | функции | 1 ч | `labs/go/03-arrays-strings-slices` | Почему `len` строки — не число символов? | [go-start](../exams/go-start.md) |
| `modules/01-go-start/05-slices.md` | Slices, `append`, capacity | core | массивы | 2 ч | `labs/go/03-arrays-strings-slices` | Когда два slice меняют один массив? | [go-start](../exams/go-start.md) |
| `modules/01-go-start/06-maps.md` | Maps, zero value, `ok` | core | синтаксис | 1 ч | `labs/go/04-maps` | Почему запись в nil map паникует? | [go-start](../exams/go-start.md) |
| `modules/01-go-start/07-structs-methods.md` | Struct, методы, receiver | core | функции | 1 ч | `labs/go/05-structs-methods` | Что выбирает value или pointer receiver? | [go-start](../exams/go-start.md) |
| `modules/01-go-start/08-values-pointers.md` | Значения, указатели, копирование | core | struct | 1 ч | `labs/go/02-values-pointers` | Что именно копируется при передаче? | [go-start](../exams/go-start.md) |
| `modules/01-go-start/09-defer-panic-recover.md` | `defer`, errors, `panic`, `recover` | interview | функции | 2 ч | `labs/go/06-errors-panic` | Где граница normal error и panic? | [go-start](../exams/go-start.md) |

## 02 · Go Core

| Код и теория | Тема | Уровень | Предпосылки | Время | Практика | Пересказ | Экзамен |
|---|---|---|---|---:|---|---|---|
| `modules/02-go-core/01-interfaces.md` | Интерфейсы и композиция | core | struct/methods | 2 ч | `labs/go/07-interfaces` | Почему interface описывает поведение? | [go-core](../exams/go-core.md) |
| `modules/02-go-core/02-errors.md` | Wrapping, `errors.Is`, `errors.As` | interview | errors | 1 ч | `labs/go/06-errors-panic` | Что сохраняет `%w`? | [go-core](../exams/go-core.md) |
| `modules/02-go-core/03-generics.md` | Прикладные generics | interview | interfaces | 1 ч | `labs/go/07-interfaces` | Когда generic хуже интерфейса? | [go-core](../exams/go-core.md) |
| `modules/02-go-core/04-table-tests.md` | Table-driven tests | core | functions | 1 ч | `labs/go/08-tests-bench-fuzz` | Зачем subtest имеет имя? | [go-core](../exams/go-core.md) |
| `modules/02-go-core/05-bench-fuzz.md` | Benchmarks и fuzzing | interview | tests | 1 ч | `labs/go/08-tests-bench-fuzz` | Что benchmark измеряет, а чего нет? | [go-core](../exams/go-core.md) |
| `modules/02-go-core/06-slice-map-internals.md` | Устройство slice/map, версии map | advanced | slices/maps | 2 ч | `labs/go/09-memory-gc` | Почему нельзя опираться на порядок map? | [go-core](../exams/go-core.md) |
| `modules/02-go-core/07-memory-escape.md` | Stack, heap, escape analysis | interview | pointers | 1 ч | `labs/go/09-memory-gc` | Почему указатель не всегда означает heap? | [go-core](../exams/go-core.md) |
| `modules/02-go-core/08-gc.md` | Сборщик мусора | advanced | memory | 1 ч | `labs/go/09-memory-gc` | Что GC не освобождает? | [go-core](../exams/go-core.md) |
| `modules/02-go-core/09-pprof.md` | `pprof` и профилирование | production | tests | 2 ч | `labs/go/10-pprof` | Что сравниваем до оптимизации? | [go-core](../exams/go-core.md) |

## 03 · Конкурентность

| Код и теория | Тема | Уровень | Предпосылки | Время | Практика | Пересказ | Экзамен |
|---|---|---|---|---:|---|---|---|
| `modules/03-concurrency/01-gmp.md` | GMP и scheduler | interview | Go Core | 1 ч | `labs/concurrency/channels` | Что делает P? | [concurrency](../exams/concurrency.md) |
| `modules/03-concurrency/02-goroutines.md` | Горутины, preemption, блокировки | core | GMP | 1 ч | `labs/concurrency/context` | Почему goroutine не равна OS thread? | [concurrency](../exams/concurrency.md) |
| `modules/03-concurrency/03-channels.md` | Buffered/unbuffered, close ownership | core | goroutines | 2 ч | `labs/concurrency/channels` | Кто и почему закрывает channel? | [concurrency](../exams/concurrency.md) |
| `modules/03-concurrency/04-select.md` | `select`, timeouts | core | channels | 1 ч | `labs/concurrency/channels` | Что будет с nil channel? | [concurrency](../exams/concurrency.md) |
| `modules/03-concurrency/05-context.md` | Cancellation tree, deadlines | production | goroutines | 2 ч | `labs/concurrency/context` | Почему context не хранит бизнес-данные? | [concurrency](../exams/concurrency.md) |
| `modules/03-concurrency/06-sync.md` | `WaitGroup`, `Mutex`, `RWMutex` | core | goroutines | 1 ч | `labs/concurrency/workerpool` | Когда mutex проще channel? | [concurrency](../exams/concurrency.md) |
| `modules/03-concurrency/07-atomic.md` | Typed atomics, CAS | interview | sync | 1 ч | `labs/concurrency/semaphore` | Где CAS не заменяет инвариант? | [concurrency](../exams/concurrency.md) |
| `modules/03-concurrency/08-races-deadlocks-leaks.md` | Race, deadlock, leaks | production | sync/context | 2 ч | `labs/concurrency/race`, `labs/concurrency/deadlock` | Как докажешь отсутствие утечки? | [concurrency](../exams/concurrency.md) |
| `modules/03-concurrency/09-worker-pool.md` | Worker pool | production | channels/context | 2 ч | `labs/concurrency/workerpool` | Где backpressure? | [concurrency](../exams/concurrency.md) |
| `modules/03-concurrency/10-semaphore-pipelines.md` | Semaphore, fan-in/out, pipeline | interview | worker pool | 2 ч | `labs/concurrency/semaphore`, `labs/concurrency/pipeline` | Что отменяется при ошибке этапа? | [concurrency](../exams/concurrency.md) |
| `modules/03-concurrency/11-graceful-shutdown.md` | Graceful shutdown | production | context | 2 ч | `labs/concurrency/shutdown` | В каком порядке прекращается сервис? | [concurrency](../exams/concurrency.md) |

## 04 · HTTP

| Код и теория | Тема | Уровень | Предпосылки | Время | Практика | Пересказ | Экзамен |
|---|---|---|---|---:|---|---|---|
| `modules/04-http/01-net-http.md` | `net/http`, request/response | core | Go Core | 1 ч | `labs/http/01-task-service` | Где живёт handler? | [http](../exams/http.md) |
| `modules/04-http/02-client.md` | Client, timeout, body close | production | net/http | 1 ч | `labs/http/01-task-service` | Чем опасен незакрытый body? | [http](../exams/http.md) |
| `modules/04-http/03-server.md` | Server, handler, JSON | core | net/http | 2 ч | `labs/http/01-task-service` | Кто декодирует JSON? | [http](../exams/http.md) |
| `modules/04-http/04-middleware-json.md` | Middleware, validation, error envelope | production | server | 1 ч | `labs/http/01-task-service` | Где создаётся request ID? | [http](../exams/http.md) |
| `modules/04-http/05-timeouts-cancellation.md` | Cancellation и shutdown | production | context | 1 ч | `labs/http/01-task-service` | Почему timeout должен иметь бюджет? | [http](../exams/http.md) |
| `modules/04-http/06-httptest.md` | `httptest` | core | tests | 1 ч | `labs/http/01-task-service` | Что тест изолирует? | [http](../exams/http.md) |
| `modules/04-http/07-graceful-shutdown.md` | HTTP graceful shutdown | production | server | 1 ч | `labs/http/01-task-service` | Почему `Close` не всегда достаточно? | [http](../exams/http.md) |

## 05 · PostgreSQL

| Код и теория | Тема | Уровень | Предпосылки | Время | Практика | Пересказ | Экзамен |
|---|---|---|---|---:|---|---|---|
| `modules/05-postgresql/01-ddl-dml.md` | DDL, DML, constraints, FK | core | SQL basics | 2 ч | `labs/postgres/01-ddl-dml` | Чем DDL отличается от DML? | [postgresql](../exams/postgresql.md) |
| `modules/05-postgresql/02-transactions-acid.md` | Transactions, ACID | core | DML | 1 ч | `labs/postgres/02-transactions` | Что rollback не компенсирует? | [postgresql](../exams/postgresql.md) |
| `modules/05-postgresql/03-isolation-mvcc.md` | MVCC и isolation levels | interview | transactions | 2 ч | `labs/postgres/03-isolation` | Почему reader не всегда блокирует writer? | [postgresql](../exams/postgresql.md) |
| `modules/05-postgresql/04-locks.md` | Row/table/advisory locks | production | transactions | 1 ч | `labs/postgres/04-locks` | Когда advisory lock уместен? | [postgresql](../exams/postgresql.md) |
| `modules/05-postgresql/05-deadlocks.md` | Deadlock diagnosis/prevention | production | locks | 1 ч | `labs/postgres/06-deadlock` | Почему retry не лечит порядок locks? | [postgresql](../exams/postgresql.md) |
| `modules/05-postgresql/06-indexes.md` | B-tree/hash/GIN, composite/partial | interview | DDL | 2 ч | `labs/postgres/07-indexes` | Почему индекс может мешать? | [postgresql](../exams/postgresql.md) |
| `modules/05-postgresql/07-explain.md` | Selectivity, cardinality, `EXPLAIN` | production | indexes | 2 ч | `labs/postgres/08-explain` | Почему planner выбрал Seq Scan? | [postgresql](../exams/postgresql.md) |
| `modules/05-postgresql/08-pool-migrations-vacuum.md` | Pool, long tx, migrations, vacuum | production | transactions | 2 ч | `labs/postgres/02-transactions` | Чем опасна долгая tx? | [postgresql](../exams/postgresql.md) |
| `modules/05-postgresql/09-wal-replication.md` | WAL, physical/logical replication | advanced | MVCC | 1 ч | `labs/postgres/09-logical-replication` | Что сохраняет WAL? | [postgresql](../exams/postgresql.md) |
| `modules/05-postgresql/10-logical-replication-debezium.md` | Publication/subscription, identity, CDC/Debezium | interview | replication | 1 ч | `labs/postgres/09-logical-replication` | Почему CDC не отменяет идемпотентность? | [postgresql](../exams/postgresql.md) |

## 06–07 · Redis и Kafka

| Код и теория | Тема | Уровень | Предпосылки | Время | Практика | Пересказ | Экзамен |
|---|---|---|---|---:|---|---|---|
| `modules/06-redis/01-data-types-ttl.md` | Redis types и TTL | core | HTTP | 1 ч | `labs/redis/01-types-ttl` | Почему TTL — часть модели данных? | [redis-kafka](../exams/redis-kafka.md) |
| `modules/06-redis/02-cache-aside.md` | Cache-aside/invalidation | production | Redis | 1 ч | `labs/redis/02-cache-aside` | Кто источник истины? | [redis-kafka](../exams/redis-kafka.md) |
| `modules/06-redis/03-stampede-locks.md` | Stampede, distributed locks limits | interview | cache | 1 ч | `labs/redis/03-stampede` | Почему lock не гарантирует correctness? | [redis-kafka](../exams/redis-kafka.md) |
| `modules/06-redis/04-pubsub.md` | Pub/Sub и потеря сообщения | core | Redis | 1 ч | `labs/redis/04-pubsub` | Что увидит поздний subscriber? | [redis-kafka](../exams/redis-kafka.md) |
| `modules/06-redis/05-streams.md` | Streams и consumer groups | interview | Pub/Sub | 1 ч | `labs/redis/05-streams` | Чем stream отличается от Pub/Sub? | [redis-kafka](../exams/redis-kafka.md) |
| `modules/07-kafka/01-log-brokers-topics.md` | Log, brokers, topics, replication | core | Redis | 1 ч | `labs/kafka/01-basics` | Почему Kafka — log? | [redis-kafka](../exams/redis-kafka.md) |
| `modules/07-kafka/02-partitions-offsets.md` | Partitions, offsets, ordering | core | Kafka basics | 1 ч | `labs/kafka/02-consumer` | Где сохраняется порядок? | [redis-kafka](../exams/redis-kafka.md) |
| `modules/07-kafka/03-producers-consumers.md` | Producer/consumer | core | partitions | 1 ч | `labs/kafka/02-consumer` | Что означает commit offset? | [redis-kafka](../exams/redis-kafka.md) |
| `modules/07-kafka/04-groups-rebalance.md` | Consumer groups/rebalance | interview | consumers | 1 ч | `labs/kafka/02-consumer` | Что происходит при rebalance? | [redis-kafka](../exams/redis-kafka.md) |
| `modules/07-kafka/05-delivery-semantics.md` | At-most/at-least/exactly-once limits | interview | offsets | 1 ч | `labs/kafka/03-delivery` | Почему exactly-once не равно exactly-once side effect? | [redis-kafka](../exams/redis-kafka.md) |
| `modules/07-kafka/06-retries-dlq.md` | Retry, backoff, DLQ, poison message | production | delivery | 1 ч | `labs/kafka/03-delivery` | Когда message попадёт в DLQ? | [redis-kafka](../exams/redis-kafka.md) |
| `modules/07-kafka/07-order-idempotency.md` | Idempotency, ordering, CDC transport | production | delivery | 1 ч | `labs/kafka/03-delivery` | Каким ключом сделать consumer idempotent? | [redis-kafka](../exams/redis-kafka.md) |

## 08–10 · Архитектура, инфраструктура, итоговый сервис

| Код и теория | Тема | Уровень | Предпосылки | Время | Практика | Пересказ | Экзамен |
|---|---|---|---|---:|---|---|---|
| `modules/08-architecture/01-cap.md` | CAP и отказы | interview | distributed basics | 1 ч | `labs/architecture/01-cap` | Какой trade-off в данном отказе? | [architecture-infra](../exams/architecture-infra.md) |
| `modules/08-architecture/02-sharding-consistent-hashing.md` | Sharding, consistent hashing | interview | CAP | 1 ч | `labs/architecture/02-sharding` | Что меняется при добавлении узла? | [architecture-infra](../exams/architecture-infra.md) |
| `modules/08-architecture/03-outbox.md` | Transactional outbox | production | PostgreSQL/Kafka | 2 ч | `project/backend-lab` | Как закрыть dual write? | [architecture-infra](../exams/architecture-infra.md) |
| `modules/08-architecture/04-saga.md` | Saga: orchestration/choreography | interview | messaging | 1 ч | `labs/architecture/03-saga` | Что компенсирует saga? | [architecture-infra](../exams/architecture-infra.md) |
| `modules/08-architecture/05-cdc-debezium.md` | CDC/Debezium | advanced | outbox/replication | 1 ч | `labs/architecture/04-cdc` | Какая граница outbox и CDC? | [architecture-infra](../exams/architecture-infra.md) |
| `modules/08-architecture/06-monolith-microservices.md` | Монолит vs микросервисы | interview | architecture | 1 ч | `labs/architecture/05-boundaries` | Когда микросервисы не нужны? | [architecture-infra](../exams/architecture-infra.md) |
| `modules/08-architecture/07-observability-security.md` | Observability и security | production | HTTP | 1 ч | `project/backend-lab` | Зачем correlation ID? | [architecture-infra](../exams/architecture-infra.md) |
| `modules/09-infrastructure/01-docker.md` | Image/container/layers/Dockerfile | core | Go | 1 ч | `labs/nginx/01-replicas` | Что попадает в image? | [architecture-infra](../exams/architecture-infra.md) |
| `modules/09-infrastructure/02-compose-networking.md` | Volumes, networks, Compose | core | Docker | 1 ч | `labs/nginx/01-replicas` | Что переживает restart? | [architecture-infra](../exams/architecture-infra.md) |
| `modules/09-infrastructure/03-nginx-reverse-proxy.md` | Nginx reverse proxy | production | HTTP | 1 ч | `labs/nginx/01-replicas` | Где заканчивается TLS? | [architecture-infra](../exams/architecture-infra.md) |
| `modules/09-infrastructure/04-load-balancing.md` | L4/L7, RR/weighted/least/conhash | interview | Nginx | 1 ч | `labs/nginx/01-replicas` | Какой алгоритм для long connections? | [architecture-infra](../exams/architecture-infra.md) |
| `modules/09-infrastructure/05-tls-timeouts-retries.md` | Health, sticky, TLS, timeout, retry storm | production | LB | 1 ч | `labs/nginx/01-replicas` | Почему retry усугубляет аварию? | [architecture-infra](../exams/architecture-infra.md) |
| `modules/09-infrastructure/06-kubernetes-core.md` | Pod, Deployment, Service, Ingress, ConfigMap, Secret | interview | Compose | 1 ч | `labs/kubernetes/01-manifests` | Что даёт Deployment? | [architecture-infra](../exams/architecture-infra.md) |
| `modules/09-infrastructure/07-kubernetes-delivery.md` | Probes, requests/limits, rolling update, autoscaling | interview | Kubernetes core | 1 ч | `labs/kubernetes/01-manifests` | Чем readiness отличается от liveness? | [architecture-infra](../exams/architecture-infra.md) |
| `modules/10-microservice/README.md` | Итог: OpenAPI/ogen, PostgreSQL, Redis, Kafka/outbox, workers, config/logs/health, tests, Compose | production | весь minimum path | 8 ч | `project/backend-lab` | Защити trade-offs сервиса | [full-middle](../exams/full-middle.md) |

## Непереговорные правила структуры

Для каждой обязательной темы проверяющий ожидает: уровень, предпосылки, ожидаемую длительность, ссылку/путь к теории, лаборатории и `CHECK.md` при практике, prompt пересказа и ссылку на экзамен. Шаблоны находятся в [templates](../templates/TOPIC.md). `DEEP_DIVE.md` создаётся только там, где есть отдельная польза от уровня `advanced`; он не входит в minimum path.

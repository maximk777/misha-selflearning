# Redis cache-aside

> Уровень: `production` · Время: 45 минут · Практика: `internal/platform/redis/cache.go`

Результат: прочитай in-memory cache port с TTL. **Один вопрос:** что происходит при miss? Сымитируй отсутствие entry, затем объясни read-through from repository и Set. Redis outage не ломает корректный read: это degraded latency, а не повод потерять source of truth.

## Где это применяется в реальном backend

1. **Ускорение GET task** — cache-aside возвращает hot object без DB round trip. Cache miss и Redis outage должны сходиться к source of truth, иначе availability cache становится correctness зависимостью.
2. **Инвалидация после complete** — write в PostgreSQL затем delete/update cache снижает stale window. Если изменить cache до commit, rollback оставит выдуманное состояние.
3. **TTL как ограничение stale data и памяти** — expiration уменьшает срок ошибки при потерянной invalidation. Слишком короткий TTL даёт miss storm, слишком длинный — старые статусы.

## Глубокое погружение

Cache-aside ownership: repository authoritative, service решает read/fill/invalidate, cache best-effort. Serialization/versioning являются частью cache key/value contract; TTL не обеспечивает мгновенную consistency. Costs — network hop, memory, serialization, connection pool и stampede на hot miss. Edge cases: cache contains malformed/old schema, concurrent read during update, negative caching, timeout, partial outage, hot key expiration. Production failure не должен менять 200/404 correctness; он меняет latency/load. Доказывай fake clock TTL tests, integration outage, stale-window experiment, hit/miss/error metrics и DB call counter.

## Мини-проект

### Результат

Бизнес-сценарий: повторные чтения Task API должны ускоряться, но outage cache не должен менять правильный ответ. Подключи Redis adapter к GET task в `project/backend-lab`, сохрани PostgreSQL source of truth и явную invalidation после mutation. Расширь cumulative `project/backend-lab/compose.yaml` service `redis`; покажи cold miss, hit, expiry, stale prevention и работу при остановленном Redis. Kafka worker пока не добавляй.

### Разрешённые знания

Предыдущие checkpoints, Redis types/TTL/cache-aside/stampede limits, context/timeouts/metrics. Kafka worker ещё не обязателен.

### Проверка

Из корня репозитория выполни `cd project/backend-lab && go test ./... -count=1 -race`, `docker compose -f project/backend-lab/compose.yaml up -d postgres redis` и созданный на этом checkpoint `make -C project/backend-lab integration-redis`. HTTP scenario сравнивает ответы до/после cache hit, complete и `docker compose -f project/backend-lab/compose.yaml stop redis`; метрики показывают hits, misses, cache errors и repository reads.

### Критерии приёмки

- [ ] cache никогда не становится source of truth и outage не меняет правильный результат;
- [ ] mutation не оставляет подтверждённо stale status после успешного ответа;
- [ ] corrupt value обрабатывается контролируемо и не раскрывается клиенту;
- [ ] TTL, invalidation и stampede trade-offs защищены измерениями.

### Усложнение после первой версии

Добавь singleflight/lock только для hot miss и докажи concurrent test, что repository reads ограничены без превращения lock в correctness guarantee.

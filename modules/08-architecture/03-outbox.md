# Transactional outbox

В одной DB transaction запиши business state и outbox record. Publisher читает outbox, публикует, помечает доставленным; crash создаёт duplicate, поэтому consumer idempotent. Outbox исправляет dual-write, не создаёт exactly-once.

## Где это применяется в реальном backend

1. **Заказ создан — отправить событие в Kafka** — task/order и outbox row фиксируются одной PostgreSQL transaction. Если писать в broker отдельно, crash между двумя writes оставляет системы в разных состояниях.
2. **Оплата подтверждена — запустить уведомление** — publisher повторяет непомеченные records после сбоя. Повторная доставка нормальна; без event ID и idempotent consumer письмо или списание повторятся.
3. **Интеграция с медленным внешним API** — request handler быстро фиксирует намерение, worker доставляет позже. Outbox не гарантирует SLA внешней стороны и требует lag/age monitoring, retry policy и DLQ/ручного разбора.

## Глубокое погружение

Главный инвариант: business mutation и append immutable outbox record принадлежат одной DB transaction. Publisher владеет claim/lease, публикацией и отметкой; consumer — дедупликацией side effect. `FOR UPDATE SKIP LOCKED` позволяет workers делить backlog, но crash после publish до mark создаёт duplicate. Polling даёт DB load и latency; batch size удерживает locks/connection, cleanup влияет на table/index bloat. Edge cases: poison payload, schema evolution, записи «in flight», clock-based lease, broker unavailable и бесконечный retry. Доказывай atomicity rollback-тестом, duplicates fault injection, конкурентный claim тестом и метриками oldest-event age, publish failures, batch size.

## Мини-проект

### Результат

Бизнес-сценарий: после создания task аналитика должна гарантированно получить `task.created`, даже если broker временно недоступен. Сначала расширь сценарий в существующем `labs/architecture/03-outbox-saga/starter/scenarios.md`, затем `project/backend-lab`: task и outbox row попадают в PostgreSQL атомарно, два publisher workers не владеют одной строкой одновременно, а тестовый consumer применяет один event ID к side effect только один раз. До кода опиши transaction boundary и точки crash; не добавляй Docker/Nginx/Kubernetes.

### Разрешённые знания

Go, context/concurrency, HTTP, PostgreSQL transactions/locks/indexes, Kafka delivery/idempotency, CAP/sharding как контекст и текущий transactional outbox. Используй только эти уже пройденные механизмы; не требуются Docker, Nginx или Kubernetes.

### Проверка

Из корня репозитория выполни `cd project/backend-lab && go test ./... -count=1 -race`; отдельный integration test с PostgreSQL проверяет rollback, concurrent claim и crash после publish до mark. Выведи `pending`, `oldest_age` и количество duplicate deliveries; если PostgreSQL runtime не предоставлен, integration proof остаётся явно незачтённым, а unit test не выдаётся за его замену.

### Критерии приёмки

- [ ] task и outbox row либо commit вместе, либо вместе отсутствуют;
- [ ] fault injection после publish создаёт duplicate delivery, но не duplicate side effect;
- [ ] конкурентный claim не теряет records и не держит network call внутри DB transaction без обоснования;
- [ ] Миша защищает at-least-once trade-off и план очистки backlog.

### Усложнение после первой версии

Добавь quarantine для poison event после ограниченного числа попыток, сохранив исходную ошибку и возможность ручного replay.

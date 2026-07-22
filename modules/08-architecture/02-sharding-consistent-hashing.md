# Sharding и consistent hashing

Shard key определяет locality и hotspots. Consistent hash уменьшает объём переназначения при добавлении node; virtual nodes сглаживают распределение. Не используйте хэш как excuse без стратегии rebalance и replication.

## Где это применяется в реальном backend

1. **Распределение профилей по PostgreSQL shards** — `user_id` даёт стабильный owner и локальные запросы. Поиск всех пользователей по городу становится scatter-gather; неверный key переносит стоимость в каждый запрос.
2. **Роутинг cache keys между Redis nodes** — consistent hashing уменьшает cache miss storm при добавлении узла. Он не реплицирует данные и не спасает от потери node без отдельной стратегии replicas.
3. **Partition key событий заказа** — ключ `order_id` сохраняет порядок одного заказа и распределяет нагрузку. Популярный tenant создаёт hot partition, даже если среднее распределение выглядит ровным.

## Глубокое погружение

Обычный `hash(key) % N` переназначает большинство ключей при изменении `N`; ring переносит диапазоны к ближайшему clockwise owner. Virtual nodes уменьшают variance, но увеличивают metadata и число движущихся диапазонов. Инварианты: один key имеет детерминированного owner в конкретной версии topology; rebalance не должен делать запись бесхозной; replication placement не должен совпасть с единственной failure domain. Costs — hash lookup, routing table, cross-shard queries, data movement и dual-read/write на миграции. Edge cases: skew, один huge tenant, hash collision, node flapping и одновременная смена topology. Доказывай распределение histogram/max-to-mean, долей перемещённых ключей и тестом доступности во время rebalance.

## Мини-проект

### Результат

Бизнес-сценарий: каталог заказов растёт и должен распределяться по узлам без массового cold-cache при изменении topology. Расширь реальный starter `labs/architecture/02-consistent-hashing/starter`: размести 10 000 заказов на 3 узлах, затем добавь четвёртый, сравни modulo и consistent hash, зафиксируй ADR выбора shard key и failure experiment с hot tenant. Не реализуй replication или сетевой cluster: scope ограничен детерминированным routing и измерением распределения.

### Разрешённые знания

Все предыдущие темы, включая CAP, Go maps/slices, tests и текущие sharding/consistent hashing. Реальная distributed DB и Kubernetes не нужны.

### Проверка

Из корня репозитория выполни `cd labs/architecture/02-consistent-hashing/starter && go test ./... -count=1`; тест или команда симулятора печатает нагрузку каждого узла, `max/mean` и процент переназначенных ключей до/после добавления узла. Отдельный сценарий с 30% заказов одного tenant показывает hotspot.

### Критерии приёмки

- [ ] одинаковые key и topology всегда дают одинакового owner;
- [ ] сравнение modulo и ring опирается на измеренный процент движения, а не на описание;
- [ ] ADR разбирает locality, scatter-gather и hot key;
- [ ] тесты покрывают пустую topology, добавление/удаление узла и virtual nodes.

### Усложнение после первой версии

Смоделируй draining узла: новые ключи туда не попадают, существующие перемещаются порциями, а метрика показывает незавершённый rebalance.

# Cache-aside

Read: cache → source on miss → cache with TTL. Write: source-of-truth, затем invalidate/update cache. Stale cache — ожидаемый риск: выбери TTL и invalidation из требований, не из удобства. Не кэшируй ошибки бесконечно.

## Где это применяется в реальном backend

1. **GET order** — cache hit снимает read с PostgreSQL; cache miss не должен возвращать пустоту как найденный order.
2. **Order update** — DB commit предшествует invalidation; обратный порядок создаёт окно старого cache.
3. **Negative cache** — кратко защищает от запросов неизвестных ids; долгий TTL скрывает недавно созданный order.

## Глубокое погружение

Cache-aside оставляет consistency приложению. Между DB read и cache set возможна гонка с writer; TTL ограничивает, но не исключает staleness. Costs: serialization, extra hop, memory, invalidations. Failures: Redis down, stale fill, poisoned value, stampede. Доказывай fake clock/controlled interleaving, hit/miss counters и fallback test.

## Мини-проект

### Результат

Расширь order store cache-aside чтением: hit/miss, DB update с invalidation и доступность чтения при отказе Redis.

### Разрешённые знания

Redis types/TTL и все предыдущие PostgreSQL/HTTP/Go темы.

### Проверка

Integration/unit tests hit, miss, stale invalidation, corrupt value, Redis unavailable; метрики hit/miss.

### Критерии приёмки

- [ ] PostgreSQL остаётся truth;
- [ ] update/invalidation order обоснован;
- [ ] cache failure не превращает корректные данные в ложный успех.

### Усложнение после первой версии

После MVP добавить bounded negative caching и проверить создание после miss.

# Cache-aside

Read: cache → source on miss → cache with TTL. Write: source-of-truth, затем invalidate/update cache. Stale cache — ожидаемый риск: выбери TTL и invalidation из требований, не из удобства. Не кэшируй ошибки бесконечно.

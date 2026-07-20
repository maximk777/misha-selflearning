# Stampede и locks

На одновременном miss много запросов бьют в source. Ограничивай singleflight/lock, jitter TTL и stale-while-revalidate. Наивный distributed lock опасен при pause и expiry: lock не превращает операцию в транзакцию без fencing/idempotency.

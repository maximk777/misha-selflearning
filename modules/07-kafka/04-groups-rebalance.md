# Groups и rebalance

В group partition назначается одному active consumer. Join/leave вызывает rebalance и временно останавливает consumption; обработчик должен переживать повтор записи после rebalance. Число consumer больше partitions не ускоряет topic.

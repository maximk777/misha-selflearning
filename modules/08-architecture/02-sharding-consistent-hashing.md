# Sharding и consistent hashing

Shard key определяет locality и hotspots. Consistent hash уменьшает объём переназначения при добавлении node; virtual nodes сглаживают распределение. Не используйте хэш как excuse без стратегии rebalance и replication.

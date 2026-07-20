# Compose networking

Compose service name — DNS имя в сети; host port нужен только хосту. `depends_on` не заменяет readiness: используйте healthcheck/wait. Port Redis 6379, Kafka-compatible broker 9092 и UI-порты документируй, но не хардкодь в production.

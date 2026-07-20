# Retries и DLQ

Transient error: ограниченный retry с backoff; poison message: причина + metadata в DLQ. Не ретраить бесконечно и не коммитить исходный offset до маршрутизации в DLQ. Наблюдай retry count и возраст сообщения.

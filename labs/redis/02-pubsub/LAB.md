# Pub/Sub

В одном terminal: `redis-cli SUBSCRIBE course`; в другом: `redis-cli PUBLISH course hello`. Останови subscriber, снова выполни PUBLISH и предскажи потерю сообщения. Не оставляй процесс: `Ctrl-C` завершает SUBSCRIBE. Перескажи, почему для durable обработки нужен Stream.

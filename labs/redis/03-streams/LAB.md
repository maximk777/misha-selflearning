# Streams

CLI: `redis-cli XADD orders '*' id 1`, `redis-cli XGROUP CREATE orders workers 0 MKSTREAM`, `redis-cli XREADGROUP GROUP workers misha COUNT 1 STREAMS orders '>'`, затем `XACK orders workers <id>`. Поломка: не делай XACK и наблюдай pending через `XPENDING orders workers`. Перескажи at-least-once и idempotency.

# Pub/Sub

Pub/Sub доставляет только подключённым подписчикам: offline message потерян. Это полезно для эфемерных событий, не для гарантированной обработки. CLI: `SUBSCRIBE topic` в одном окне, `PUBLISH topic hi` в другом; останови subscriber и наблюдай потерю.

## Где это применяется в реальном backend

1. **UI refresh hint** — online instances получают order-changed; offline subscriber пропускает hint и обязан перечитать truth.
2. **Cache invalidation hint** — событие ускоряет сброс cache; потеря не должна нарушать correctness благодаря TTL.
3. **Live operational notification** — low-latency broadcast приемлем без replay; billing/job processing — нет.

## Глубокое погружение

Publish доставляет подключённым subscribers без durable log/ack. Connection subscriber переходит в push mode, slow consumer ограничен buffers/network. Reconnect требует новой подписки и гарантированно оставляет непокрытый gap: сообщения периода disconnect потеряны. Сам reconnect старые сообщения не повторяет; duplicate возможен только из-за app-level retry или повторной публикации. Доказывай disconnect/reconnect experiment и сверку с PostgreSQL truth.

## Мини-проект

### Результат

Добавь поверх order store/cache эфемерное уведомление `order changed`, которое ускоряет refresh, но потеря которого не ломает чтение.

### Разрешённые знания

Пройденные Redis cache темы, HTTP/PostgreSQL/concurrency; Streams следующей темы не требуются.

### Проверка

Два subscribers получают event; отключённый пропускает его, после reconnect корректный state читается из store.

### Критерии приёмки

- [ ] offline loss воспроизведён;
- [ ] Pub/Sub не source of truth и не job queue;
- [ ] reconnect/cancel lifecycle обработан.

### Усложнение после первой версии

Добавить медленного subscriber и наблюдать latency/backpressure boundary.

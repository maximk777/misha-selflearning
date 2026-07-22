# Nginx reverse proxy

Nginx принимает client connection и проксирует upstream, выставляя `X-Request-ID`, `X-Forwarded-*`. Timeout защищает занятые resources; retry небезопасен для non-idempotent POST без ключа. Config проверяй `nginx -t` внутри container.

## Где это применяется в реальном backend

1. **Единая HTTP-точка входа** — Nginx принимает public port и направляет `/api` в Go service. Неверный `proxy_pass` или path rewrite меняет URL и даёт неожиданный 404.
2. **Передача client context** — proxy формирует trusted `X-Forwarded-For`, proto и request ID. Приложение не должно доверять тем же headers напрямую от произвольного клиента.
3. **Ограничение медленного upstream** — connect/read/send timeouts освобождают proxy resources. Слишком короткий timeout обрывает валидную работу, а автоматический retry POST может удвоить side effect.

## Глубокое погружение

Reverse proxy разделяет client и upstream connections: buffering, keepalive и timeout существуют на обеих границах. URI normalization и trailing slash у `location/proxy_pass` имеют observable semantics. Ownership public headers и client IP начинается на доверенной proxy boundary; hop-by-hop headers не передаются как end-to-end. Costs — extra hop, buffers, file descriptors, worker connections и log volume. Edge cases: client disconnect, large body, websocket/streaming buffering, upstream DNS change, spoofed forwarding header. Failures доказывай `nginx -t`, access/error logs с одним request ID, timeout experiment и запросом с поддельным header.

## Мини-проект

### Результат

Бизнес-сценарий: клиенты Task API должны ходить через единую public boundary с едиными headers, timeout и access log. Расширь cumulative `project/backend-lab/compose.yaml`: наружу опубликован только Nginx, `/api/` проксируется в API, `/healthz` имеет осознанную semantics, а request ID проходит до application log. Сначала нарисуй две connection boundaries и бюджеты времени; load balancing и TLS пока не добавляй.

### Разрешённые знания

Docker, Compose networking/health, HTTP/config/logging и текущий Nginx reverse proxy. Алгоритмы load balancing и TLS будут позже и в обязательную версию не входят.

### Проверка

Из корня репозитория выполни `docker compose -f project/backend-lab/compose.yaml config`, `docker compose -f project/backend-lab/compose.yaml up -d --build` и `docker compose -f project/backend-lab/compose.yaml exec nginx nginx -t`; сделай `curl -i http://localhost/api/...`, запрос с собственным `X-Request-ID` и delayed request для read timeout. Сверь proxy/application logs по ID и убедись, что API port не опубликован на host.

### Критерии приёмки

- [ ] path и status не меняются случайно при proxying, forwarding headers заданы явно;
- [ ] внешний request ID не позволяет подделать trusted audit identity;
- [ ] timeout даёт воспроизводимый status/log и не запускает скрытый retry unsafe request;
- [ ] конфигурация проходит `nginx -t`, а приложение не публикуется отдельно.

### Усложнение после первой версии

Добавь streaming response и сравни поведение с buffering on/off по time-to-first-byte и памяти proxy.

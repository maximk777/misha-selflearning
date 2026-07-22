# TLS, timeouts, retries

TLS завершается на ingress/proxy с управляемым сертификатом. У каждой границы есть deadline; retry имеет лимит, backoff и только для безопасных операций. Сумма downstream timeout должна укладываться в upstream deadline.

## Где это применяется в реальном backend

1. **TLS termination на edge** — proxy проверяет certificate/key и передаёт корректный scheme приложению. Шифрование до edge не защищает внутренний hop, если threat model требует end-to-end/mTLS.
2. **Бюджет запроса API → DB/external API** — Go application владеет общим context deadline и timeout своих downstream calls, оставляя время вернуть контролируемую ошибку. Независимые таймауты длиннее client deadline продолжают бесполезную работу.
3. **Разные owners повторов** — Nginx может переключить один idempotent request на следующий upstream только по явно заданным `proxy_next_upstream` conditions/tries/timeout; Go client отдельно владеет bounded attempts, exponential backoff и jitter при вызове внешнего API. Если оба слоя повторяют одну операцию без общего budget, amplification перемножается; POST без idempotency key не повторяется.

## Глубокое погружение

TLS handshake подтверждает identity certificate chain и согласует session keys; expiry, hostname/SAN и trust store являются частью доступности. Deadline — абсолютная граница, timeout — бюджет конкретного этапа. Nginx upstream retry не выполняет application backoff/jitter: его owner — edge config, лимиты задаются `proxy_next_upstream_tries` и `proxy_next_upstream_timeout`. Retry Go HTTP client — отдельный owner с bounded attempts, exponential backoff, jitter и проверкой остатка context deadline. Для одной operation должен быть один согласованный retry budget, иначе два слоя перемножают attempts. Health имеет startup/readiness/liveness смыслы даже до Kubernetes: «process жив» не равно «готов принимать traffic». Costs — handshake CPU/latency, open connections, queued work и extra traffic. Edge cases: partial response, client cancellation, stale certificate, retry после side effect до response, synchronized probes. Доказывай certificate inspection, отдельные attempt counters edge/application и fault experiments с delay/503.

## Мини-проект

### Результат

Бизнес-сценарий: публичный Task API шифрует traffic и не превращает краткий отказ replica или внешнего API в retry storm. Расширь cumulative `project/backend-lab/compose.yaml`: добавь локальный TLS и end-to-end deadline budget. Отдельно опиши и настрой (1) Nginx upstream failover с owner `nginx`, `tries/timeout` без backoff и только для безопасной операции; (2) Go client retry с owner application, bounded attempts, backoff/jitter и общим context budget. Проведи experiments «slow upstream» и «одна replica отвечает 503»; не разрешай обоим слоям повторять один и тот же hop.

### Разрешённые знания

Вся последовательность Docker → Compose → Nginx → load balancing, ранее изученные context/timeouts/idempotency и текущие TLS/health/retries. Kubernetes пока не используется.

### Проверка

Из корня репозитория выполни `docker compose -f project/backend-lab/compose.yaml config`, `docker compose -f project/backend-lab/compose.yaml up -d --build` и `docker compose -f project/backend-lab/compose.yaml exec nginx nginx -t`; `openssl s_client -connect localhost:443 -servername localhost` или `curl -vk https://localhost/...` подтверждает certificate/hostname. Два отдельных fault experiments печатают `edge_attempts`, `application_attempts`, final status и total latency: Nginx переключает upstream не больше своего `tries`, Go client показывает bounded attempts с наблюдаемыми backoff/jitter, сумма укладывается в общий deadline; POST без idempotency key не повторяется.

### Критерии приёмки

- [ ] TLS config проверена, private key не попал в image/history/repository;
- [ ] budgets образуют осознанную цепочку и cancellation доходит до Go handler;
- [ ] Nginx upstream retry и Go client retry имеют разных owners, отдельные лимиты и не накладываются на один hop;
- [ ] backoff/jitter проверены только у application/client retry, а Nginx failover ограничен `tries/timeout` и безопасной семантикой;
- [ ] failure experiment показывает число исходных requests, edge/application attempts и общий storm limit.

### Усложнение после первой версии

Добавь общий retry budget на окно времени и покажи по отдельным edge/application counters, как degraded dependency перестаёт получать бесконечное усиление traffic.

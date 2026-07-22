# Observability и security

Логи связывай request/correlation ID, метрики — rate/error/latency/saturation, traces — границы вызовов. Secrets не в image/репозитории; least privilege, authn/authz, input validation и TLS — baseline. Redact PII прежде чем логировать.

## Где это применяется в реальном backend

1. **Диагностика медленного POST заказа** — request ID связывает proxy, handler, DB и event publisher; histogram показывает p95/p99. Лог каждого body увеличит стоимость и может утечь PII.
2. **Разбор backlog outbox** — метрики pending/oldest age и trace span отличают медленный broker от DB claim. Один средний latency скрывает хвост и не объясняет причинность.
3. **Защита admin endpoint** — authn устанавливает identity, authz проверяет действие и resource, audit log фиксирует решение. TLS без authorization не запрещает обычному пользователю чужую операцию.

## Глубокое погружение

Telemetry имеет ownership и cardinality budget: logs — дискретные события, metrics — агрегаты, traces — причинный путь. Correlation ID нельзя слепо доверять от внешнего клиента; его валидируют/генерируют на границе и передают через context. RED/USE инварианты требуют стабильных labels; `user_id` как metric label взрывает cardinality. Security boundaries: input validation не заменяет authorization, secret redaction должна происходить до sink, least privilege проверяется реальными credentials. Edge cases: sampling теряет редкий trace, clock skew, log injection, panic с токеном, метрика успеха при ошибочном статусе. Доказывай failure injection, golden log без PII, metrics assertions и audit trail denied/allowed решений.

## Мини-проект

### Результат

Бизнес-сценарий: on-call должен связать медленный или запрещённый запрос с repository/outbox и не раскрыть данные пользователя. Расширь `project/backend-lab`: один request проходит с безопасным request ID через HTTP, repository и outbox; endpoint отдаёт RED-метрики, health различает процесс и зависимости, а protected operation имеет authz/audit. Составь threat/observability ADR и проведи два failure experiments; не добавляй сторонний telemetry stack, если его ещё нет в проекте.

### Разрешённые знания

Все пройденные Go/HTTP/concurrency/data/messaging/architecture темы и текущие observability/security. Можно использовать standard library и уже имеющиеся зависимости проекта; инфраструктура следующего модуля не обязательна.

### Проверка

Из корня репозитория выполни `cd project/backend-lab && go test ./... -count=1 -race`; запросы с корректным/некорректным identity проверяют 2xx/4xx, искусственная задержка repository отражается в histogram/trace, поиск по логам находит request ID и не находит тестовый secret/email.

### Критерии приёмки

- [ ] логи, метрики и trace отвечают на разные вопросы и связаны одним безопасным ID;
- [ ] labels имеют ограниченную cardinality, PII/secrets redacted до записи;
- [ ] denied и allowed операции доказаны тестом и audit событием;
- [ ] защита объясняет sampling, health semantics, least privilege и стоимость telemetry.

### Усложнение после первой версии

Введи sampling traces и докажи, что error traces сохраняются всегда, а обычный traffic остаётся в заданном telemetry budget.

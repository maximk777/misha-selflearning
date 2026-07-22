# Workers, context и shutdown

> Уровень: `production` · Время: 45 минут · Практика: `cmd/api/main.go`

Результат: server завершает `Shutdown` по SIGINT/SIGTERM с bounded timeout. **Один вопрос:** почему shutdown context не равен request context? Запусти API, отправь Ctrl-C и объясни порядок stop intake → wait → timeout. Worker pool добавляется вокруг outbox claim с bounded concurrency и cancellation.

## Где это применяется в реальном backend

1. **Остановка API во время deploy** — процесс прекращает intake, ждёт in-flight handlers и завершает в grace period. Немедленный exit рвёт ответы; бесконечное ожидание блокирует rollout.
2. **Outbox publisher pool** — bounded workers ограничивают DB/broker concurrency и создают backpressure. Неограниченная goroutine на record исчерпывает pool/memory при backlog.
3. **Consumer rebalance/termination** — cancellation останавливает получение новой работы, текущий side effect завершается или безопасно повторяется. Commit offset после отмены без завершённого effect теряет сообщение.

## Глубокое погружение

Lifecycle имеет владельца root context и порядок: mark unready/stop intake → cancel producers → drain queues/in-flight → close clients/resources → exit by deadline. Channel закрывает sender; WaitGroup Add происходит до запуска goroutine; shutdown context не наследуется от уже отменённого request context. Costs — reserved workers/connections, queue memory, shutdown latency. Edge cases: blocked send, worker игнорирует context, panic без Done, signal during startup, second signal, DB call longer grace, claimed outbox lease. Доказывай goroutine leak/race tests, fake signal, in-flight barrier и измерение accepted/completed/abandoned jobs.

## Мини-проект

### Результат

Бизнес-сценарий: deploy Task API не должен терять принятые HTTP requests или events при SIGTERM и одновременно не может ждать бесконечно. Объедини API, outbox publisher и consumer lifecycle в `project/backend-lab`: bounded pools, root cancellation и один coordinator завершают процесс в документированном порядке. Проведи SIGTERM experiment с активным request и очередью events; не меняй persistence/delivery semantics прошлых checkpoints.

### Разрешённые знания

Только предыдущие checkpoints, goroutines/channels/select/context/sync/race/leaks/worker pool/graceful shutdown, DB/Kafka clients и infrastructure time budgets; новый orchestrator/framework не добавляй.

### Проверка

Из корня репозитория выполни `cd project/backend-lab && go test ./... -count=1 -race`, затем `docker compose -f project/backend-lab/compose.yaml up -d --build` и `docker compose -f project/backend-lab/compose.yaml kill -s SIGTERM api`. Lifecycle test проверяет запрет нового intake, завершение/повтор текущей работы и deadline; Compose logs/metrics показывают `accepted`, `completed`, `in_flight`.

### Критерии приёмки

- [ ] concurrency bounded относительно DB/broker capacity, ownership очередей/close однозначен;
- [ ] shutdown имеет один coordinator и конечный deadline;
- [ ] после SIGTERM нет silent lost side effect: завершение или повтор доказаны;
- [ ] tests/race detector не показывают race/leak, порядок защищён своими словами.

### Усложнение после первой версии

Добавь второй сигнал: первый запускает graceful drain, второй — немедленный controlled exit с логом оставшейся работы.

# Monolith или microservices

Начинай с modular monolith, когда команда и domain ещё меняются. Выделяй сервис при независимом scaling/deploy/ownership и чёткой границе данных. Microservices добавляют network failure, observability и контрактную стоимость.

## Где это применяется в реальном backend

1. **Новый продукт одной команды** — modular monolith сохраняет быстрые refactoring и одну transaction boundary. Раннее разделение превращает обычные function calls в сеть без доказанной пользы.
2. **Независимо растущий image processor** — отдельный deploy и scaling оправданы измеримой CPU-нагрузкой и другим release cadence. Выделение «потому что код большой» не доказывает service boundary.
3. **Платежи с отдельным ownership/compliance** — сервис изолирует данные и доступ команды. Цена — contracts, on-call, retries, distributed tracing и невозможность общей DB transaction.

## Глубокое погружение

Boundary должен совпадать с ownership бизнес-инварианта и данных; shared database оставляет distributed monolith. Синхронный call связывает availability/latency, async event — schema и eventual consistency. Costs измеряются deploy lead time, change failure rate, cross-boundary calls, p95 latency и on-call load, а не числом repositories. Edge cases: циклические зависимости, chatty APIs, version skew, shared library lockstep и distributed transaction. Production failure experiment отключает зависимость и показывает blast radius. Доказательство решения — context map, ADR с альтернативой modular monolith и метрики, по которым выделение или обратное объединение считается успешным.

## Мини-проект

### Результат

Бизнес-сценарий: команда решает, отделять ли независимо растущий event-processing контур от Task API. Создай scenario/ADR-файл `project/backend-lab/architecture-boundaries.md`: разметь `backend-lab` на модули, предложи одну кандидатную границу и сравни два варианта — оставить modular monolith или выделить. Построй измеримый request/event scenario и reasoned failure experiment; не ссылайся на отсутствующий `labs/architecture/05-boundaries`, не создавай второй production-сервис и не добавляй Kubernetes.

### Разрешённые знания

Все предыдущие темы курса, включая HTTP, data stores, messaging, CAP, outbox, saga и CDC, плюс текущая тема. Docker/Kubernetes не нужны для решения.

### Проверка

Из корня репозитория выполни `cd project/backend-lab && go test ./... -count=1`; dependency check или таблица в `architecture-boundaries.md` подтверждает отсутствие обратной зависимости между выбранными модулями. Reasoned simulation fake dependency с timeout показывает, какие endpoint/status/latency должны измениться; runtime failure proof требуется только если fake действительно реализован.

### Критерии приёмки

- [ ] boundary владеет данными и бизнес-инвариантом, а не только папкой;
- [ ] ADR содержит baseline metrics, ожидаемый выигрыш и цену эксплуатации;
- [ ] failure experiment показывает blast radius и degraded behavior;
- [ ] защита объясняет, почему выбранный вариант лучше альтернативы сейчас, а не навсегда.

### Усложнение после первой версии

Добавь сценарий version skew: старый caller и новый contract должны сосуществовать, а ADR — назвать срок и стоимость совместимости.

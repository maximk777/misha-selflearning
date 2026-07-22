# Saga

Saga разбивает distributed transaction на local transactions и compensations. Orchestration централизует flow; choreography распределяет реакции, но усложняет наблюдаемость. Compensation — новое бизнес-действие, не технический rollback.

## Где это применяется в реальном backend

1. **Оформление заказа: reserve → pay → ship** — saga связывает local transactions без общей DB transaction. Ошибка оплаты запускает release reservation; компенсация может сама временно не выполниться.
2. **Бронирование поездки** — отель и билет подтверждаются независимо, а отмена имеет штраф и бизнес-условия. Compensation не возвращает мир «как было», если внешний эффект уже увидел пользователь.
3. **Онбординг merchant** — choreography позволяет сервисам реагировать на события профиля. Рост числа событий скрывает flow и циклы; при сложных ветвлениях orchestration обычно легче наблюдать.

## Глубокое погружение

Saga хранит durable state machine: step, outcome, retry/compensation status и correlation ID. Инварианты задаются как допустимые конечные состояния, а не как мгновенная глобальная consistency. Orchestrator владеет переходами и проще показывает timeline, но становится важным компонентом; choreography снижает центральную связанность, но переносит её в contracts/events. Costs — промежуточные состояния, idempotency каждого шага, retries и operational recovery. Edge cases: late success после timeout, повтор compensation, необратимый side effect, reordered event, ручное вмешательство. Доказывай model/table tests переходов, fault injection на каждом шаге и проверку, что повтор сообщения не нарушает конечный инвариант.

## Мини-проект

### Результат

Бизнес-сценарий: оформление заказа связывает reserve, charge и confirm delivery, но общей distributed transaction нет. В существующем `labs/architecture/03-outbox-saga/starter/scenarios.md` добавь state-transition table, ADR orchestration против choreography и reasoned simulation минимум трёх точек отказа; для каждой укажи observable state, owner следующего действия и compensation. Ограничение: не придумывай готовый отдельный saga-service и не требуй отсутствующий lab runtime; внешние вызовы описывай через уже знакомые interfaces/fakes.

### Разрешённые знания

Все пройденные Go, concurrency, HTTP, PostgreSQL, Redis/Kafka, outbox и текущая saga. Новые infrastructure-механизмы не нужны.

### Проверка

Проверь `labs/architecture/03-outbox-saga/starter/scenarios.md` по `labs/architecture/03-outbox-saga/CHECK.md`: таблица должна воспроизводимо проводить happy path, failure/timeout после каждого local step и повтор команды с тем же operation ID. Если добавлена исполняемая модель в `project/backend-lab`, запусти из корня `cd project/backend-lab && go test ./... -count=1 -race`; без неё принимается только явно обозначенная reasoned simulation, а не выдуманный runtime proof.

### Критерии приёмки

- [ ] happy path приходит в единственное объявленное успешное состояние;
- [ ] каждая обратимая операция имеет idempotent compensation и тест повторного вызова;
- [ ] late success и failure compensation не маскируются как rollback;
- [ ] ADR связывает выбранный стиль с ownership, observability и стоимостью изменений.

### Усложнение после первой версии

Добавь необратимый шаг «отправить письмо» и измени порядок/политику так, чтобы защита решения явно учитывала этот side effect.

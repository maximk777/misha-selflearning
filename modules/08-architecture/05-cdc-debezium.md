# CDC и Debezium

CDC читает change log базы и публикует изменения. Debezium снимает polling-нагрузку, но требует schema evolution, snapshots и контроля lag. CDC event отражает изменение данных, не обязательно намерение домена.

## Где это применяется в реальном backend

1. **Наполнение поискового индекса из PostgreSQL** — CDC передаёт insert/update без изменения request path. Raw row change не сообщает бизнес-намерение и может раскрыть лишние columns.
2. **Миграция read model без остановки writes** — snapshot даёт базу, WAL changes догоняют хвост. Неверная граница snapshot/stream создаёт пропуски или повторы.
3. **Outbox Event Router** — Debezium читает только outbox rows и публикует domain events без polling publisher. Это не отменяет schema compatibility, idempotent consumer и контроль connector lag.

## Глубокое погружение

Connector читает ordered database log position, сохраняет offset и преобразует записи в event envelope. Ownership разделён: PostgreSQL хранит WAL/replication slot, connector — offset/snapshot state, consumer — интерпретацию и дедупликацию. Retained WAL расходует disk при отстающем connector; snapshot нагружает I/O и может длиться часами. Edge cases: DDL during snapshot, column rename/type change, replica identity, delete/tombstone, connector restart и offset loss. Production failures: заполненный disk из-за slot, schema deserialization error, silent lag. Доказательство — snapshot+concurrent-write experiment, restart from saved offset, compatibility test и метрики source timestamp lag/WAL retained bytes.

## Мини-проект

### Результат

Бизнес-сценарий: поисковая read model должна получить исходные tasks и последующие изменения без окна потери на границе snapshot/stream. Расширь CDC-card в существующем `labs/architecture/03-outbox-saga/starter/scenarios.md`: проведи reasoned simulation offset/snapshot/restart и зафиксируй ADR «raw table CDC или outbox CDC» для `backend-lab`. Если создаёшь исполняемый experiment, положи его в явно новый project path внутри `project/backend-lab`, а не ссылайся на отсутствующий `labs/architecture/04-cdc`; не требуй реальный Debezium runtime без предоставленного окружения.

### Разрешённые знания

PostgreSQL WAL/logical replication, Kafka producer/consumer/delivery, outbox, saga и текущие CDC/Debezium. Docker Compose ещё не обязателен: допускается предоставленный lab runtime или fake change log для unit model.

### Проверка

Проверь сценарий по существующему `labs/architecture/03-outbox-saga/CHECK.md`: запись, вставленная на границе snapshot/stream, должна появиться в read model, а restart от сохранённого offset не должен терять её. Для добавленного к `backend-lab` experiment выполни из корня `cd project/backend-lab && go test ./... -count=1`; реальный connector proof, если runtime предоставлен, отдельно фиксирует lag до и после искусственной паузы.

### Критерии приёмки

- [ ] snapshot и stream имеют явную границу offset без потерянной записи;
- [ ] duplicate/restart не создаёт duplicate side effect;
- [ ] schema change даёт контролируемую ошибку или совместимое чтение;
- [ ] ADR объясняет domain intent, data exposure, polling load и operational ownership.

### Усложнение после первой версии

Смоделируй несовместимое переименование column и предложи двухфазную evolution, подтверждённую тестом старого и нового consumer.

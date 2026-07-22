# Контракт сначала

> Уровень: `production` · Время: 45 минут · Практика: `project/backend-lab/api/openapi.yaml`

Результат: прочитай OpenAPI и назови 201, 200 replay, 400, 404 и 409. **Один вопрос:** почему повтор POST с тем же idempotency key не должен создавать вторую задачу? Измени required field в контракте, предскажи последствия и верни его. Перескажи, почему generated `ogen` слой отделён от domain.

## Где это применяется в реальном backend

1. **Параллельная работа frontend и backend** — OpenAPI фиксирует paths, schemas и errors, позволяет генерировать client/server glue. Контракт не доказывает бизнес-инварианты и может идеально описывать неверное поведение.
2. **Безопасное изменение public API** — compatibility review ловит удалённое поле, новый required input или изменённый status до deploy. Добавление optional field обычно безопаснее, но старые consumers всё равно нужно проверить.
3. **Idempotent create после client timeout** — key и ответы 201/200 replay задают observable semantics повтора. Один и тот же key с другим payload должен конфликтовать, иначе key скрывает ошибку клиента.

## Глубокое погружение

Contract владеет wire representation, а domain — значением операции; generated `ogen` types/adapters не должны становиться domain model. Инварианты включают стабильный error envelope, однозначные status codes и совместимость старого consumer с новой schema. Стоимость codegen — build step, version pinning и regeneration diff; ручной код — drift. Edge cases: nullable против absent, unknown enum, duplicate headers, large body, content type и key reuse с другим payload. Production failure — rollout server раньше client или наоборот. Доказывай schema validation, contract tests, generated-code clean diff и compatibility сценариями old/new.

## Мини-проект

### Результат

Бизнес-сценарий: frontend и backend должны независимо реализовать create/get/complete Task API с одинаковой семантикой retry и ошибок. Начни финальное расширение `project/backend-lab`: уточни OpenAPI для idempotency replay и стабильных 400/404/409. Сначала перечисли domain outcomes, затем отобрази их в HTTP; на этом checkpoint не реализуй PostgreSQL/Redis/Kafka adapters и не редактируй generated код вручную.

### Разрешённые знания

Весь minimum path модулей 01–09, текущие OpenAPI/ogen и существующий `backend-lab`. Используй только уже выбранный generator/toolchain проекта; новые framework не добавляй.

### Проверка

Из корня репозитория выполни `make -C project/backend-lab generate` и `make -C project/backend-lab test`; при отсутствии сети используй уже pinned generator cache либо честно отметь generation proof незавершённым. Сравни generated diff; contract tests проверяют 201, replay 200, 400, 404, 409 и key reuse с другим payload.

### Критерии приёмки

- [ ] OpenAPI не противоречит существующим domain outcomes и не содержит двусмысленных error bodies;
- [ ] повторная генерация идемпотентна, generated код не редактировался вручную;
- [ ] backward-compatible и намеренно breaking изменения различаются тестом/объяснением;
- [ ] Миша защищает границу contract/generated adapter/domain.

### Усложнение после первой версии

Добавь optional `description` так, чтобы старый client продолжал работать, и докажи совместимость contract test.

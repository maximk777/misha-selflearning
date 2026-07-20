# Outbox and saga cards

- Заказ записан, publish упал: DB transaction содержит order+outbox; publisher повторит запись, consumer idempotent.
- Оплата прошла, резерв не прошёл: saga запускает compensation/refund, а не DB rollback другой системы.
- CDC snapshot догоняет log: consumer учитывает schema/version и duplicate events.

# Ordering и idempotency

Key сохраняет порядок одной сущности в одной partition. Idempotency key хранится с результатом side effect; duplicate возвращает уже принятый результат. Нельзя гарантировать порядок между partitions — меняй модель данных или компенсируй.

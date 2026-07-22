# Индексы

B-tree — равенство, диапазон и сортировка; hash — равенство; GIN — составные значения, arrays, JSONB, full-text. Composite index работает прежде всего по левому префиксу. Partial index хранит только строки условия. Индекс ускоряет чтение, но занимает место и удорожает запись.

## Где это применяется в реальном backend

1. **Поиск order по id** — B-tree ускоряет equality; primary key уже создаёт индекс.
2. **Лента customer orders** — composite `(customer_id, created_at)` поддерживает filter/order; неправильный порядок колонок бесполезен.
3. **Pending queue** — partial index уменьшает размер; запрос обязан логически соответствовать predicate.

## Глубокое погружение

B-tree хранит ordered keys/page tree; проверка visibility может потребовать чтения heap. Каждый индекс удорожает INSERT/UPDATE, занимает cache/disk и создаёт bloat. Edge cases: low selectivity, expression mismatch, NULL, left prefix. На текущем уроке доказывай пользу повторяемыми latency measurements, размером индекса и write benchmark; чтение `EXPLAIN` — тема следующего урока и здесь только будущий контекст.

## Мини-проект

### Результат

Подбери для order store максимум два индекса под три заданных query patterns и докажи пользу, не индексируя всё.

### Разрешённые знания

DDL/DML и пройденные PostgreSQL темы. `EXPLAIN` будет следующим уроком и в обязательную работу здесь не входит.

### Проверка

Загрузить representative data, многократно сравнить query latency и index size до/после; отдельно измерить write cost.

### Критерии приёмки

- [ ] каждый индекс привязан к query;
- [ ] redundant index отсутствует;
- [ ] read gain и write/storage cost объяснены.

### Усложнение после первой версии

Изменить distribution статусов и проверить, остаётся ли partial index полезен.

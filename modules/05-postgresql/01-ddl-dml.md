# DDL и DML

DDL меняет схему: `CREATE/ALTER/DROP`. DML меняет или читает строки: `SELECT/INSERT/UPDATE/DELETE`. Ограничения `NOT NULL`, `CHECK`, `UNIQUE`, `FOREIGN KEY` защищают данные ближе к источнику истины. Практика: `labs/postgres/01-ddl-dml/`.

Этот модуль переносит знакомый flow Task API на домен заказов: HTTP-граница остаётся знакомой, а новым source of truth становится PostgreSQL order store.

## Где это применяется в реальном backend

1. **Order schema** — constraints защищают amount/status/customer; nullable-by-default пропускает некорректный заказ.
2. **Создание заказа** — INSERT фиксирует факт, UPDATE меняет состояние; update без predicate повреждает весь набор.
3. **Связь order/items** — FK запрещает сирот; каскадное удаление опасно без бизнес-решения.

## Глубокое погружение

Constraints исполняются БД для каждого writer и образуют invariant source of truth. DDL берёт locks и может переписывать table; modifying DML (`INSERT/UPDATE/DELETE`) создаёт row versions и WAL, тогда как обычный `SELECT` tuple/WAL не пишет. Edge cases: NULL в CHECK/UNIQUE, FK order, concurrent writes. Доказывай SQL negative cases, schema inspection и повторяемой миграцией.

## Мини-проект

### Результат

Создай основу order store: schema и набор SQL-операций create/get/update с намеренно проверяемыми constraints; структуру решения сначала предложи сам.

### Разрешённые знания

SQL basics, текущая тема и ранее пройденные HTTP/Go темы; transactions следующих уроков не требуются.

### Проверка

Выполнить schema дважды в disposable DB, positive/negative SQL cases и `go test ./...`, если есть Go adapter.

### Критерии приёмки

- [ ] valid order сохраняется/читается;
- [ ] invalid amount/status/FK отвергаются БД;
- [ ] DDL/DML и выбранные constraints объяснены.

### Усложнение после первой версии

Добавить безопасное изменение schema для уже существующих строк.

# 12 недель до Go backend middle

Старт: **21 июля 2026**. Контрольная дата: **21 октября 2026**. Базовый темп: **10–12 часов в неделю**.

Этот календарь — маршрут, а не автоматическая отметка о прохождении. Завершённым считается только материал с доказательством: выводом команды, тестом, работающим кодом или принятым устным объяснением. Полная карта тем — в [SYLLABUS.md](SYLLABUS.md), критерии готовности — в [READINESS.md](READINESS.md).

## Ритм недели

Минимальный путь — пять сессий: 2 ч теория и прогноз, 3 ч лабораторная, 2 ч самостоятельная починка, 2 ч повторение и мини-проверка, 1 ч ретроспектива. Это 10 часов. Ещё 2 часа — опциональный `advanced`-слой, разбор ошибки или mock-интервью. `advanced` нельзя использовать как причину не сдать обязательную практику.

Каждая сессия начинается с результата, который можно увидеть за 5–15 минут: запуск программы, тест, запрос, план запроса или контейнер. До подсказки Миша делает прогноз, запускает, намеренно ломает, чинит и пересказывает наблюдение.

| Неделя и даты | Фокус и минимальный результат | Обязательный путь, 10 ч | Опционально, до 2 ч | Контрольное доказательство |
|---|---|---|---|---|
| 1 · 21–26 июля | Окружение, модуль, синтаксис, функции | Диагностика; первый модуль; переменные, условия, циклы, функции; 2 мини-проверки | История GOPATH и устройство `go env` | `go test` первого модуля и объяснение пакета/модуля |
| 2 · 27 июля–2 августа | Данные и ошибки Go | Массивы, строки, slices, maps, структуры, методы, указатели, `defer`/ошибки/`panic`; 3 лабораторные | Внутренности `slice` и UTF-8 | Код с тестом, намеренная поломка `nil map`, мини-экзамен Go Start |
| 3 · 3–9 августа | API Go Core | Интерфейсы, композиция, видимость, пакеты, wrapping ошибок, generics | Ограничения generics и interface nil trap | Лабораторная с интерфейсом и `errors.Is`/`errors.As` |
| 4 · 10–16 августа | Качество и runtime | Table-driven tests, benchmark, fuzzing; память, escape, GC, `pprof` | Реализация map в используемой версии Go | Тестовый пакет, benchmark и короткий разбор `pprof`; mock №1 |
| 5 · 17–23 августа | Конкурентность I | Горутины, GMP, scheduler, channels, `select`, `context` | Preemption и блокирующие syscalls | Демонстрация блокировки/отмены и объяснение владельца закрытия channel |
| 6 · 24–30 августа | Конкурентность II | `WaitGroup`, mutexes, typed atomics/CAS, race/deadlock/leak, worker pool, semaphore, pipeline, shutdown | Тонкости memory ordering | `go test -race`, остановка worker pool; накопительный экзамен и mock №2 |
| 7 · 31 августа–6 сентября | HTTP | `net/http`, клиент/таймауты/body close, server/handler/middleware/JSON, validation/errors, `httptest`, cancellation/shutdown | Разбор HTTP keep-alive | Маленький HTTP-сервис с тестом и graceful shutdown |
| 8 · 7–13 сентября | PostgreSQL I | DDL/DML, constraints/FK/indexes, transactions, ACID, MVCC, isolation, locks | Table/advisory locks | Две сессии SQL и воспроизведённая блокировка |
| 9 · 14–20 сентября | PostgreSQL II | `FOR UPDATE`, `SKIP LOCKED`, atomic update, deadlock, B-tree/hash/GIN, selectivity, `EXPLAIN (ANALYZE, BUFFERS)` | WAL, физическая/логическая репликация, CDC/Debezium | План до/после индекса и deadlock diagnosis; mock №3 |
| 10 · 21–27 сентября | Redis и Kafka | Redis TTL/structures/cache-aside/stampede/PubSub/Streams; Kafka partitions, offsets, groups, delivery, retry/DLQ/idempotency/CDC | Distributed lock и exactly-once без магии | Cache-aside сценарий и объяснение повторной доставки |
| 11 · 28 сентября–4 октября | Архитектура и инфраструктура | CAP, sharding, outbox, saga, monolith vs microservices; Docker, Compose, Nginx, L4/L7/LB, Kubernetes overview | Consistent hashing и Kubernetes autoscaling | Несколько реплик за Nginx и разбор manifests |
| 12 · 5–11 октября | Итоговый сервис и интервью | OpenAPI/ogen, PostgreSQL, Redis, Kafka/outbox, worker pool/context, migrations/config/logging/correlation/health, tests, Compose | Анализ компромиссов проекта | Запуск интеграционного сценария, защита решения, mock №4 |

## Буфер до контрольной даты: 12–21 октября

Буфер не добавляет новую обязательную теорию. Он закрывает конкретные пробелы из доказательств, повторяет слабые темы по очереди повторения и проводит итоговый mock. Если минимум недели не выполнен, сначала восстановить минимальный путь, затем возвращаться к `advanced`.

## Принятие недели

1. Есть хотя бы один свежий артефакт: вывод команды, тест, работающий код или принятая защита.
2. Сдана мини-проверка из 3–5 вопросов; один вопрос может быть неожиданным из старого материала.
3. Ошибка, если была, записана как «неверная модель → правильная модель → проверка».
4. Следующая неделя выбирается по факту готовности, а не только по календарю.

# Итоговая лаборатория: Task API

Это учебный skeleton: по умолчанию он запускается без Docker на in-memory adapters. PostgreSQL, Redis и Kafka представлены compileable портами и SQL-контрактом; подключать реальные зависимости следует только на соответствующих checkpoints.

## Быстрый результат

```bash
cd project/backend-lab
go test ./...
go run ./cmd/api
```

В другом terminal: `curl http://127.0.0.1:8080/readyz`. Для создания задачи отправь JSON с `title` и `idempotency_key`; повтор с тем же ключом возвращает ту же задачу без второго event.

## Checkpoints

1. **Контракт.** Прогноз: чем 201 отличается от replay 200? Прочитай `api/openapi.yaml`, измени validation и проверь ответ.
2. **Domain/HTTP.** Прогноз: где request ID? Запусти test, сломай пустой title, почини и объясни error envelope.
3. **PostgreSQL/outbox.** Прогноз: почему task и outbox пишутся одной tx? Прочитай migration и `ClaimOutboxSQL`; до Compose не подменяй source of truth cache-ом.
4. **Redis/Kafka.** Прогноз: что произойдёт при cache miss и повторе event? Добавь adapter только после выполнения unit tests; outage cache должен деградировать в repository read.
5. **Shutdown/container.** Прогноз: какие запросы успеет завершить `Shutdown`? Запусти server, отправь SIGINT, объясни timeout и non-root image.

После каждого checkpoint: прогноз → запуск → намеренная безопасная поломка → самостоятельная починка → пересказ → мини-экзамен. Не открывай готовые решения до своей попытки.

## Ogen

`cmd/api/main.go` содержит воспроизводимую pinned directive `ogen@v1.16.0`. Перед `make generate` проверь версию Go и доступность сети; generated output остаётся в `internal/ogen`, отдельно от handwritten adapters. В текущем starter HTTP adapter намеренно handwritten, чтобы unit tests не требовали generator или Docker.

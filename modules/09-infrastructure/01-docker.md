# Docker

Image — immutable template, container — process, volume — persistent data, network — DNS/изоляция. Pin image tag, не клади секрет в Dockerfile, используй multi-stage build. `docker logs` и healthcheck важнее «Up».

## Где это применяется в реальном backend

1. **Одинаковый запуск Go API на laptop и CI** — image фиксирует binary, CA certificates и runtime user. `latest` и зависимости из host делают результат невоспроизводимым.
2. **Уменьшение attack surface** — multi-stage оставляет compiler и source в builder, runtime запускается non-root. Маленький image не гарантирует безопасность, если binary или base уязвимы.
3. **Диагностика crash-loop процесса** — stdout/stderr, exit code и healthcheck дают наблюдаемое состояние. Контейнер со статусом `Up` может не принимать запросы или ждать недоступную DB.

## Глубокое погружение

Image — content-addressed набор read-only layers и config; container добавляет writable layer, namespaces/cgroups и один главный process. Build context и порядок Dockerfile команд влияют на cache и утечки: удалённый в следующем layer secret остаётся в предыдущем. PID 1 должен получать signals и завершать children; filesystem container не является persistent storage. Costs — размер transfer, cold start, duplicate layers, memory/CPU limits и image CVE maintenance. Edge cases: architecture mismatch, missing CA/timezone, static/dynamic linking, root-owned files, stale cache. Доказывай digest, `docker history`, размер, `docker inspect`, signal/shutdown experiment и запуск от non-root.

## Мини-проект

### Результат

Бизнес-сценарий: один и тот же Task API artifact должен одинаково запускаться у разработчика и в delivery pipeline. Упакуй существующий `project/backend-lab` API в reproducible multi-stage image: container стартует non-root, отвечает на health endpoint, пишет logs в stdout и корректно завершается по `docker stop`. Сначала предскажи, что останется в runtime layer; не добавляй Compose, proxy или cluster manifests.

### Разрешённые знания

Все предыдущие Go/backend темы, текущие image/container/layers/Dockerfile и shell-команды из лаборатории. Compose, Nginx и Kubernetes пока не использовать.

### Проверка

Из корня репозитория выполни `docker build -t backend-lab:local project/backend-lab`, затем `docker run --rm -p 8080:8080 --name backend-lab-local backend-lab:local` и `curl -f http://localhost:8080/readyz`; проверь `docker history backend-lab:local`, runtime user, image size и время `docker stop backend-lab-local`.

### Критерии приёмки

- [ ] build повторяем с pinned base и не включает `.git`, secrets или builder toolchain в runtime;
- [ ] API доступен, invalid config завершает process с понятной ошибкой;
- [ ] container работает non-root и принимает SIGTERM в bounded time;
- [ ] Миша объясняет layers, cache, writable layer и отличие image от container.

### Усложнение после первой версии

Собери image для второй CPU architecture через `buildx` и объясни, что именно проверено без запуска на реальном целевом host.

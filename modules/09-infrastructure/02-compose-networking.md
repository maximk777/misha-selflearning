# Compose networking

Compose service name — DNS имя в сети; host port нужен только хосту. `depends_on` не заменяет readiness: используйте healthcheck/wait. Port Redis 6379, Kafka-compatible broker 9092 и UI-порты документируй, но не хардкодь в production.

## Где это применяется в реальном backend

1. **Локальный stack API + PostgreSQL + Redis** — Compose создаёт общую сеть, а service names становятся DNS names. `localhost` внутри API указывает на сам API-container, не на host и не на DB.
2. **Сохранение данных между restart** — named volume переживает пересоздание container. Volume не является backup и `down -v` удаляет данные, поэтому destructive command должен быть явным.
3. **Ожидание готовности зависимости** — healthcheck показывает, что PostgreSQL принимает запросы. `depends_on` без health semantics задаёт порядок запуска, но не гарантирует готовность приложения.

## Глубокое погружение

Compose materializes services, user-defined networks, DNS aliases, volumes и environment/config. Port publishing создаёт host→container mapping; межсервисный traffic идёт на container port. DNS и IP меняются при recreate, поэтому clients используют имя и reconnect. Ownership persistence принадлежит volume и migration policy, не lifecycle container. Costs — локальные CPU/RAM, startup sequencing, hidden env drift и platform-specific mounts. Edge cases: port collision, stale volume schema, healthcheck к неправильной поверхности, app стартует раньше migration, container restart с новым IP. Доказывай `docker compose -f project/backend-lab/compose.yaml config`, `ps`, health status, DNS lookup из container, restart и checksum данных.

## Мини-проект

### Результат

Бизнес-сценарий: разработчику нужен одной командой воспроизводимый Task API с PostgreSQL и Redis, где restart API не теряет данные. Создай отсутствующий cumulative файл `project/backend-lab/compose.yaml` для stack `api + postgres + redis`: API обращается по service DNS, данные task переживают restart, readiness dependencies наблюдаема. Этот checkpoint становится владельцем compose-файла для всех следующих тем; не добавляй proxy, Kafka или Kubernetes в первую версию.

### Разрешённые знания

Docker из предыдущей темы, PostgreSQL/Redis, config/health и текущие Compose networks/volumes. Nginx, load balancing и Kubernetes ещё не требуются.

### Проверка

Из корня репозитория выполни `docker compose -f project/backend-lab/compose.yaml config`, `docker compose -f project/backend-lab/compose.yaml up -d --build` и `docker compose -f project/backend-lab/compose.yaml ps`; сделай HTTP create/read, затем `docker compose -f project/backend-lab/compose.yaml restart api` и повторный read. Из API-container проверь service DNS; останови PostgreSQL и зафиксируй health/degraded response. Завершение — `docker compose -f project/backend-lab/compose.yaml down` без `-v`.

### Критерии приёмки

- [ ] API не использует `localhost` для container dependencies и не публикует лишние DB/cache ports;
- [ ] healthchecks проверяют реальную готовность, а не существование process;
- [ ] task переживает restart API и PostgreSQL container без удаления volume;
- [ ] secrets не записаны в image или committed compose file; trade-off локального способа передачи объяснён.

### Усложнение после первой версии

Добавь one-shot migration service и докажи, что API не начинает принимать traffic при неуспешной migration.

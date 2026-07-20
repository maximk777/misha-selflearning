# Первичная проверка курса — 21.07.2026

Проверено локально на macOS arm64, Go 1.24.5, Docker 29.5.3 и Compose 5.1.4.

- `bash scripts/check-skills.sh` — семь repo-local skills валидны;
- `bash scripts/doctor.sh` — read-only диагностика завершена;
- `bash scripts/check-course.sh` — структура и Markdown-ссылки прошли;
- `GOCACHE=/tmp/misha-final-go-cache bash scripts/test-go-labs.sh` — все найденные Go-модули прошли;
- `go test -race ./...` в concurrency и HTTP starter — прошло;
- `docker compose -f deploy/compose.yaml config` — конфигурация валидна;
- Live integration выполнена: PostgreSQL, Redis и Kafka/Redpanda получили статус `healthy`; seed добавил 100 000 customers, 300 000 orders и 1 000 jobs, затем выполнил `ANALYZE`.
- Локальные порты `5432` и `55432` уже были заняты другими проектами, поэтому учебный PostgreSQL проверен через `POSTGRES_PORT=15432`. Это подтверждает работу предусмотренного port override.

Проверка репозитория не отмечала ученику ни одной пройденной темы. Стартовый статус остаётся onboarding.

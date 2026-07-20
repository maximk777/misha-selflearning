# Первичная проверка курса — 21.07.2026

Проверено локально на macOS arm64, Go 1.24.5, Docker 29.5.3 и Compose 5.1.4.

- `bash scripts/check-skills.sh` — семь repo-local skills валидны;
- `bash scripts/doctor.sh` — read-only диагностика завершена;
- `bash scripts/check-course.sh` — структура и Markdown-ссылки прошли;
- `GOCACHE=/tmp/misha-final-go-cache bash scripts/test-go-labs.sh` — все найденные Go-модули прошли;
- `go test -race ./...` в concurrency и HTTP starter — прошло;
- `docker compose -f deploy/compose.yaml config` — конфигурация валидна;
- PostgreSQL/Redis/Kafka live integration не запускалась в первичном прогоне, чтобы не скачивать образы без необходимости. Команда: `RUN_INTEGRATION=1 bash scripts/verify-all.sh`.

Проверка репозитория не отмечала ученику ни одной пройденной темы. Стартовый статус остаётся onboarding.

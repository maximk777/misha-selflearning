# PostgreSQL labs
Подними PostgreSQL: `bash scripts/deps-up.sh postgres && bash scripts/wait-deps.sh postgres && bash scripts/seed-postgres.sh`. Выполняй SQL: `docker compose -f deploy/compose.yaml exec -T postgres psql -U misha -d misha < файл.sql`.

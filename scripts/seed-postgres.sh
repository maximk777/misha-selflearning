#!/bin/sh
set -eu
root=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
docker compose -f "$root/deploy/compose.yaml" exec -T postgres psql -v ON_ERROR_STOP=1 -U misha -d misha < "$root/scripts/seed-postgres.sql"

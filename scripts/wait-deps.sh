#!/bin/sh
set -eu
root=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
service=${1:-postgres}
i=0
while [ "$i" -lt 60 ]; do status=$(docker inspect --format '{{if .State.Health}}{{.State.Health.Status}}{{else}}running{{end}}' "$(docker compose -f "$root/deploy/compose.yaml" ps -q "$service")" 2>/dev/null || true); [ "$status" = healthy ] && { printf '%s готов.\n' "$service"; exit 0; }; i=$((i+1)); sleep 1; done
printf '%s не стал healthy за 60 секунд.\n' "$service" >&2; exit 1

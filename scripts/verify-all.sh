#!/bin/sh
set -eu
root=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
cd "$root"

printf '1/5 shell syntax\n'
find scripts -type f -name '*.sh' -print | sort | while IFS= read -r file; do
	if sed -n '1p' "$file" | grep -q 'bash'; then bash -n "$file"; else sh -n "$file"; fi
done
printf '2/5 Codex skills\n'
bash scripts/check-skills.sh
printf '3/5 course structure\n'
bash scripts/check-course.sh
printf '4/5 Go modules\n'
GOCACHE=${GOCACHE:-/tmp/misha-verify-go-cache} bash scripts/test-go-labs.sh
printf '5/5 Compose config\n'
docker compose -f deploy/compose.yaml config >/dev/null

if [ "${RUN_INTEGRATION:-0}" = 1 ]; then
	bash scripts/deps-up.sh
	bash scripts/wait-deps.sh postgres
	bash scripts/wait-deps.sh redis
	bash scripts/wait-deps.sh kafka
	bash scripts/seed-postgres.sh
else
	printf 'Integration: пропущен; запусти RUN_INTEGRATION=1 bash scripts/verify-all.sh\n'
fi
printf 'Все обязательные проверки пройдены.\n'

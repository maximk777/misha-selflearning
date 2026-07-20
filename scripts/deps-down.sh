#!/bin/sh
set -eu
root=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
if [ "${1:-}" = "--volumes" ]; then printf 'Это удалит учебные данные. Введи DELETE: '; read answer; [ "$answer" = DELETE ] || exit 1; docker compose -f "$root/deploy/compose.yaml" down -v; else docker compose -f "$root/deploy/compose.yaml" down; fi

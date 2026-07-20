#!/bin/sh
set -eu
root=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
docker compose -f "$root/deploy/compose.yaml" up -d "$@"

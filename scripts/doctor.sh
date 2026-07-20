#!/bin/sh
set -u
printf 'ОС: %s\nАрхитектура: %s\n' "$(uname -s 2>/dev/null || printf '?')" "$(uname -m 2>/dev/null || printf '?')"
check() { name=$1; shift; if command -v "$name" >/dev/null 2>&1; then printf '%-16s найден: %s\n' "$name" "$($@ 2>&1 | sed -n '1p')"; else printf '%-16s не найден\n' "$name"; fi; }
check go go version
check git git --version
check docker docker --version
if command -v docker >/dev/null 2>&1 && docker compose version >/dev/null 2>&1; then printf '%-16s найден: %s\n' 'docker compose' "$(docker compose version 2>&1 | sed -n '1p')"; else printf '%-16s не найден\n' 'docker compose'; fi
check make make --version
check psql psql --version
check curl curl --version
printf 'Doctor ничего не устанавливал.\n'

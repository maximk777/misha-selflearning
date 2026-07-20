#!/usr/bin/env bash
set -euo pipefail

root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

while IFS= read -r module; do
	module_dir="$(dirname "$module")"
	relative_dir="${module_dir#"$root"/}"
	if [[ "$module_dir" == "$root" ]]; then
		relative_dir="."
	fi

	echo "Проверяем Go-модуль: $relative_dir"
	(
		cd "$module_dir"
		go test ./...
	)
done < <(find "$root" -type f -name go.mod -not -path "$root/.git/*" -print | LC_ALL=C sort)

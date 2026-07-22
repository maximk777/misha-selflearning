#!/bin/sh
set -eu

root=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
skills_dir="$root/.agents/skills"
tone_file="$root/.agents/references/tommy-shelby-tone.md"
learner_names='misha-onboarding misha-lesson misha-exam misha-review misha-progress misha-mock-interview misha-next misha-notes'
maintainer_names='misha-course-audit'
names="$learner_names $maintainer_names"
failed=0

require_phrase() {
	file=$1
	phrase=$2
	if ! grep -Fq "$phrase" "$file"; then
		printf '%s: нет обязательной фразы: %s\n' "$file" "$phrase" >&2
		failed=1
	fi
}

if [ ! -f "$tone_file" ]; then
	printf 'Нет общего tone reference: %s\n' "$tone_file" >&2
	failed=1
else
	require_phrase "$tone_file" '# Открытие урока'
	require_phrase "$tone_file" '# Открытие экзамена'
	require_phrase "$tone_file" '# Закрытие занятия'
	require_phrase "$tone_file" 'Я Томми Шелби.'
fi

require_phrase "$root/AGENTS.md" '.agents/references/tommy-shelby-tone.md'

if grep -R -Eiq 'Острые[[:space:]]+козырьки|Peaky[[:space:]]+Blinders' "$root/.agents"; then
	printf 'В агентских инструкциях найдена запрещённая связь образа с сериалом.\n' >&2
	failed=1
fi

for name in $names; do
	file="$skills_dir/$name/SKILL.md"
	if [ ! -f "$file" ]; then
		printf 'Нет skill manifest: %s\n' "$file" >&2
		failed=1
		continue
	fi
	if [ "$(sed -n '1p' "$file")" != '---' ] || ! grep -q '^---$' "$file"; then
		printf '%s: нет YAML frontmatter\n' "$file" >&2
		failed=1
	fi
	if ! grep -q "^name: $name$" "$file"; then
		printf '%s: неверное поле name\n' "$file" >&2
		failed=1
	fi
	if ! grep -q '^description: Use when' "$file"; then
		printf '%s: description должен начинаться с Use when\n' "$file" >&2
		failed=1
	fi
	case " $learner_names " in
		*" $name "*)
			require_phrase "$file" 'ровно один вопрос'
			require_phrase "$file" 'progress/PROFILE.md'
			require_phrase "$file" 'progress/STATUS.md'
			require_phrase "$file" 'progress/EVIDENCE.md'
			require_phrase "$file" 'progress/EXAMS.md'
			require_phrase "$file" 'progress/REVIEW_QUEUE.md'
			require_phrase "$file" 'Не показывай решения'
			require_phrase "$file" '.agents/references/tommy-shelby-tone.md'
			;;
		*)
			require_phrase "$file" 'progress/'
			require_phrase "$file" 'solutions/'
			;;
	esac
done

if [ "$failed" -ne 0 ]; then
	exit 1
fi

printf 'Проверка навыков наставника пройдена.\n'

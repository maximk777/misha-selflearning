# Stack, heap и escape analysis

> Уровень: `interview`
> Время: 60 минут
> Предпосылки: pointers
> Практика: `labs/go/09-memory-gc/LAB.md`
> Экзамен: `exams/go-core.md`, GC-6

## Результат за 5–15 минут

Выполни `go -C labs/go/09-memory-gc/starter build -gcflags=-m .` и найди сообщения компилятора про escape.

## Модель

Stack/heap — решение компилятора и runtime, а не прямое следствие синтаксиса `&`. Значение escape-ится, если его время жизни или доступ нельзя безопасно ограничить текущим stack frame. Сначала измерь профиль; не переписывай API ради угадывания allocation.

## Прогноз → опыт → поломка

**Один вопрос:** означает ли return `*payload`, что любой указатель всегда живёт на heap? Выполни анализ и сравни с выводом компилятора.

## Пересказ и сдача

**Один вопрос:** чем доказательство компилятора отличается от интуиции «указатель медленный»? Сохрани строку output и ответь на GC-6.

## Где это применяется в реальном backend

1. **Hot-path расчёт** — лишний heap allocation увеличивает GC work; pointer сам по себе не доказывает escape.
2. **Closure для правила** — захваченная переменная может жить дольше frame; blindly убранная closure ухудшает API без измерения.
3. **Возврат большого значения** — compiler может оптимизировать копии; преждевременный pointer повышает aliasing.

## Глубокое погружение

Compiler размещает значение там, где безопасно его lifetime; escape analysis консервативен и зависит от версии/inlining. Ownership и API correctness первичны, stack address безопасно возвращать именно потому, что compiler перенесёт объект при необходимости. Costs — allocation, GC scanning и locality. Edge cases: interface boxing, closures, formatting. Докажи `go -C project/order-report build -gcflags='-m=2' ./...`, `go -C project/order-report test -bench=. -benchmem ./domain` и profile; сообщение компилятора — evidence, не вечная гарантия.

## Мини-проект

### Результат

Продолжи `project/order-report/domain`: найди одну allocation hypothesis и проверь её, не меняя публичное поведение.

### Разрешённые знания

Предыдущие темы, escape analysis, benchmarks; без GC tuning и pprof.

### Проверка

Из корня репозитория: `go -C project/order-report test ./...`, `go -C project/order-report build -gcflags='-m=2' ./...`, `go -C project/order-report test -bench=. -benchmem ./domain`; сравни benchmark before/after.

### Критерии приёмки

- [ ] есть точная строка анализа и сценарий benchmark;
- [ ] изменение не основано на «pointer всегда быстрее»;
- [ ] ученик объясняет lifetime и trade-off aliasing.

### Усложнение после первой версии

Сравнить closure и явный параметр только измерением, не объявляя победителя заранее.

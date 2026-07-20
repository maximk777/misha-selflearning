# Модуль 02: Go Core

Здесь базовый синтаксис превращается в код, который можно тестировать, измерять и объяснять на собеседовании. Minimum path: interfaces, errors, generics, tests, benchmarks/fuzzing, memory/escape и `pprof`; внутренности map и GC отмечены `advanced` и не задерживают практику.

Лаборатории: `labs/go/07-interfaces`, `06-errors-panic`, `08-tests-bench-fuzz`, `09-memory-gc`, `10-pprof`. Сначала добейся наблюдаемого результата, а затем открывай deep dive.

| Тема | Доказательство | Банк |
|---|---|---|
| interfaces/errors/generics | test и объяснение API | GC-1–GC-4 |
| tests/benchmark/fuzz | `go test`, benchmark, bounded fuzz | GC-5–GC-6 |
| memory/GC/pprof | escape output или profile | GC-6 |

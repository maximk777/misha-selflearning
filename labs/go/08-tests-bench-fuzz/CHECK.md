# Проверка: tests, benchmark и fuzz

> Банк: `exams/go-core.md`, GC-5, GC-6

- [ ] `go test ./...` проходит с named subtests.
- [ ] Показан `go test -bench=BenchmarkAdd -benchmem` и прочитано `allocs/op`.
- [ ] Показан bounded `go test -fuzz=FuzzAddZero -fuzztime=2s`.
- [ ] Поломка `return 0` найдена test/fuzz и исправлена.

Экзаменатор задаёт GC-5 и GC-6; benchmark без baseline не считается production-выводом.

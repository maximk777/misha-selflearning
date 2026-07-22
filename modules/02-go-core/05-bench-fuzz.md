# Benchmarks и fuzzing

> Уровень: `interview`
> Время: 75 минут
> Предпосылки: table tests
> Практика: `labs/go/08-tests-bench-fuzz/LAB.md`
> Экзамен: `exams/go-core.md`, GC-6

## Результат за 5–15 минут

В starter выполни `go -C labs/go/08-tests-bench-fuzz/starter test -bench=BenchmarkAdd -benchmem` и увидь `allocs/op`. Затем запусти bounded fuzz: `go -C labs/go/08-tests-bench-fuzz/starter test -fuzz=FuzzAddZero -fuzztime=2s`.

## Модель

Benchmark измеряет конкретный сценарий на конкретной машине; сравнивай baseline и изменения, не делай глобальных выводов. Fuzzing генерирует входы и ищет нарушенную инварианту; найденный input превращается в регрессионный test.

## Прогноз → опыт → поломка

**Один вопрос:** сколько allocations ждёшь у сложения двух `int`? Запусти benchmark. Затем сломай `Add` на `return 0`, запусти test/fuzz и восстанови.

## Пересказ и сдача

**Один вопрос:** чего benchmark не доказывает о production latency? Покажи оба запуска и ответь на GC-6.

## Где это применяется в реальном backend

1. **Сравнение расчёта отчёта** — benchmark сравнивает версии на одном сценарии; он не моделирует сеть и production traffic.
2. **Поиск panic на строках** — fuzzing генерирует Unicode/пустые входы; invariant должен быть проверяемым, а не «не упало» без контракта.
3. **Контроль allocations** — `-benchmem` замечает копии slices; microbenchmark с устранённым результатом вводит в заблуждение.

## Глубокое погружение

Benchmark loop задаёт `b.N`; setup исключают таймером, результат удерживают наблюдаемым. Шум уменьшают несколькими запусками и сравнением distributions. Fuzzer хранит corpus и минимизирует crashing input; nondeterminism делает его нестабильным. Costs — CPU/время CI и ложные выводы из нерепрезентативных данных. Докажи baseline/changed outputs, regression test найденного input и отдельным `go -C project/order-report test -race ./...`.

## Мини-проект

### Результат

Продолжи `project/order-report/domain`: добавь benchmark агрегации заказов и fuzz invariant — результат нормализации валиден и повторный вызов не меняет его.

### Разрешённые знания

Предыдущие темы, tests, benchmarks, fuzzing; без concurrency/pprof.

### Проверка

Из корня репозитория: `go -C project/order-report test ./...`, `go -C project/order-report test -bench=. -benchmem ./...`, `go -C project/order-report test -fuzz=. -fuzztime=3s ./domain`.

### Критерии приёмки

- [ ] benchmark имеет representative input и baseline;
- [ ] fuzz проверяет сформулированную invariant;
- [ ] найденный дефект превращается в regression test.

### Усложнение после первой версии

Сравнить вариант с защитной копией slice и объяснить allocations/op.

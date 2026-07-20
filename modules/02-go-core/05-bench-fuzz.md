# Benchmarks и fuzzing

> Уровень: `interview`
> Время: 75 минут
> Предпосылки: table tests
> Практика: `labs/go/08-tests-bench-fuzz/LAB.md`
> Экзамен: `exams/go-core.md`, GC-6

## Результат за 5–15 минут

В starter выполни `go test -bench=BenchmarkAdd -benchmem` и увидь `allocs/op`. Затем запусти bounded fuzz: `go test -fuzz=FuzzAddZero -fuzztime=2s`.

## Модель

Benchmark измеряет конкретный сценарий на конкретной машине; сравнивай baseline и изменения, не делай глобальных выводов. Fuzzing генерирует входы и ищет нарушенную инварианту; найденный input превращается в регрессионный test.

## Прогноз → опыт → поломка

**Один вопрос:** сколько allocations ждёшь у сложения двух `int`? Запусти benchmark. Затем сломай `Add` на `return 0`, запусти test/fuzz и восстанови.

## Пересказ и сдача

**Один вопрос:** чего benchmark не доказывает о production latency? Покажи оба запуска и ответь на GC-6.

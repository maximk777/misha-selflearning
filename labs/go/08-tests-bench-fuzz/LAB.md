# Лабораторная: test, benchmark и fuzz

> Тема: `modules/02-go-core/04-table-tests.md`, `05-bench-fuzz.md`
> Уровень: `interview`
> Время: 35 минут
> Запуск из: `labs/go/08-tests-bench-fuzz/starter/`

## Результат за 5–15 минут

```bash
cd labs/go/08-tests-bench-fuzz/starter
go test ./...
go test -bench=BenchmarkAdd -benchmem
go test -fuzz=FuzzAddZero -fuzztime=2s
```

Ожидаются named subtests, строка benchmark с `allocs/op` и успешная короткая fuzz-сессия.

## До запуска: прогноз

**Один вопрос:** сколько allocations на операцию ожидаешь у `int`-сложения и почему это не production benchmark?

## Поломка и самостоятельная починка

Измени `Add` на `return 0`. Запусти обычный test, затем bounded fuzz; оба должны найти нарушение. Верни `left + right`.

## Пересказ

**Один вопрос:** какой найденный fuzz input обязан стать регрессией? Сдача — три команды, failure и пояснение `allocs/op`.

# Lifetime goroutine

Starter сравнивает обычный вызов и `go f()` без WaitGroup, channels, context и `time.Sleep`.

Из корня репозитория выполни:

```bash
cd labs/concurrency/goroutines/starter
go run . -mode sync -work 10000000
go run . -mode async -work 10000000
```

Повтори async-команду минимум пять раз и запиши фактический вывод. Процесс заканчивается при возврате из `main`, поэтому работа goroutine может не начаться или оборваться. Не добавляй `Sleep` или другие ещё не изученные primitives, чтобы «починить» пример: это наблюдение, а не runner.

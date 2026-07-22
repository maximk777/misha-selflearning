# Наблюдение за GMP scheduler

Starter — готовый observation harness. Внутри есть `sync.WaitGroup` только для того, чтобы все jobs успели завершиться. На этой теме не нужно понимать, менять или воспроизводить WaitGroup.

Из корня репозитория выполни:

```bash
cd labs/concurrency/gmp/starter
go run . -jobs 4 -work 1000000 -gomaxprocs 1
go run . -jobs 4 -work 1000000 -gomaxprocs 2
```

Каждую команду повтори три раза. До запуска спрогнозируй, что может измениться. Запиши порядок `start`/`finish`, но не принимай его или время одного запуска за гарантию scheduler.

Затем измени только flags `-jobs`, `-work`, `-gomaxprocs`. Код и WaitGroup не меняй.

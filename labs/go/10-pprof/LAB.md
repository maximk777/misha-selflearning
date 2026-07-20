# Лабораторная: pprof endpoint

> Тема: `modules/02-go-core/09-pprof.md`
> Уровень: `production`
> Время: 30 минут
> Запуск из: `labs/go/10-pprof/starter/`

## Результат за 5–15 минут

В terminal A:

```bash
cd labs/go/10-pprof/starter
go run .
```

В terminal B:

```bash
curl http://127.0.0.1:6060/work
curl http://127.0.0.1:6060/debug/pprof/
```

Ожидается `sum=4999950000` и HTML index pprof. После опыта останови server через Ctrl-C.

## До запуска: прогноз

**Один вопрос:** почему импорт `net/http/pprof` записан с `_`?

## Поломка и самостоятельная починка

Временно поменяй порт в `ListenAndServe` на `6061`, получи ошибку curl на 6060, верни 6060. Не оставляй сервер работать после опыта.

## Пересказ

**Один вопрос:** какой рабочий сценарий нужно зафиксировать до сравнения profiles? Сдача — два curl outputs и объяснение baseline.

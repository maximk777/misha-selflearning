# Профилирование через pprof

> Уровень: `production`
> Время: 75 минут
> Предпосылки: tests, memory
> Практика: `labs/go/10-pprof/LAB.md`
> Экзамен: `exams/go-core.md`, GC-6

## Результат за 5–15 минут

В одном terminal: `go -C labs/go/10-pprof/starter run .`. В другом: `curl http://127.0.0.1:6060/work`, затем открой `http://127.0.0.1:6060/debug/pprof/`. Останови сервер Ctrl-C.

## Модель

Profile отвечает на вопрос о наблюдаемой работе процесса: CPU, heap, goroutines и другие точки. Сначала зафиксируй baseline и representative action, затем меняй одну гипотезу и сравни. `pprof` не оправдывает оптимизацию без сценария.

## Прогноз → опыт → поломка

**Один вопрос:** появится ли заметная работа в профиле, если ни разу не вызвать `/work`? Проверь. Затем временно измени адрес порта, увидь ошибку curl и верни 6060.

## Пересказ и сдача

**Один вопрос:** что сравниваешь до оптимизации? Покажи endpoint и один profile URL, затем ответь на GC-6.

## Где это применяется в реальном backend

1. **CPU spike отчёта** — CPU profile показывает горячие функции; одна sample не доказывает причинность.
2. **Рост heap** — alloc/inuse profiles различают churn и удержание; profile без representative workload бесполезен.
3. **Зависшие процессы** — goroutine profile показывает stacks ожидания; он не заменяет понимание ownership.

## Глубокое погружение

Профиль — sampling/снимок наблюдаемого процесса, а не полный trace. Выбирают CPU, heap alloc/inuse, goroutine под конкретную гипотезу. Сбор имеет overhead и endpoint нельзя бездумно публиковать. Invariant исследования: baseline, одинаковая нагрузка, одно изменение, повторное измерение. Edge cases — inlining labels, short workload, debug endpoint с чувствительными данными. Докажи сохранёнными before/after profiles и тестами корректности.

## Мини-проект

### Результат

Продолжи `project/order-report/domain`: добавь диагностический benchmark, который создаёт CPU и allocation workload и позволяет сравнить один обоснованный вариант.

### Разрешённые знания

Весь Go Core, `pprof`, tests/benchmarks и стандартная библиотека; concurrency не требуется.

### Проверка

Из корня репозитория: `go -C project/order-report test ./...`, затем `go -C project/order-report test -run '^$' -bench=. -cpuprofile=cpu.out -memprofile=mem.out ./domain`. Открой profiles командами `go tool pprof project/order-report/cpu.out` и `go tool pprof project/order-report/mem.out`; запиши before/after наблюдение.

### Критерии приёмки

- [ ] выбран правильный profile под вопрос;
- [ ] workload воспроизводим и корректность сохранена;
- [ ] ученик защищает вывод без обещания production latency.

### Усложнение после первой версии

Проверить вторую гипотезу и честно оставить старый код, если профиль её не подтверждает.

# Semaphore и pipeline

Semaphore — buffered channel с capacity равной лимиту: acquire перед работой, release через `defer`. Pipeline: stage читает input и закрывает свой output, когда input исчерпан; fan-out делит работу, fan-in собирает. Каждый stage должен уметь завершиться при cancellation, иначе pipeline течёт.

## Где это применяется в реальном backend

1. **Лимит дорогой операции** — semaphore ограничивает активные вызовы внутри workers; acquire без cancellation зависает при shutdown.
2. **Pipeline validate→execute→format** — stages разделяют работу; один непрочитанный output блокирует upstream.
3. **Fan-out/fan-in batch** — несколько workers ускоряют независимые jobs, collector объединяет; порядок результатов не гарантирован.

## Глубокое погружение

Channel-semaphore имеет tokens/slots: acquire должен быть cancel-aware, успешный acquire всегда имеет ровно один release. Worker pool ограничивает число долгоживущих workers, semaphore — число одновременно находящихся в критической внешней операции; сочетание лимитов должно быть осмысленным. У каждой pipeline stage должен быть один owner lifecycle и выход при cancel, но конкретные channels, owners и close order ученик сначала проектирует сам. Наставник после первой попытки раскрывает по одному нарушенному инварианту, не готовую архитектуру. Для проекта фиксируется **partial-results policy**: ошибка одной job становится её outcome и не отменяет siblings. Cutoff cancellation — момент, когда runner наблюдает `ctx.Done()` и перестаёт подтверждать новые jobs; после него input отклоняется. Каждая подтверждённая до cutoff job получает ровно один outcome: `success`, `validation_error`, `execution_error` или `cancelled`. В tests injected gates заранее определяют, какие IDs успели закончить до cancel, поэтому exact set и допустимая категория каждого ID детерминированы. Happens-before channel передачи не решает ordering продукта. Costs — лишние goroutines/channels, head-of-line blocking, buffer memory. Failures: token leak, release без acquire, forgotten drain. Докажи atomic active/max-active, fault-injected stage, `-race` и завершение всех goroutines. Будущий HTTP dispatcher использует этот механизм только после изучения HTTP client.

## Мини-проект

### Результат

Предложи pipeline для проверки, выполнения инъецируемой job и форматирования с отдельным лимитом дорогой стадии. До кода сам определи stages, channels, owners, close/cancel order и обоснуй, где нужен semaphore; наставник не выдаёт схему до попытки.

### Разрешённые знания

Все предыдущие concurrency темы, worker pool, semaphore, pipelines; без HTTP.

### Проверка

Из корня: `cd project/concurrency-runner && go test -race -count=50 ./...` на limit, empty input, exact outcomes всех accepted-before-cutoff IDs и управляемый cancel. Для запуска без cancel exact set заранее равен всему входу; для cancel test gates фиксируют IDs с `success` и `cancelled`.

### Критерии приёмки

- [ ] max active никогда не превышает лимит;
- [ ] acquire/release и close ownership сначала предложены учеником и доказаны;
- [ ] partial-results policy, cutoff и четыре допустимые outcome-категории соблюдены без потери accepted IDs.

### Усложнение после первой версии

Только после рабочей первой версии наставник задаёт разные worker/semaphore limits; ученик защищает, зачем нужны оба.

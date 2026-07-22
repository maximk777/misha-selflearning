# Worker pool

Pool ограничивает число worker и получает jobs из канала. Producer закрывает jobs; owner ждёт `WaitGroup` и закрывает results. При отмене producer и workers прекращают ожидание через context. Размер pool — осознанное ограничение ресурсов, не «побольше для скорости».

## Где это применяется в реальном backend

1. **Обработка фонового batch** — фиксированные workers ограничивают CPU/зависимость; goroutine на каждую job не даёт backpressure.
2. **Повторная обработка файлов** — очередь отделяет producer от workers; unbounded input всё равно требует внешнего лимита.
3. **Будущий HTTP dispatcher** — pool сможет вызывать HTTP client после модуля HTTP; сейчас job — инъецируемая функция без сети.

## Глубокое погружение

У pool должны быть однозначные owners каналов, условие прекращения producer/workers/collector и момент, после которого новые jobs не считаются принятыми. Конкретных owners и close order ученик сначала предлагает сам; наставник проверяет модель и только после первой попытки раскрывает по одному нарушенному инварианту. Buffer задаёт in-process backpressure, но не durable queue. Context участвует во всех potentially blocking send/receive. Error policy (fail-fast или partial) выбирается контрактом; worker count ограничивает одновременно исполняемые jobs, если job не запускает скрытые goroutines. Exact-once в этом проекте означает: каждая job, которую runner подтвердил принятой до cutoff отмены, имеет не более одного запуска и ровно один документированный outcome — result, error или `cancelled`; jobs после cutoff отклоняются. Costs — stacks, channel contention, queue memory. Докажи injected job с atomic active/max-active, `-race`, cancellation и accepted-ID ledger.

## Мини-проект

### Результат

Спроектируй background job runner с ограниченной конкурентностью и инъецируемой `func(context.Context, Job) (Result, error)`. До кода нарисуй lifecycle и сам предложи owners каналов, close order, cutoff принятия и политику результата при cancellation; готовая схема заранее не выдаётся.

### Разрешённые знания

Все темы concurrency до worker pool и Go Core; HTTP запрещён.

### Проверка

После принятия модели, из корня: `cd project/concurrency-runner && go test -race -count=50 ./...` на max concurrency, empty batch, error/cancel и outcome каждой accepted-before-cutoff job.

### Критерии приёмки

- [ ] channel ownership и порядок close сначала предложены учеником, затем доказаны тестом;
- [ ] worker limit доказан, а не выведен из Sleep;
- [ ] после cancel завершены все goroutines, которыми владеет runner; accepted-before-cutoff jobs имеют документированный outcome.

### Усложнение после первой версии

Только после рабочей первой версии наставник выдаёт ограничение bounded input buffer; ученик сам определяет, где возникает backpressure.

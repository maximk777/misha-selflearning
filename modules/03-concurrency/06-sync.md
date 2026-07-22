# sync

`Mutex` защищает инвариант, `RWMutex` — только при оправданных read-heavy нагрузках, `WaitGroup` ждёт известное число работ, `Once` выполняет действие раз. Lock ordering должен быть одинаковым во всех путях. `defer mu.Unlock()` уменьшает риск забыть unlock, но не делает lock-free код безопасным.

## Где это применяется в реальном backend

1. **Счётчики runner** — Mutex защищает согласованный набор полей; отдельные locks могут показать невозможное состояние.
2. **Ожидание workers** — WaitGroup связывает Add/Done/Wait; локальный инвариант runner: owner делает все положительные `Add` до запуска workers и до своего `Wait`.
3. **Ленивая инициализация** — Once выполняет setup один раз; он не предоставляет retry после panic/error без дополнительного дизайна.

## Глубокое погружение

Mutex устанавливает happens-before между Unlock и последующим Lock. Он не reentrant, копировать использованный mutex нельзя. WaitGroup — счётчик lifecycle. Для этого runner выбираем простой локальный контракт: owner вызывает положительный `Add` до запуска соответствующей goroutine и до `Wait`, worker делает `Done`. Это не объявляется универсальным запретом любого `Add` после начала любого `Wait`: точные правила зависят от состояния счётчика и повторного использования. RWMutex добавляет bookkeeping и выигрывает не всегда; benchmark решает. Инвариант формулируется до выбора lock, lock scope минимален, порядок нескольких locks един. Failures: double Unlock, negative WaitGroup, lock across slow job, self-deadlock. Докажи `-race`, timeout-guarded deadlock test и benchmark Mutex/RWMutex при representative workload.

## Мини-проект

### Результат

Добавь runner потокобезопасную статистику accepted/running/succeeded/failed с согласованным snapshot.

### Разрешённые знания

Предыдущие темы, `sync.Mutex`, `WaitGroup`, `Once`; без atomic.

### Проверка

Из корня: `cd project/concurrency-runner && go test -race -count=50 ./...`; проверка invariant `accepted = running + succeeded + failed` в выбранной точке lifecycle.

### Критерии приёмки

- [ ] один lock защищает явно названный invariant;
- [ ] job не вызывается под лишним lock;
- [ ] Add/Done/Wait ownership объяснён.

### Усложнение после первой версии

Сравнить Mutex и RWMutex benchmark и оставить более простой вариант без доказанной выгоды.

# Load balancing

Round-robin — default; weight задаёт долю; least_conn полезен при разной длительности запросов; sticky привязывает session, но ухудшает перераспределение. Instance должен быть stateless, session — внешней.

## Где это применяется в реальном backend

1. **Два stateless API replicas** — round-robin распределяет однородные короткие запросы. Он считает выдачи, а не реальную загрузку, поэтому медленный instance может накопить очередь.
2. **Долгие report requests** — least connections лучше отражает занятость. Длительность connection не равна CPU cost, а один hung connection требует health/timeout отдельно.
3. **Canary новой версии** — weighted routing отправляет малую долю traffic в новую replica. Weight не гарантирует точную долю на малой выборке и без одинаковых contracts превращает rollout в случайные ошибки.

## Глубокое погружение

L4 балансирует connections без HTTP semantics; L7 видит method/path/headers и может маршрутизировать/ретраить осознаннее. Round-robin, weighted, least_conn и consistent hash оптимизируют разные сигналы. Инвариант stateless replicas: любой healthy instance способен обработать запрос; sticky session ослабляет его и усложняет failover. Costs — proxy hop, connection pools, health probes и uneven load. Edge cases: keepalive закрепляет много requests в нескольких connections, slow replica остаётся technically healthy, long-lived streams, retry amplification. Доказывай распределение по instance ID на достаточно большой выборке, latency histogram и остановку replica во время нагрузки.

## Мини-проект

### Результат

Бизнес-сценарий: Task API должен распределять short и long requests между replicas и переживать потерю одной из них. В cumulative `project/backend-lab/compose.yaml` запусти две API replicas за Nginx и сравни round-robin с least_conn на коротких и искусственно медленных запросах. Вывод должен содержать доли requests, p95 и поведение при остановке одной replica; TLS и Kubernetes не добавляй.

### Разрешённые знания

Docker → Compose → Nginx из предыдущих тем, HTTP/concurrency/metrics и текущие L4/L7/алгоритмы balancing. TLS и Kubernetes не нужны.

### Проверка

Из корня репозитория выполни `docker compose -f project/backend-lab/compose.yaml config`, `docker compose -f project/backend-lab/compose.yaml up -d --build --scale api=2` и `docker compose -f project/backend-lab/compose.yaml exec nginx nginx -t`; генератор выполняет минимум 100 запросов и группирует `instance_id`, errors и latency. Во время серии останови одну API replica и измерь failed/recovered requests.

### Критерии приёмки

- [ ] оба instances обслуживают запросы без local-session correctness;
- [ ] выбранный алгоритм объяснён измерениями для короткой и долгой нагрузки;
- [ ] failure experiment фиксирует ошибки/время восстановления, а не только итоговое `Up`;
- [ ] retry unsafe POST отсутствует или защищён ранее изученной idempotency.

### Усложнение после первой версии

Сделай weighted canary 90/10 и определи размер выборки, после которого измеренная доля достаточно убедительна.

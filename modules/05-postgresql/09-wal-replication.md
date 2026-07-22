# WAL и физическая репликация

Сначала изменение пишется в WAL, затем страницы данных. Физическая реплика воспроизводит WAL-блоки всего кластера. Синхронная репликация ждёт подтверждение выбранной standby и уменьшает риск потери; асинхронная отвечает быстрее, но replica может отставать.

## Где это применяется в реальном backend

1. **Crash recovery order store** — WAL replay восстанавливает committed changes; WAL не является backup от логического DELETE.
2. **Read replica** — разгружает reports; replication lag даёт stale order status.
3. **Durability/latency** — synchronous standby уменьшает loss window, но добавляет commit latency.

## Глубокое погружение

WAL records предшествуют dirty page flush; LSN задаёт позицию. Checkpoint ограничивает recovery path, replication slots удерживают WAL. Physical replica повторяет cluster-level bytes и обычно read-only. Failures: lag, slot disk growth, failover data loss. Доказывай LSN/lag metrics и controlled write/read observation.

## Мини-проект

### Результат

Проведи order-store эксперимент: записать order на primary, наблюдать его появление на replica и измерить lag/LSN, не строя failover automation.

### Разрешённые знания

Все предыдущие PostgreSQL темы; Kafka/CDC не требуются.

### Проверка

SQL для current/replay LSN и timestamped write/read; зафиксировать поведение при остановленной replica, если lab позволяет.

### Критерии приёмки

- [ ] commit и replica visibility различаются измеримо;
- [ ] WAL retention/slot risk объяснён;
- [ ] реплика не объявлена backup или strong read.

### Усложнение после первой версии

После MVP определить read-after-write policy для одного endpoint и проверить её.

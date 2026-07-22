# Kubernetes delivery

Readiness исключает Pod из traffic, liveness перезапускает зависший process, startup защищает медленный старт. HPA масштабирует по наблюдаемому сигналу, не по желанию. Rollout проверяй, rollback планируй до deploy.

## Где это применяется в реальном backend

1. **Rolling update без отправки traffic в старующий Pod** — readiness открывает endpoint только после готовности. Liveness, проверяющий внешнюю DB, может перезапустить весь fleet во время общей аварии.
2. **Защита медленного startup** — startup probe даёт время migration/warmup прежде, чем liveness начнёт рестарты. Слишком большой budget скрывает настоящий hang и удлиняет recovery.
3. **HPA по saturation** — replicas растут по CPU или бизнес-метрике backlog. Без requests CPU utilization бессмысленна, а scaling не лечит bottleneck в DB и может его усилить.

## Глубокое погружение

Deployment rollout создаёт новый ReplicaSet и соблюдает `maxSurge/maxUnavailable`; readiness участвует в availability, liveness — в restart, startup временно выключает liveness/readiness cadence. Termination требует endpoint removal, SIGTERM, grace period и прекращения intake; race между ними даёт dropped requests. Scheduler использует requests, runtime limits влияют на throttling/OOM; HPA реагирует с задержкой и может oscillate. Costs — duplicate capacity during rollout, slow metric loops и probe traffic. Edge cases: bad image, readiness flapping, PDB blocking drain, schema incompatibility old/new versions, HPA scaling consumers beyond DB pool. Без cluster поведение доказывается только schema/render и reasoned state simulation; rollout status/history, probe fault injection, in-flight termination и реальный rollback становятся доказательством лишь в optional cluster run.

## Мини-проект

### Результат

Бизнес-сценарий: Task API нужно обновлять без потери доступности, даже если у ученика сейчас нет Kubernetes cluster. Расширь `project/backend-lab/deploy/kubernetes/`: добавь probes с разной semantics, requests/limits, rolling strategy и осознанный autoscaling signal. Обязательный результат — static manifests, schema validation/render и reasoned simulation v1→v2→bad release→rollback с расчётом replicas/surge/unavailable и условием отката. Реальный rollout не требуй без cluster; он остаётся отдельным optional proof.

### Разрешённые знания

Все предыдущие темы, включая Kubernetes core, graceful shutdown, metrics, load balancing и data-store constraints. Новые service mesh/GitOps/managed-cloud механизмы не требуются.

### Проверка

Обязательный no-cluster path из корня: `kubectl apply --dry-run=client -o yaml -f project/backend-lab/deploy/kubernetes/`; команда валидирует локально известную schema и печатает inspectable render. Затем таблица симуляции показывает для каждого шага desired/old/new/ready/unavailable replicas, соблюдение `maxSurge/maxUnavailable`, failed readiness и выбранную rollback revision. Для HPA рассчитай исходную метрику, replicas и предел downstream pool. Optional real-cluster proof выполняется отдельно: server-side dry run, `kubectl rollout status/history`, request loop, probe failure injection, events и реальный rollback.

### Критерии приёмки

- [ ] readiness/liveness/startup проверяют разные свойства и имеют обоснованные budgets;
- [ ] no-cluster acceptance проходит local schema validation/render и reasoned rollout/rollback simulation без runtime assertions;
- [ ] simulation bad release останавливается и выбирает rollback по заранее записанному сигналу;
- [ ] утверждения про реальные in-flight requests, probes и rollback предъявляются только при optional cluster proof;
- [ ] autoscaling ограничен downstream capacity, а защита trade-offs включает surge, shutdown и schema compatibility.

### Усложнение после первой версии

Добавь PodDisruptionBudget и смоделируй maintenance одного node: объясни, когда PDB защищает availability, а когда блокирует операцию.

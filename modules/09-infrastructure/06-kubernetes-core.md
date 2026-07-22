# Kubernetes core

Pod запускает containers; Deployment управляет ReplicaSet/rolling update; Service даёт стабильный DNS; ConfigMap — non-secret config, Secret — чувствительные значения. Requests влияют на scheduling, limits сдерживают потребление.

## Где это применяется в реальном backend

1. **Поддержание трёх API replicas** — Deployment reconciliation восстанавливает desired count и управляет template revision. Он не гарантирует готовность нового Pod без probes следующей темы.
2. **Стабильный адрес эфемерных Pods** — Service выбирает Pods по labels и даёт DNS/VIP. Ошибка selector создаёт Service без endpoints при полностью Running Pods.
3. **Разделение config и image** — ConfigMap хранит non-secret settings, Secret — чувствительные bytes. Secret не становится безопасным автоматически: доступ RBAC, encryption at rest и исключение из logs остаются обязанностями.

## Глубокое погружение

Kubernetes — control loop: desired objects лежат в API server, controllers сравнивают их с observed state. Pod — минимальная scheduling unit и обычно disposable; Deployment владеет ReplicaSet, Service находит endpoints по labels, Ingress требует controller. Requests используются scheduler и QoS, limits enforced runtime и могут дать CPU throttling/OOMKilled. Config change не обязательно перезапускает процесс. Costs — cluster control plane, YAML/API evolution, image pulls и capacity reservations. Edge cases: wrong namespace/selector, mutable tag, Pending from insufficient requests, CrashLoopBackOff, mounted Secret update semantics. Доказывай server-side dry run, `describe`, events, endpoints и deletion/self-healing experiment.

## Мини-проект

### Результат

Бизнес-сценарий: platform team должна получить переносимые manifests Task API до предоставления реального cluster. Создай отсутствующий path `project/backend-lab/deploy/kubernetes/` и опиши `backend-lab` как Deployment, Service, ConfigMap и Secret reference. Обязательная версия ограничена static manifests, локальным render/schema validation и reasoned reconciliation table; реальный cluster proof выполняется отдельно только при доступном окружении.

### Разрешённые знания

Предыдущая infrastructure-цепочка Docker/Compose/Nginx/LB/TLS/timeouts/retries, observability/security и текущие Kubernetes core objects. Probes, rolling strategy и autoscaling следующей темы не обязательны.

### Проверка

Обязательный no-cluster path из корня: `kubectl apply --dry-run=client -o yaml -f project/backend-lab/deploy/kubernetes/` валидирует локально известную schema и печатает inspectable render; отдельно проверь selectors/ports/image/config по таблице. Optional real-cluster proof: `kubectl apply --dry-run=server -f project/backend-lab/deploy/kubernetes/`, `kubectl diff -f project/backend-lab/deploy/kubernetes/`, `kubectl get pods,deploy,svc,endpoints`, `kubectl describe`, DNS/request и удаление Pod с наблюдением replacement.

### Критерии приёмки

- [ ] static labels/selectors и ports согласованы; реальные Service endpoints требуются только в optional cluster proof;
- [ ] image pin, ports, config и secret references согласованы с Docker/Compose версией;
- [ ] requests/limits заданы и trade-off throttling/OOM объяснён;
- [ ] no-cluster acceptance содержит manifests, schema validation/render и reasoned reconciliation; runtime claims появляются только с optional cluster evidence.

### Усложнение после первой версии

Создай отдельный namespace и минимальный ServiceAccount/RBAC для чтения только нужного ConfigMap, затем докажи запрещённый доступ к Secret.

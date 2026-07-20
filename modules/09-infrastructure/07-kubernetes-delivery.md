# Kubernetes delivery

Readiness исключает Pod из traffic, liveness перезапускает зависший process, startup защищает медленный старт. HPA масштабирует по наблюдаемому сигналу, не по желанию. Rollout проверяй, rollback планируй до deploy.

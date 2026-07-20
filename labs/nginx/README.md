# Nginx balancing

Compose должен дать несколько HTTP replicas и Nginx с конфигом `deploy/nginx/nginx.conf`. Запусти инфраструктуру через родительские scripts, затем посылай повторные запросы и наблюдай label replica. Останавливай `go run` через `Ctrl-C`.

# Balancing

Проверь config: `docker compose -f deploy/compose.yaml exec nginx nginx -t`. Затем `for i in $(seq 1 8); do curl -s http://localhost/; done` и наблюдай replica label. Поломка: временно убери один upstream и наблюдай `502`; верни config. Сравни round-robin, `weight`, `least_conn` и sticky: sticky не даёт free failover.

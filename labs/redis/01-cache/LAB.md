# Cache-aside

Запусти `cd starter && go test ./...`. CLI: `redis-cli SET task:1 value EX 2`, `redis-cli TTL task:1`, затем после паузы `GET task:1`. Поломка: временно задай TTL `0` и предскажи постоянные miss; верни TTL. Перескажи stale cache и почему mutex локального процесса не решает distributed stampede.

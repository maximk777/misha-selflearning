# net/http

`Handler` получает `ResponseWriter` и `*Request`; router может быть `http.ServeMux`. Сервер обязан иметь read/write/idle timeouts. Handler сначала валидирует method/path/input, затем пишет headers/status/body. Один handler не должен тайно владеть запуском process.

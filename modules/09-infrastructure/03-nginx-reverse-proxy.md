# Nginx reverse proxy

Nginx принимает client connection и проксирует upstream, выставляя `X-Request-ID`, `X-Forwarded-*`. Timeout защищает занятые resources; retry небезопасен для non-idempotent POST без ключа. Config проверяй `nginx -t` внутри container.

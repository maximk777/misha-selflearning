# Middleware и JSON

Middleware оборачивает `next`: создаёт/принимает request ID, кладёт его в header/context и вызывает `next`. JSON decoder не должен принимать неизвестные поля без решения команды; validation возвращает `400` с `{\"error\": ...}`. Нельзя писать два разных ответа после ошибки.

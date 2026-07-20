# Client

Клиент имеет timeout или request context. После успешного `Do` всегда `defer resp.Body.Close()`, даже при не-2xx статусе. Сначала проверь status, затем ограниченно декодируй JSON. Timeout защищает ожидание; cancellation сообщает, что результат больше не нужен.

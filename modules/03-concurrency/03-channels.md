# Channels

Unbuffered channel синхронизирует sender и receiver; buffered разрешает до capacity отправок без receiver. Закрывает только владелец отправки, ровно один раз. Receive после close возвращает zero value и `ok=false`; send after close паникует. Выполни `go test ./labs/concurrency/...` после предсказания блокировки.

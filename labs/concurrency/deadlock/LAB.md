# Deadlock

Сначала запусти `go run ./deadlock/starter`. Затем прогноз: почему два lock в разном порядке опасны? Выполни `go run ./deadlock/starter -deadlock`; Go завершит процесс с `fatal error: all goroutines are asleep - deadlock!`, поэтому cleanup не нужен. Почини упражнение самостоятельно единым порядком lock и перескажи lock ordering.

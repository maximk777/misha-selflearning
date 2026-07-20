# Worker pool

Предскажи, что ограничивает число worker. Выполни `go test ./workerpool/starter`. Поломка: временно убери `case <-ctx.Done()` у producer и запусти `go test -timeout 1s ./workerpool/starter`; ожидай timeout при cancellation. Восстанови cancel path, проверь тест и перескажи, кто закрывает jobs и кто ждёт workers.

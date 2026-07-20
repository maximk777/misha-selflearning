# Pipeline

Запусти `go test ./pipeline/starter` и предскажи закрытие output. Поломка: временно убери `defer close(out)` и выполни `go test -timeout 1s ./pipeline/starter`; ожидай timeout range. Восстанови close, затем перескажи владельца каждого output и реакцию stage на cancellation.

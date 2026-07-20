# Context и leak

Прогноз: что вернёт worker после `cancel`? Запусти `go test ./context/starter`. Поломка: временно замени `<-ctx.Done()` на `select {}` и выполни `go test -timeout 1s ./context/starter`; ожидай timeout. Верни cancellation-путь, запусти тест и перескажи, почему context предотвращает leak.

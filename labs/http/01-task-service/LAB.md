# JSON task service

Из `starter` предскажи статус и JSON: `go test ./...`, затем `go test -race ./...`. Проверь реальный handler через тесты `httptest`, а не постоянный server process. Bounded failure: временно удали `defer resp.Body.Close()` в client и объясни leak; верни строку до финального теста. Второй: временно убери `DisallowUnknownFields`, добавь неизвестное поле в тест и наблюдай ошибочно принятый запрос; восстанови validation. Перескажи: timeout client/server, request context, error envelope, request ID и `Shutdown`.

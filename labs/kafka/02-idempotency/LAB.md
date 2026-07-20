# Idempotency и DLQ

Запусти `cd starter && go test ./...`. Предскажи duplicate delivery: тот же key дважды должен примениться раз. Поломка: временно убери map check и наблюдай падение теста. Для poison record опиши bounded retry/backoff и отправку исходного payload+ошибки в DLQ до commit.

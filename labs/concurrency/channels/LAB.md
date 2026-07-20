# Channels

Прогноз: unbuffered send ждёт receiver, buffered — лишь до capacity. Запусти `go test ./channels/starter`. Поломка: временно замени `make(chan int, len(values))` на `make(chan int)` и запусти `go test -timeout 1s ./channels/starter`; ожидай timeout, затем верни код. Добавь временный `SendOnce(... ) <- 1` после `close` и предскажи panic `send on closed channel`; немедленно убери строку. Почини сам и перескажи, почему закрывает producer.

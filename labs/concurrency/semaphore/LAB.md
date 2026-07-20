# Semaphore

Предскажи максимум одновременных jobs. Запусти `go test ./semaphore/starter`. Поломка: временно перенеси `<-sem` перед `job()` и запусти тест; ожидай `max > 2`. Верни release в `defer`, проверь тест и перескажи acquire/release.

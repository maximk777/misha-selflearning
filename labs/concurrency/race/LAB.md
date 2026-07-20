# Race detector

Запусти безопасный путь: `go test -race ./race/starter`. Затем предскажи результат и выполни bounded demo: `RACE_DEMO=1 go test -race -run TestDeliberateRace ./race/starter`; ожидай `WARNING: DATA RACE` и non-zero exit. Не исправляй detector: объясни, почему mutex защищает оба доступа, затем вернись к обычной команде.

# httptest

`httptest.NewRequest` + `httptest.NewRecorder` проверяют handler без сети. `httptest.NewServer` проверяет реальный client path и автоматически даёт URL. Проверяй status, headers и JSON body. Закрывай test server через `defer server.Close()`.

## Где это применяется в реальном backend

1. **Handler unit test** — recorder проверяет mapping; real transport он не доказывает.
2. **Client integration test** — test server воспроизводит non-2xx, slow/malformed body без внешней сети.
3. **Contract regression** — field/status checks ловят несовместимость; строковый snapshot хрупок.

## Глубокое погружение

Recorder вызывает handler в памяти, NewServer использует listener/transport. Server и bodies закрываются owner'ом. Инварианты: inputs контролируемы, semantic assertions, нет network/timing зависимости. Надёжность доказывают `-count=10`, `-race` и intentional break.

## Мини-проект

### Результат

Создай тесты Task API/client, ловящие неверный status/header/body, timeout, malformed response и concurrent ошибку.

### Разрешённые знания

Все HTTP темы до `httptest`, testing/table tests и пройденная concurrency.

### Проверка

`go test ./... -count=10`, `go test -race ./...`; намеренно сломать contract, увидеть failure и починить.

### Критерии приёмки

- [ ] Recorder/NewServer выбраны по границе;
- [ ] happy path и три failures покрыты, ресурсы закрыты;
- [ ] тесты детерминированы, ограничения объяснены.

### Усложнение после первой версии

Проверить streaming/flush либо disconnect без fixed sleep.

# Server

`http.Server` отделяет конфигурацию от handler. JSON service обязан задать `Content-Type`, использовать понятные статусы и не возвращать внутренние ошибки клиенту. Ошибки кодируй единым envelope. Проверяй server через `httptest`, не через реальный порт в unit-test.

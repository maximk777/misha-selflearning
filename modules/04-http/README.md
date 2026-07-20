# HTTP

Dependency-free путь: `net/http` server → JSON handler → middleware → client/timeouts → `httptest` → shutdown. На каждом шаге: прогноз → запуск → ограниченная поломка → починка → пересказ.

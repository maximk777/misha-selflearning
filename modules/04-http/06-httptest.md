# httptest

`httptest.NewRequest` + `httptest.NewRecorder` проверяют handler без сети. `httptest.NewServer` проверяет реальный client path и автоматически даёт URL. Проверяй status, headers и JSON body. Закрывай test server через `defer server.Close()`.

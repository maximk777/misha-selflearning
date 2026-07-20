# Банк: HTTP

Использовать после HTTP-лаборатории. Минимум один вопрос должен сопровождаться `httptest` или живым запросом. Оценивание: [RUBRIC.md](RUBRIC.md).

| ID и теги | Тип | Вопрос экзаменатора | Follow-up и практическое доказательство |
|---|---|---|---|
| HT-1 · `R1 R7` | Концептуальный | Какие timeout нужны HTTP client и почему один бесконечный default опасен? | Назови бюджет timeout для запроса и cancellation. |
| HT-2 · `R1` | Чтение/debugging | В клиенте после `Do` не вызывают `resp.Body.Close()`. Какая неисправность вероятна? | Исправь код и объясни влияние на соединения. |
| HT-3 · `R7` | Практический | Напиши handler: decode JSON, validate, вернуть единый JSON error и request ID. | Покажи `httptest` на invalid JSON и validation error. |
| HT-4 · `R21` | Misconception probe | «HTTP 500 достаточно, клиент сам поймёт ошибку». Что теряет API? | Предложи минимальный error envelope без утечки internal details. |
| HT-5 · `R1 R7` | Чтение/debugging | Request context отменён: где это должно быть проверено в handler/service? | Добавь blocking dependency и тест cancellation. |
| HT-6 · `surprise-old` | Production | Как сделать graceful shutdown сервера, не бросив активные запросы? | Покажи порядок: stop intake, timeout, drain/close и лог. |

<!-- ТОЛЬКО ЭКЗАМЕНАТОРУ
HT-1: client timeout/deadline prevents stuck connections; budgets include upstream. Avoid universal numeric answer.
HT-2: body close enables reuse/resources; also handle error response semantics.
HT-3: decode limits optional; reject malformed input, stable error schema, test status/body.
HT-4: stable code/message/request id; no raw DB/internal errors.
HT-5: propagate r.Context, dependencies honour it; do not create background context.
HT-6: Server.Shutdown context, stop accepting, bounded wait; distinguish readiness/liveness if relevant.
-->

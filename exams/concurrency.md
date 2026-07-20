# Банк: Конкурентность

Использовать после каждой 3–4 тем и перед HTTP. Критичные нули: ownership закрытия channel, cancellation и race. Оценивание: [RUBRIC.md](RUBRIC.md).

| ID и теги | Тип | Вопрос экзаменатора | Follow-up и практическое доказательство |
|---|---|---|---|
| CO-1 · `R1 R7` | Концептуальный | Объясни P, M и G так, чтобы не спутать goroutine с OS thread. | Что изменится, если goroutine блокируется в syscall? |
| CO-2 · `R1` | Чтение/debugging | Кто должен закрывать channel и почему receiver обычно не делает этого? | Исправь send-on-closed-channel в минимальном примере. |
| CO-3 · `R7 R21` | Практический | Сделай worker pool, который прекращает принимать работу по `context` и дожидается начатой работы с bounded timeout. | Покажи тест/лог shutdown и отсутствие зависания. |
| CO-4 · `R1 R7` | Misconception probe | «Mutex всегда медленнее channel, поэтому всё надо делать channel». Что неверно? | Выбери mutex или channel для защиты счётчика и обоснуй. |
| CO-5 · `R21` | Чтение/debugging | `go test -race` сообщает две строки доступа к map. Что это доказывает, а чего не доказывает? | Почини с mutex или ownership и снова запусти `-race`. |
| CO-6 · `surprise-old` | Практический | Две goroutine берут locks в разном порядке. Как диагностировать и предотвратить deadlock? | Воспроизведи с timeout, затем введи общий порядок locks. |

<!-- ТОЛЬКО ЭКЗАМЕНАТОРУ
CO-1: G goroutine, M OS thread, P resource/context scheduler; blocked syscall may require another M/P handling, avoid demanding runtime internals beyond model.
CO-2: sender/producer owns close; close signals no more sends; coordinate many senders.
CO-3: expect context propagation, close/stop intake, WaitGroup and bounded shutdown; discuss policy for in-flight jobs.
CO-4: choose by ownership/communication vs simple shared state; performance first measure.
CO-5: race detector finds exercised races only, not proof absence; synchronize or single-owner.
CO-6: global lock order, short critical sections, retry only after cause considered.
-->

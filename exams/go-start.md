# Банк: Go Start

Использовать после тем Go Start. Возьми 3–5 вопросов, среди них один reading/debugging и один практический. Оценивание: [RUBRIC.md](RUBRIC.md). Теги повторения сохраняются при записи результата.

| ID и теги | Тип | Вопрос экзаменатора | Follow-up и практическое доказательство |
|---|---|---|---|
| GS-1 · `R1 R7` | Концептуальный | Чем package отличается от module и что делает `go test ./...`? | Покажи команду из своего модуля и объясни, какой package тестируется. |
| GS-2 · `R1 R21` | Чтение кода | Предскажи вывод и aliasing: `a := []int{1,2}; b := a[:1]; b = append(b, 9); fmt.Println(a, b)`. От чего зависит ответ? | Измени capacity или используй full slice expression и докажи тестом. |
| GS-3 · `R7` | Misconception probe | «Slice — это динамический массив, поэтому передача slice всегда копирует все элементы». Согласен? | Назови состав slice header и собери минимальный пример общей backing array. |
| GS-4 · `R1 R7` | Чтение/debugging | Почему `var m map[string]int; m["x"] = 1` паникует, а чтение `m["x"]` допустимо? | Исправь код и добавь проверку отсутствующего ключа через comma-ok. |
| GS-5 · `R21` | Практический | Напиши `ParsePort(string) (int, error)`, который не паникует и возвращает контекст ошибки. | Запусти table-driven test с пустой строкой и нечисловым значением. |
| GS-6 · `surprise-old` | Концептуальный | Когда выбрать pointer receiver, а когда value receiver? | Покажи, как копирование struct меняет наблюдаемое поведение. |

<!-- ТОЛЬКО ЭКЗАМЕНАТОРУ
GS-1: module — единица версионирования и go.mod; package — компилируемая единица. go test ./... обходит packages под текущим module.
GS-2: append может писать в общую backing array при достаточной capacity; ответ зависит от cap. Full slice expression ограничивает cap.
GS-3: slice header содержит pointer/len/cap; элементы не копируются автоматически.
GS-4: nil map доступна на чтение, но запись требует allocate через make или literal.
GS-5: ожидаем strconv.Atoi плюс fmt.Errorf с %w при уместном sentinel; доказательство — тест.
GS-6: pointer нужен для mutation/избежания крупного copy/методного набора; value — small immutable semantics. Не требовать «всегда pointer».
-->

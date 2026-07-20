# Банк: Go Core

Использовать после Go Core и перед первым mock. Взять 3–5 вопросов; минимум один вопрос требует чтения кода. Оценивание: [RUBRIC.md](RUBRIC.md).

| ID и теги | Тип | Вопрос экзаменатора | Follow-up и практическое доказательство |
|---|---|---|---|
| GC-1 · `R1 R7` | Концептуальный | Почему интерфейс в Go обычно описывает поведение со стороны потребителя? | Перепроектируй зависимость от конкретного storage в маленький interface. |
| GC-2 · `R7` | Чтение/debugging | Что выведет `var p *T = nil; var x I = p; fmt.Println(x == nil)` и почему? | Покажи безопасную проверку и объясни dynamic type/value. |
| GC-3 · `R1 R21` | Практический | Сделай error API, где caller отличает not found от другой ошибки через `errors.Is` или `errors.As`. | Добавь тест на wrapped error и объясни, почему `%v` недостаточно. |
| GC-4 · `R7` | Misconception probe | «Generics всегда заменяют interfaces и делают код быстрее». Что неверно? | Приведи по одному случаю для generic и interface. |
| GC-5 · `R1` | Чтение/debugging | В table-driven test почему нельзя бездумно захватывать переменную цикла в parallel subtest? | Напиши минимальный test с отдельной переменной/современной семантикой и докажи запуском. |
| GC-6 · `R21 surprise-old` | Production | Benchmark показывает allocation. Что проверишь до оптимизации? | Покажи baseline и назови, почему microbenchmark не заменяет profile. |

<!-- ТОЛЬКО ЭКЗАМЕНАТОРУ
GC-1: малый interface располагается у consumer; concrete implementation не навязывается заранее.
GC-2: x != nil: interface имеет dynamic type *T и nil dynamic value. Нужны explicit design/typed nil handling.
GC-3: %w сохраняет unwrap chain; проверить errors.Is/As тестом.
GC-4: generic — параметризация алгоритма/type constraints; interface — runtime polymorphism/behavior boundary. Нет универсального faster.
GC-5: оценить понимание capture и фактической версии Go, не ловушку на синтаксис.
GC-6: baseline, representative load, pprof/allocs and semantic cost.
-->

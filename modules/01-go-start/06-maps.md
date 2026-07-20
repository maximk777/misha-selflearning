# Maps и zero value

> Уровень: `core`
> Время: 60 минут
> Предпосылки: variables, slices
> Практика: `labs/go/04-maps/LAB.md`
> Экзамен: `exams/go-start.md`, GS-4

## Результат за 5–15 минут

Запусти `cd labs/go/04-maps/starter && go test ./...`. Тест отличает существующий ключ со значением `0` от отсутствующего ключа.

## Модель

Чтение из nil map безопасно и возвращает zero value типа значения. Запись в nil map паникует: для записи нужна `make(map[K]V)` или literal. Форма `value, ok := m[key]` отделяет «нет ключа» от «есть zero value».

## Прогноз → опыт → поломка

**Один вопрос:** что напечатает чтение `var scores map[string]int; fmt.Println(scores["Миша"])`? Проверь в коротком временном коде. Затем измени `NewScores` на `return nil`, получи panic в test и почини.

## Пересказ и сдача

**Один вопрос:** когда одного `scores[name]` недостаточно? Покажи test и объясни comma-ok.

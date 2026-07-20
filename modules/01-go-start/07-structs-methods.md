# Structs и методы

> Уровень: `core`
> Время: 60 минут
> Предпосылки: functions
> Практика: `labs/go/05-structs-methods/LAB.md`
> Экзамен: `exams/go-start.md`, GS-6

## Результат за 5–15 минут

Выполни test `labs/go/05-structs-methods/starter`. `Deposit` меняет balance, а `Label` читает account и формирует строку.

## Модель

`struct` объединяет данные в именованный тип. Метод с receiver делает операцию частью поведения типа. Pointer receiver выбирают, когда метод должен изменить исходное значение или когда копирование дорого; value receiver — для маленького значения без изменения его состояния.

## Прогноз → опыт → поломка

**Один вопрос:** изменит ли balance метод `func (a Account) Deposit(...)`? Поменяй receiver на value, запусти test, прочитай failure и верни pointer receiver.

## Пересказ и сдача

**Один вопрос:** почему `Label` может иметь value receiver, а `Deposit` — нет? Покажи test и ответь на GS-6.

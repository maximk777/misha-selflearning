# Лабораторная: методы и receiver

> Тема: `modules/01-go-start/07-structs-methods.md`
> Уровень: `core`
> Время: 25 минут
> Запуск из: `labs/go/05-structs-methods/starter/`

## Результат за 5–15 минут

```bash
cd labs/go/05-structs-methods/starter
go test ./...
```

`Deposit` меняет balance до 15; `Label` возвращает `Миша: 10`.

## До запуска: прогноз

**Один вопрос:** почему `Deposit` использует `*Account`, а `Label` — `Account`?

## Поломка и самостоятельная починка

Замени `func (account *Account) Deposit` на value receiver, запусти test и верни pointer receiver после объяснения failure.

## Пересказ

**Один вопрос:** когда value receiver безопаснее и проще? Сдача — зелёный test, поломка и аргумент выбора receiver.

# Лабораторная: интерфейс и typed nil

> Тема: `modules/02-go-core/01-interfaces.md`
> Уровень: `core`
> Время: 30 минут
> Запуск из: `labs/go/07-interfaces/starter/`

## Результат за 5–15 минут

```bash
cd labs/go/07-interfaces/starter
go test ./...
```

Ожидаются два факта: `Notify` зависит от поведения, а interface с dynamic type `*emailNotifier` и nil value не равен `nil`.

## До запуска: прогноз

**Один вопрос:** будет ли `typedNil == nil`, если в interface положить `(*emailNotifier)(nil)`?

## Поломка и самостоятельная починка

Временно измени `IsNilInterface` на `return true`, посмотри падение test, затем верни честное сравнение с `nil`. Не вызывай `Notify` на typed nil: это отдельный риск, который нужно явно проектировать.

## Пересказ

**Один вопрос:** из каких двух частей логически состоит interface value? Сдача — два tests и объяснение typed nil.

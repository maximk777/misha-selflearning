# Лабораторная: wrapping errors и panic

> Тема: `modules/01-go-start/09-defer-panic-recover.md`
> Уровень: `interview`
> Время: 30 минут
> Запуск из: `labs/go/06-errors-panic/starter/`

## Результат за 5–15 минут

```bash
cd labs/go/06-errors-panic/starter
go test ./...
```

Ожидаются tests: `errors.Is` находит `ErrNotFound`; `MustPositive` паникует только для `0`.

## До запуска: прогноз

**Один вопрос:** найдёт ли `errors.Is` исходную ошибку после `fmt.Errorf("...: %w", err)`?

## Поломка и самостоятельная починка

Замени `%w` на `%v`, запусти test и объясни, почему текст похож, а цепочка причины пропала. Верни `%w`. Не превращай обычную ошибку поиска в panic.

## Пересказ

**Один вопрос:** где `panic` оправдан, а где вернуть `error`? Сдача — зелёный test, failure и собственный пример.

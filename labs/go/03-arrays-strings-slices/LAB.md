# Лабораторная: UTF-8 и общий backing array

> Тема: `modules/01-go-start/04-arrays-strings.md`, `05-slices.md`
> Уровень: `core`
> Время: 35 минут
> Запуск из: `labs/go/03-arrays-strings-slices/starter/`

## Результат за 5–15 минут

```bash
cd labs/go/03-arrays-strings-slices/starter
go test ./...
```

Ожидаются green tests: `FirstRune("Жора") == 'Ж'`; `append` с capacity 2 пишет marker в общий массив.

## До запуска: прогноз

**Один вопрос:** `len("Ж")` — это число bytes или runes? А изменится ли `base[:2][1]` после append?

## Поломка и самостоятельная починка

Сделай `FirstRune` через `text[0]`, получи test failure, потом восстанови `range`. Затем сделай capacity base равной 1 и объясни, почему старая проверка общей backing array больше не применима; верни 2.

## Пересказ

**Один вопрос:** какие три поля у slice header? Сдача — оба tests, наблюдение capacity и пересказ.

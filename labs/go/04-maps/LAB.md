# Лабораторная: nil map и comma-ok

> Тема: `modules/01-go-start/06-maps.md`
> Уровень: `core`
> Время: 25 минут
> Запуск из: `labs/go/04-maps/starter/`

## Результат за 5–15 минут

```bash
cd labs/go/04-maps/starter
go test ./...
```

Ожидаются tests на существующий ключ со значением `0`, отсутствующий ключ и пригодную для записи map.

## До запуска: прогноз

**Один вопрос:** почему чтение из nil map работает, а запись нет?

## Поломка и самостоятельная починка

Верни `nil` из `NewScores`, запусти test и прочитай panic `assignment to entry in nil map`. Верни `make(map[string]int)`.

## Пересказ

**Один вопрос:** зачем нужен второй результат `ok`? Сдача — output test, panic signature и объяснение.

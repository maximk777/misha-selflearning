# Producer/consumer

Создай topic через CLI image из Compose, отправь два records с одинаковым key и наблюдай одну partition/упорядоченные offsets. Запусти два consumers одной group и наблюдай assignment/rebalance при остановке одного через `Ctrl-C`. Поломка: commit до обработки — предскажи потерю. Верни manual commit после success.

# Разбор 03: UTF-8 и slices

**Мисконцепция:** string индексируется символами или slice автоматически копирует элементы.

`text[0]` возвращает первый byte UTF-8; `range` декодирует rune. Slice содержит pointer/len/cap; при spare capacity `append` использует backing array, поэтому второй view видит marker. При capacity 1 append может выделить другой массив — нельзя строить API на случайном факте о capacity.

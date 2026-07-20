# Разбор 05: methods

**Мисконцепция:** method receiver всегда меняет исходный struct.

Value receiver получает copy. Поэтому `Deposit` использует `*Account`: mutation остаётся видимой caller. `Label` не меняет state и может работать с copy. Выбор receiver должен быть согласован по типу, особенно когда методы образуют API.

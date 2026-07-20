# Блокировки
Row locks защищают выбранные строки, table locks — операции над отношением, advisory locks — договорённость приложения. `SELECT ... FOR UPDATE` нужен для read-modify-write, но держи транзакцию короткой и бери locks в стабильном порядке.

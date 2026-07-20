# Isolation и MVCC
MVCC хранит версии строк, поэтому читатель обычно не блокирует писателя. `READ COMMITTED` даёт новый snapshot каждому statement; `REPEATABLE READ` — транзакции; `SERIALIZABLE` может отменить конфликтующую транзакцию. Важно уметь назвать dirty/non-repeatable/phantom/read-write anomaly.

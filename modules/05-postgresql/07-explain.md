# EXPLAIN
Читай план изнутри наружу: node type, estimated/actual rows, loops, time, buffers. `EXPLAIN (ANALYZE, BUFFERS)` реально выполняет запрос. Seq Scan нормален для маленькой таблицы или низкой selectivity; наличие индекса не обязывает planner его использовать.

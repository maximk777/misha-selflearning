# Logical replication и CDC
Publication выбирает таблицы/операции, subscription применяет логические изменения. `REPLICA IDENTITY` определяет, как найти старую строку для UPDATE/DELETE. Debezium читает log/WAL через connector, превращает изменения в события и часто публикует их в Kafka; это CDC, а не замена бизнес-outbox во всех случаях.

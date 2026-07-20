# Streams

Streams хранят записи и consumer groups отслеживают pending/ack. `XADD`, `XREADGROUP`, `XACK` дают повторное чтение, но delivery всё равно at-least-once: consumer обязан быть idempotent. Pending без ACK — сигнал retry/claim.

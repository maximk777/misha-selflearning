# Producers и consumers

Producer подтверждает запись по требуемому acks; consumer читает и коммитит offset после безопасной обработки. Manual commit после side effect даёт at-least-once. Commit до обработки даёт at-most-once и возможную потерю.

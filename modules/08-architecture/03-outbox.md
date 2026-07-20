# Transactional outbox

В одной DB transaction запиши business state и outbox record. Publisher читает outbox, публикует, помечает доставленным; crash создаёт duplicate, поэтому consumer idempotent. Outbox исправляет dual-write, не создаёт exactly-once.

# Delivery semantics

At-most-once: commit до обработки. At-least-once: обработка до commit, значит duplicates. Exactly-once — граница конкретной transactional системы, не обещание обычного HTTP/database side effect. Проектируй consumer idempotent.

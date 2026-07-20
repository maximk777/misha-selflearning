# TLS, timeouts, retries

TLS завершается на ingress/proxy с управляемым сертификатом. У каждой границы есть deadline; retry имеет лимит, backoff и только для безопасных операций. Сумма downstream timeout должна укладываться в upstream deadline.

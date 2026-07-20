package postgres

import "context"

// ClaimOutboxSQL documents the transaction-safe worker query for the Compose checkpoint.
const ClaimOutboxSQL = `SELECT id, payload FROM outbox WHERE delivered_at IS NULL ORDER BY created_at FOR UPDATE SKIP LOCKED LIMIT $1`

type Outbox struct{}

func (Outbox) Claim(context.Context, int) error { return nil }

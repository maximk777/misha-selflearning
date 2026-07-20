CREATE TABLE tasks (
  id text PRIMARY KEY,
  idempotency_key text NOT NULL UNIQUE,
  title text NOT NULL CHECK (length(trim(title)) > 0),
  status text NOT NULL CHECK (status IN ('open', 'completed')),
  created_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE outbox (
  id text PRIMARY KEY,
  event_type text NOT NULL,
  payload jsonb NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now(),
  delivered_at timestamptz
);
CREATE INDEX outbox_undelivered_idx ON outbox (created_at) WHERE delivered_at IS NULL;

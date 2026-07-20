CREATE TABLE IF NOT EXISTS customers (id bigserial PRIMARY KEY, email text NOT NULL UNIQUE, region text NOT NULL, created_at timestamptz NOT NULL DEFAULT now());
CREATE TABLE IF NOT EXISTS orders (id bigserial PRIMARY KEY, customer_id bigint NOT NULL REFERENCES customers(id), status text NOT NULL CHECK (status IN ('new','paid','done')), total_cents bigint NOT NULL CHECK (total_cents >= 0), created_at timestamptz NOT NULL DEFAULT now());
CREATE TABLE IF NOT EXISTS jobs (id bigserial PRIMARY KEY, payload jsonb NOT NULL, status text NOT NULL DEFAULT 'new', taken_at timestamptz);
CREATE TABLE IF NOT EXISTS documents (id bigserial PRIMARY KEY, body jsonb NOT NULL, search tsvector GENERATED ALWAYS AS (to_tsvector('simple', body::text)) STORED);
CREATE TABLE IF NOT EXISTS outbox (id uuid PRIMARY KEY, topic text NOT NULL, payload jsonb NOT NULL, created_at timestamptz NOT NULL DEFAULT now(), published_at timestamptz);

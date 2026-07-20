CREATE TABLE IF NOT EXISTS notes(id bigserial PRIMARY KEY, body text NOT NULL); INSERT INTO notes(body) VALUES ('hello'); SELECT * FROM notes;

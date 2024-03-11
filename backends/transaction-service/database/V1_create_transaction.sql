

CREATE TABLE IF NOT EXISTS "transaction" (
  id             uuid,
  user_id        uuid,
  value         numeric,
  created_at     timestamp,
  PRIMARY KEY (id));

CREATE INDEX IF NOT EXISTS "transaction_user_id" ON "transaction" (user_id);

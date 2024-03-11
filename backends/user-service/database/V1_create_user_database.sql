

CREATE TABLE IF NOT EXISTS "user" (
  id             uuid,
  email          text NOT NULL unique,
  surname        text,
  first_name     text,
  PRIMARY KEY (id));

CREATE INDEX IF NOT EXISTS user_email_idx ON "user" (email);

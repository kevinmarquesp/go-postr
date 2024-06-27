-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id              SERIAL UNIQUE,
  username        TEXT UNIQUE,
  password        TEXT,
  session_token   TEXT UNIQUE,
  session_expires TIMESTAMP,
  created_at      TIMESTAMP,
  updated_at      TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd

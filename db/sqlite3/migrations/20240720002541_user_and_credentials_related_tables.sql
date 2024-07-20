-- +goose Up
-- +goose StatementBegin
CREATE TABLE Users (
  id          TEXT UNIQUE NOT NULL CHECK (LENGTH(id) = 26),
  name        TEXT,
  username    TEXT UNIQUE NOT NULL,
  email       TEXT UNIQUE NOT NULL,
  description TEXT,
  created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  password    TEXT CHECK (LENGTH(password) >= 60),

  role TEXT DEFAULT 'STANDARD' CHECK (
    role IS 'STANDARD' OR
    role IS 'MODERATOR' OR
    role IS 'BANNED'
  )
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Users;
-- +goose StatementEnd

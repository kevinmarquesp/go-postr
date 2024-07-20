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

  auth_type TEXT CHECK (
    auth_type IS 'CREDENTIALS' OR
    auth_type IS 'ANONYMOUS' OR
    auth_type IS 'GOOGLE' OR
    auth_type IS 'GITHUB'
  ),

  role TEXT DEFAULT 'STANDARD' CHECK (
    role IS 'STANDARD' OR
    role IS 'MODERATOR' OR
    role IS 'BANNED'
  )
);

CREATE TABLE Credentials (
  userId   TEXT UNIQUE NOT NULL CHECK (LENGTH(userId) = 26),
  password TEXT NOT NULL CHECK (LENGTH(password) >= 60)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Users;
DROP TABLE Credentials;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Users (
  id           TEXT UNIQUE NOT NULL CHECK (LENGTH(id) IS 26),
  username     TEXT NOT NULL UNIQUE CHECK (username IS NOT ''),
  bio          TEXT CHECK (display_name IS NOT ''),
  display_name TEXT CHECK (display_name IS NOT ''),
  email        TEXT CHECK (display_name IS NOT ''),
  verified_at  TIMESTAMP,
  created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

  provider TEXT NOT NULL CHECK (
    provider IS 'Credentials'
  ),

  role TEXT NOT NULL DEFAULT 'Standard' CHECK (
    role IS 'Standard' OR
    role IS 'Banned'
  )
);

CREATE TABLE IF NOT EXISTS Credentials (
  id          INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id     TEXT NOT NULL CHECK (LENGTH(user_id) IS 26),
  email       TEXT NOT NULL,
  password    TEXT NOT NULL CHECK (LENGTH(password) >= 60),
  created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT foreign_key_user_id
    FOREIGN KEY (user_id)
    REFERENCES Users (id)
    ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Users;
DROP TABLE IF EXISTS Credentials;
-- +goose StatementEnd

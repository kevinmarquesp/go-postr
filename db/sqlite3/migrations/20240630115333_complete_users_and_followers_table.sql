-- +goose Up
-- +goose StatementBegin
DROP TABLE users;

CREATE TABLE IF NOT EXISTS users (
  id              INTEGER PRIMARY KEY AUTOINCREMENT,
  public_id       UUID UNIQUE NOT NULL CHECK (public_id IS NOT ''),
  fullname        TEXT NOT NULL CHECK (fullname IS NOT ''),
  username        TEXT UNIQUE NOT NULL CHECK (username IS NOT ''),
  password        TEXT NOT NULL,
  session_token   TEXT UNIQUE NOT NULL CHECK (session_token IS NOT ''),
  session_expires TIMESTAMP NOT NULL,
  description     TEXT,
  followers       INTEGER DEFAULT 0 CHECK (followers >= 0),
  following       INTEGER DEFAULT 0 CHECK (following >= 0),
  articles        INTEGER DEFAULT 0 CHECK (articles >= 0),
  created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS followers (
  id          INTEGER PRIMARY KEY AUTOINCREMENT,
  follower_id INTEGER NOT NULL,
  followed_id INTEGER NOT NULL,
  created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT check_follower_followed
    CHECK (follower_id IS NOT followed_id),

  CONSTRAINT foreign_key_follower
    FOREIGN KEY (follower_id)
    REFERENCES users (id)
    ON DELETE CASCADE,

  CONSTRAINT foreign_key_followed
    FOREIGN KEY (followed_id)
    REFERENCES users (id)
    ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TABLE followers;
-- +goose StatementEnd

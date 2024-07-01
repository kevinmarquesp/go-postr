CREATE TABLE users (
  id              INTEGER PRIMARY KEY AUTOINCREMENT,
  public_id       UUID UNIQUE NOT NULL,
  fullname        TEXT NOT NULL,
  username        TEXT UNIQUE NOT NULL,
  password        TEXT NOT NULL,
  session_token   TEXT UNIQUE NOT NULL,
  session_expires TIMESTAMP NOT NULL,
  description     TEXT,
  followers       INTEGER DEFAULT 0 CHECK (followers >= 0),
  following       INTEGER DEFAULT 0 CHECK (following >= 0),
  articles        INTEGER DEFAULT 0 CHECK (articles >= 0),
  created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE followers (
  id          INTEGER PRIMARY KEY AUTOINCREMENT,
  follower_id INTEGER NOT NULL,
  followed_id INTEGER NOT NULL,
  created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,

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

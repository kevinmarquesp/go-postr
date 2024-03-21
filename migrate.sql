CREATE TABLE IF NOT EXISTS "user" (
	id			SERIAL PRIMARY KEY,
	username	VARCHAR(255) UNIQUE NOT NULL,
	password	VARCHAR(255) NOT NULL,
	bio			TEXT,
	created_at	TIMESTAMP,
	updated_at	TIMESTAMP
);

CREATE TABLE IF NOT EXISTS article (
	id			SERIAL PRIMARY KEY,
	content		TEXT NOT NULL,
	user_id		INTEGER NOT NULL,
	created_at	TIMESTAMP,
	updated_at	TIMESTAMP,

	FOREIGN	KEY (user_id) REFERENCES "user" (id)
);

CREATE TABLE IF NOT EXISTS relationship (
	follower_user	INTEGER NOT NULL,
	followed_user	INTEGER NOT NULL,

	PRIMARY KEY (follower_user, followed_user),
	FOREIGN	KEY (follower_user) REFERENCES "user" (id),
	FOREIGN	KEY (followed_user) REFERENCES "user" (id),
	CHECK (follower_user != followed_user)
);

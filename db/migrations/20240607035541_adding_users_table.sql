-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	username VARCHAR(255) UNIQUE NOT NULL,
	password VARCHAR(255) NOT NULL,
	description TEXT,
	created_at TIMESTAMP(0) NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
	updated_at TIMESTAMP(0) NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc')
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd

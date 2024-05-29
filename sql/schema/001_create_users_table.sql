-- +goose Up
CREATE TABLE users (
		id BIGSERIAL PRIMARY KEY NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE NOT NULL,
		updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
		user_name TEXT NOT NULL
);

-- +goose Down
DROP TABLE users;

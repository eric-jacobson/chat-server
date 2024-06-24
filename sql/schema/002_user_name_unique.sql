-- +goose Up

ALTER TABLE users ADD CONSTRAINT unique_user_name UNIQUE (user_name);

-- +goose Down
ALTER TABLE users DROP CONSTRAINT IF EXISTS unique_user_name;

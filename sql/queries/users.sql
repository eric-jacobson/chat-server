-- name: CreateUser :one
INSERT INTO users (user_name)
VALUES ($1)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE user_name = $1;

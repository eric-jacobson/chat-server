-- name: CreateUser :one
INSERT INTO users (created_at, updated_at, user_name)
VALUES ($1, $2, $3)
RETURNING *;

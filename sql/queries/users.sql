-- name: CreateUser :one
INSERT INTO users (user_name, password_hash)
VALUES ($1, $2)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE user_name = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE user_name = $1;

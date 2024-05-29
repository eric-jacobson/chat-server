// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package db

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (created_at, updated_at, user_name)
VALUES ($1, $2, $3)
RETURNING id, created_at, updated_at, user_name
`

type CreateUserParams struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	UserName  string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.CreatedAt, arg.UpdatedAt, arg.UserName)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserName,
	)
	return i, err
}
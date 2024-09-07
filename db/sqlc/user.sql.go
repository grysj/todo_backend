// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package db

import (
	"context"
)

const countUsers = `-- name: CountUsers :one
SELECT count(*) FROM users
`

func (q *Queries) CountUsers(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countUsers)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    username
) VALUES (
    $1
) RETURNING user_id, username, created_at
`

func (q *Queries) CreateUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, username)
	var i User
	err := row.Scan(&i.UserID, &i.Username, &i.CreatedAt)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE from users WHERE user_id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, userID int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, userID)
	return err
}

const getUser = `-- name: GetUser :one
SELECT user_id, username, created_at FROM users
WHERE user_id=$1
`

func (q *Queries) GetUser(ctx context.Context, userID int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, userID)
	var i User
	err := row.Scan(&i.UserID, &i.Username, &i.CreatedAt)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT user_id, username, created_at FROM users
WHERE username = $1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(&i.UserID, &i.Username, &i.CreatedAt)
	return i, err
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: permission.sql

package db

import (
	"context"
)

const addPermission = `-- name: AddPermission :one
INSERT INTO permissions (
    list_id,
    user_id
) VALUES (
    $1, $2
) RETURNING permission_id, list_id, user_id, created_at
`

type AddPermissionParams struct {
	ListID int32 `json:"list_id"`
	UserID int32 `json:"user_id"`
}

func (q *Queries) AddPermission(ctx context.Context, arg AddPermissionParams) (Permission, error) {
	row := q.db.QueryRowContext(ctx, addPermission, arg.ListID, arg.UserID)
	var i Permission
	err := row.Scan(
		&i.PermissionID,
		&i.ListID,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}

const deletePermission = `-- name: DeletePermission :exec
DELETE FROM permissions WHERE list_id=$1 AND user_id=$2
`

type DeletePermissionParams struct {
	ListID int32 `json:"list_id"`
	UserID int32 `json:"user_id"`
}

func (q *Queries) DeletePermission(ctx context.Context, arg DeletePermissionParams) error {
	_, err := q.db.ExecContext(ctx, deletePermission, arg.ListID, arg.UserID)
	return err
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: list.sql

package db

import (
	"context"
	"time"
)

const createList = `-- name: CreateList :one
INSERT INTO lists (
    created_by,
    title
) VALUES (
    $1, $2
) RETURNING list_id, created_by, title, created_at
`

type CreateListParams struct {
	CreatedBy int32  `json:"created_by"`
	Title     string `json:"title"`
}

func (q *Queries) CreateList(ctx context.Context, arg CreateListParams) (List, error) {
	row := q.db.QueryRowContext(ctx, createList, arg.CreatedBy, arg.Title)
	var i List
	err := row.Scan(
		&i.ListID,
		&i.CreatedBy,
		&i.Title,
		&i.CreatedAt,
	)
	return i, err
}

const editTile = `-- name: EditTile :exec
UPDATE lists SET title=$2
WHERE list_id = $1
`

type EditTileParams struct {
	ListID int32  `json:"list_id"`
	Title  string `json:"title"`
}

func (q *Queries) EditTile(ctx context.Context, arg EditTileParams) error {
	_, err := q.db.ExecContext(ctx, editTile, arg.ListID, arg.Title)
	return err
}

const getList = `-- name: GetList :one
SELECT list_id, created_by, title, created_at from lists
WHERE list_id = $1
`

func (q *Queries) GetList(ctx context.Context, listID int32) (List, error) {
	row := q.db.QueryRowContext(ctx, getList, listID)
	var i List
	err := row.Scan(
		&i.ListID,
		&i.CreatedBy,
		&i.Title,
		&i.CreatedAt,
	)
	return i, err
}

const getListsByUserPermission = `-- name: GetListsByUserPermission :many
SELECT l.list_id, l.title, l.created_at
FROM lists l
JOIN permissions p ON l.list_id = p.list_id
WHERE p.to_user = $1
  AND p.perm_type = $2
`

type GetListsByUserPermissionParams struct {
	ToUser   int32 `json:"to_user"`
	PermType int32 `json:"perm_type"`
}

type GetListsByUserPermissionRow struct {
	ListID    int32     `json:"list_id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

func (q *Queries) GetListsByUserPermission(ctx context.Context, arg GetListsByUserPermissionParams) ([]GetListsByUserPermissionRow, error) {
	rows, err := q.db.QueryContext(ctx, getListsByUserPermission, arg.ToUser, arg.PermType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetListsByUserPermissionRow{}
	for rows.Next() {
		var i GetListsByUserPermissionRow
		if err := rows.Scan(&i.ListID, &i.Title, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

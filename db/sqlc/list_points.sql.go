// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: list_points.sql

package db

import (
	"context"
)

const checkPoint = `-- name: CheckPoint :exec
UPDATE list_points SET checked=true
WHERE point_id = $1
`

func (q *Queries) CheckPoint(ctx context.Context, pointID int32) error {
	_, err := q.db.ExecContext(ctx, checkPoint, pointID)
	return err
}

const createPoint = `-- name: CreatePoint :one
INSERT INTO list_points (
    list_id,
    content,
    position,
    added_by
) VALUES (
    $1, $2, $3, $4
) RETURNING point_id, list_id, content, position, checked, created_at, added_by
`

type CreatePointParams struct {
	ListID   int32  `json:"list_id"`
	Content  string `json:"content"`
	Position int32  `json:"position"`
	AddedBy  int32  `json:"added_by"`
}

func (q *Queries) CreatePoint(ctx context.Context, arg CreatePointParams) (ListPoint, error) {
	row := q.db.QueryRowContext(ctx, createPoint,
		arg.ListID,
		arg.Content,
		arg.Position,
		arg.AddedBy,
	)
	var i ListPoint
	err := row.Scan(
		&i.PointID,
		&i.ListID,
		&i.Content,
		&i.Position,
		&i.Checked,
		&i.CreatedAt,
		&i.AddedBy,
	)
	return i, err
}

const getMaxPositionOrDefault = `-- name: GetMaxPositionOrDefault :one
SELECT COALESCE(MAX(position), 1) AS max_position
FROM list_points
WHERE list_id = $1
`

func (q *Queries) GetMaxPositionOrDefault(ctx context.Context, listID int32) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, getMaxPositionOrDefault, listID)
	var max_position interface{}
	err := row.Scan(&max_position)
	return max_position, err
}

const getPointsByListID = `-- name: GetPointsByListID :many
SELECT point_id, list_id, content, position, checked, created_at, added_by FROM list_points lp
WHERE list_id = $1
ORDER BY position ASC
`

func (q *Queries) GetPointsByListID(ctx context.Context, listID int32) ([]ListPoint, error) {
	rows, err := q.db.QueryContext(ctx, getPointsByListID, listID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListPoint{}
	for rows.Next() {
		var i ListPoint
		if err := rows.Scan(
			&i.PointID,
			&i.ListID,
			&i.Content,
			&i.Position,
			&i.Checked,
			&i.CreatedAt,
			&i.AddedBy,
		); err != nil {
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

const uncheckPoint = `-- name: UncheckPoint :exec
UPDATE list_points SET checked=false
WHERE point_id = $1
`

func (q *Queries) UncheckPoint(ctx context.Context, pointID int32) error {
	_, err := q.db.ExecContext(ctx, uncheckPoint, pointID)
	return err
}

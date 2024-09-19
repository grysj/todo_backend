-- name: CreatePoint :one
INSERT INTO list_points (
    list_id,
    content,
    position,
    added_by
) VALUES (
    $1, $2, $3, $4
) RETURNING *;


-- name: ChangePointCheck :exec
UPDATE list_points SET checked=$2
WHERE point_id = $1;



-- name: GetPointsByListID :many
SELECT * FROM list_points lp
WHERE list_id = $1
ORDER BY position ASC;

-- name: GetMaxPositionOrDefault :one
SELECT COALESCE(MAX(position), 0)::int AS max_position
FROM list_points
WHERE list_id = $1;

-- name: ChangePointPosition :exec
UPDATE list_points SET position = sqlc.arg(new_pos)
WHERE point_id = $1;


-- name: ChangePointContent :exec
UPDATE list_points SET content = sqlc.arg(new_text)
WHERE point_id = $1;

-- name: DeletePoint :exec
DELETE FROM list_points
WHERE point_id = $1;

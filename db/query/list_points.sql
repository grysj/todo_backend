-- name: CreatePoint :one
INSERT INTO list_points (
    list_id,
    content,
    position,
    added_by
) VALUES (
    $1, $2, $3, $4
) RETURNING *;


-- name: CheckPoint :exec
UPDATE list_points SET checked=true
WHERE point_id = $1;


-- name: UncheckPoint :exec
UPDATE list_points SET checked=false
WHERE point_id = $1;


-- name: GetPointsByListID :many
SELECT * FROM list_points lp
WHERE list_id = $1
ORDER BY position ASC;

-- name: GetMaxPositionOrDefault :one
SELECT COALESCE(MAX(position), 1) AS max_position
FROM list_points
WHERE list_id = $1;

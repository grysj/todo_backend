-- name: CreatePoint :one
INSERT INTO list_points (
    list_id,
    point
) VALUES (
    $1, $2
) RETURNING *;


-- name: CheckPoint :exec
UPDATE list_points SET checked=true
WHERE point_id = $1;


-- name: UncheckPoint :exec
UPDATE list_points SET checked=false
WHERE point_id = $1;


-- name: UserPoints :many
SELECT l.list_id, l.title, lp.point_id, lp.point, lp.checked, lp.created_at
FROM lists l
JOIN list_points lp ON l.list_id = lp.list_id
WHERE l.list_id IN (
    SELECT p.list_id
    FROM permissions p
    WHERE p.user_id = $1
)
ORDER BY l.created_at ASC, lp.created_at ASC;

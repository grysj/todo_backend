-- name: CreateList :one
INSERT INTO lists (
    created_by,
    title
) VALUES (
    $1, $2
) RETURNING *;


-- name: EditTile :exec
UPDATE lists SET title=$2
WHERE list_id = $1;

-- name: GetList :one
SELECT * from lists
WHERE list_id = $1;


-- name: GetListsByUserPermission :many
SELECT l.list_id, l.title, l.created_at
FROM lists l
JOIN permissions p ON l.list_id = p.list_id
WHERE p.to_user = $1
  AND p.perm_type = $2;

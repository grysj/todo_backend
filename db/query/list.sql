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

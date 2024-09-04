-- name: AddPermission :one
INSERT INTO permissions (
    list_id,
    user_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: DeletePermission :exec
DELETE FROM permissions WHERE list_id=$1 AND user_id=$2;

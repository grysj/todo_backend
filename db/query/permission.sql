-- name: CreatePermission :one
INSERT INTO permissions (
    from_user,
    to_user,
    list_id,
    perm_type
) VALUES (
    $1, $2, $3, $4
) RETURNING *;


-- name: CheckUserPermissions :one
SELECT * FROM permissions
WHERE to_user = $1;


-- name: CheckUserPermission :one
SELECT COALESCE(4, p.perm_type ) FROM permissions p
WHERE from_user = $1 AND list_id = $2;

-- name: ListPermissions :many
SELECT p.to_user, p.perm_type FROM permissions p
WHERE p.list_id = $1;



-- name: AddPermission :one
INSERT INTO permissions (
    from_user,
    to_user,
    perm_type)
VALUES ($1, $2, $3)
RETURNING *;

-- name: EditPermission :exec
UPDATE permissions SET perm_type = $2
WHERE permission_id = $1;


-- name: DeletePermission :exec
DELETE from permissions p
WHERE p.permission_id = $1;

-- name: CreateUser :one
INSERT INTO users (
    username
) VALUES (
    $1
) RETURNING *;

-- name: DeleteUser :exec
DELETE from users WHERE user_id = $1;


-- name: GetUser :one
SELECT * FROM users
WHERE user_id=$1;

-- name: CountUsers :one
SELECT count(*) FROM users;

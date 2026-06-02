-- name: GetUser :one
SELECT * FROM users
WHERE ID = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY last_name;

-- name: CreateUser :one
INSERT INTO users (
    first_name,
    last_name
) VALUES (
    $1,
    $2
)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
    SET first_name = $2,
    last_name = $3,
    image_url = $4
WHERE id = $1
RETURNING *;

-- name: DeleteUser :one
UPDATE users
    SET deleted_at = NOW()
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

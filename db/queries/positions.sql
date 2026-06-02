-- name: GetPosition :one
SELECT * FROM positions
WHERE id = $1 AND deleted_at IS NULL LIMIT 1;

-- name: ListPositions :many
SELECT * FROM positions
WHERE deleted_at IS NULL
ORDER BY name;

-- name: CreatePosition :one
INSERT INTO positions (
    name,
    email,
    group_id
) VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: UpdatePosition :one
UPDATE positions
    SET name = $2,
    email = $3,
    group_id = $4
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeletePosition :one
UPDATE positions
    SET deleted_at = NOW()
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

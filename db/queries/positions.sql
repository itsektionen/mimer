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
    group_id,
    established_at,
    dissolved_at
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: UpdatePosition :one
UPDATE positions
    SET name = $2,
    email = $3,
    group_id = $4,
    established_at = $5,
    dissolved_at = $6
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeletePosition :one
UPDATE positions
    SET deleted_at = NOW()
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: GetPosition :one
SELECT * FROM position
WHERE id = $1 AND deleted_at IS NULL LIMIT 1;

-- name: ListPositions :many
SELECT * FROM position
WHERE deleted_at IS NULL
ORDER BY name;

-- name: CreatePosition :one
INSERT INTO position (
    name,
    email,
    committee_id
) VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: UpdatePosition :one
UPDATE position
    SET name = $2,
    email = $3,
    committee_id = $4
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeletePosition :one
UPDATE position
    SET deleted_at = NOW()
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: GetPosition :one
SELECT * FROM position
WHERE ID = $1 LIMIT 1;

-- name: ListPositions :many
SELECT * FROM position
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
WHERE ID = $1
RETURNING *;

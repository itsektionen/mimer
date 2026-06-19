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

-- name: ListGroupPositionsWithActiveTrustees :many
SELECT
    pos.id as position_id,
    pos.name as position_name,
    pos.email as position_email,
    pos.group_id,
    pos.established_at as position_established_at,
    pos.dissolved_at as position_dissolved_at,
    t.id as trustee_id,
    t.start_date,
    t.end_date,
    u.id as user_id,
    u.first_name,
    u.last_name,
    u.image_url as user_image_url
FROM positions pos
LEFT JOIN trustees t ON t.position_id = pos.id 
    AND t.deleted_at IS NULL 
    AND (t.end_date IS NULL OR t.end_date > CURRENT_DATE)
    AND t.start_date <= CURRENT_DATE
LEFT JOIN users u ON t.user_id = u.id AND u.deleted_at IS NULL
WHERE pos.group_id = $1
    AND pos.deleted_at IS NULL
    AND (pos.dissolved_at IS NULL OR pos.dissolved_at > CURRENT_DATE)
ORDER BY pos.name ASC;

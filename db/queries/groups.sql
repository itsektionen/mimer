-- name: GetGroup :one
SELECT * FROM groups
WHERE ID = $1 AND deleted_at IS NULL AND active = TRUE LIMIT 1;

-- name: ListGroups :many
SELECT * FROM groups
WHERE deleted_at IS NULL
ORDER BY name;

-- name: CreateGroup :one
INSERT INTO groups (
    name,
    slug,
    short_name,
    description,
    color,
    image_url,
    website_url,
    established_at,
    dissolved_at
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9
)
RETURNING *;

-- name: UpdateGroup :one
UPDATE groups
    SET name = $2,
    slug = $3,
    short_name = $4,
    description = $5,
    color = $6,
    image_url = $7,
    website_url = $8,
    established_at = $9,
    dissolved_at = $10
WHERE ID = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteGroup :one
UPDATE groups
    SET deleted_at = NOW()
WHERE ID = $1 AND deleted_at IS NULL
RETURNING *;

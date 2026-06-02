-- name: GetGroup :one
SELECT * FROM group
WHERE ID = $1 AND deleted_at IS NULL AND active = TRUE LIMIT 1;

-- name: ListGroups :many
SELECT * FROM group
WHERE deleted_at IS NULL
ORDER BY name;

-- name: CreateGroup :one
INSERT INTO group (
    name,
    slug,
    short_name,
    description,
    color,
    image_url,
    website_url
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)
RETURNING *;

-- name: UpdateGroup :one
UPDATE group
    SET name = $2,
    slug = $3,
    short_name = $4,
    description = $5,
    color = $6,
    image_url = $7,
    website_url = $8
WHERE ID = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteGroup :one
UPDATE group
    SET deleted_at = NOW()
WHERE ID = $1 AND deleted_at IS NULL
RETURNING *;

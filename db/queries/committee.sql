-- name: GetCommittee :one
SELECT * FROM committee
WHERE ID = $1 LIMIT 1;

-- name: ListCommittees :many
SELECT * FROM committee
ORDER BY name;

-- name: CreateCommittee :one
INSERT INTO committee (
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

-- name: UpdateCommittee :one
UPDATE committee
    SET name = $2,
    slug = $3,
    short_name = $4,
    description = $5,
    color = $6,
    image_url = $7,
    website_url = $8
WHERE ID = $1
RETURNING *;

-- name: ListApiKeys :many
SELECT * FROM api_key
ORDER BY created_at;

-- name: GetApiKey :one
SELECT * FROM api_key
WHERE ID = $1;

-- name: CreateApiKey :one
INSERT INTO api_key (
    value
), (
    $1
)
RETURNING *;

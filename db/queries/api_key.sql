-- name: ListApiKeys :many
SELECT * FROM api_key
ORDER BY created_at;

-- name: GetApiKey :one
SELECT * FROM api_key
WHERE ID = $1;

-- name: GetApiKeyByValue :one
SELECT * FROM api_key
WHERE value = $1;

-- name: CreateApiKey :one
INSERT INTO api_key (
    value
) VALUES (
    $1
)
RETURNING *;

-- name: DisableApiKey :exec
UPDATE api_key
    SET active = false
WHERE id = $1;

-- name: EnableApiKey :exec
UPDATE api_key
    SET active = false
WHERE id = $1;

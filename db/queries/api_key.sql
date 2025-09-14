-- name: ListApiKeys :many
SELECT * FROM api_key
ORDER BY created_at AND deleted_at IS NULL;

-- name: GetApiKey :one
SELECT * FROM api_key
WHERE ID = $1 AND deleted_at IS NULL;

-- name: GetApiKeyByValue :one
SELECT * FROM api_key
WHERE value = $1 AND deleted_at IS NULL;

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
WHERE id = $1 AND deleted_at IS NULL;

-- name: EnableApiKey :exec
UPDATE api_key
    SET active = false
WHERE id = $1 AND deleted_at IS NULL;

-- name: DeleteApiKey :one
UPDATE api_key
    SET deleted_at = NOW()
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

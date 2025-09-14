-- name: GetPerson :one
SELECT * FROM person
WHERE ID = $1 LIMIT 1;

-- name: ListPeople :many
SELECT * FROM person
ORDER BY last_name;

-- name: CreatePerson :one
INSERT INTO person (
    first_name,
    last_name
) VALUES (
    $1,
    $2
)
RETURNING *;

-- name: UpdatePerson :one
UPDATE person
    SET first_name = $2,
    last_name = $3,
    image_url = $4
WHERE id = $1
RETURNING *;

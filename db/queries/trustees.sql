-- name: ListGroupTrustees :many
SELECT
    t.id as trustee_id,
    t.start_date,
    t.end_date,
    t.created_at as trustee_created_at,
    t.updated_at as trustee_updated_at,
    p.id as user_id,
    p.first_name,
    p.last_name,
    p.image_url as user_image_url,
    pos.id as position_id,
    pos.name as position_name,
    pos.email as position_email,
    pos.group_id
FROM trustees t
INNER JOIN users p ON t.user_id = p.id
INNER JOIN positions pos ON t.position_id = pos.id
WHERE pos.group_id = $1
    AND t.deleted_at IS NULL
    AND p.deleted_at IS NULL
    AND pos.deleted_at IS NULL
ORDER BY t.start_date DESC, pos.name ASC;

-- name: ListTrustees :many
SELECT
  t.id trustee_id,
  t.start_date,
  t.end_date,
  u.id user_id,
  u.first_name,
  u.last_name,
  pos.id position_id,
  pos.name position_name,
  g.name group_name,
  g.id group_id
FROM trustees t
INNER JOIN users u ON t.user_id = u.id
INNER JOIN positions pos ON t.position_id = pos.id
INNER JOIN groups g ON pos.group_id = g.id;

-- name: CreateTrustee :one
INSERT INTO trustees (
  user_id,
  position_id,
  start_date,
  end_date
) VALUES (
  $1,
  $2,
  $3,
  $4
) RETURNING *;

-- name: ListTrusteesByPosition :many
SELECT
  t.id trustee_id,
  t.start_date,
  t.end_date,
  u.id user_id,
  u.first_name,
  u.last_name,
  p.id position_id,
  p.name position_name,
  g.name group_name,
  g.id group_id
FROM positions p 
INNER JOIN trustees t ON t.position_id = p.id
INNER JOIN users u ON t.user_id = u.id
INNER JOIN groups g on p.group_id = g.id
WHERE p.id = $1
ORDER BY t.start_date DESC;

-- name: ListCommitteeTrustees :many
SELECT
    t.id as trustee_id,
    t.start_date,
    t.end_date,
    t.created_at as trustee_created_at,
    t.updated_at as trustee_updated_at,
    p.id as person_id,
    p.first_name,
    p.last_name,
    p.image_url as person_image_url,
    pos.id as position_id,
    pos.name as position_name,
    pos.email as position_email,
    pos.active as position_active,
    pos.committee_id
FROM trustee t
INNER JOIN person p ON t.person_id = p.id
INNER JOIN position pos ON t.position_id = pos.id
WHERE pos.committee_id = $1
    AND t.deleted_at IS NULL
    AND p.deleted_at IS NULL
    AND pos.deleted_at IS NULL
ORDER BY t.start_date DESC, pos.name ASC;

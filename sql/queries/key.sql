-- name: CreateSCKey :one
INSERT INTO sc_keys(
    hash_key, block_key, created_at, signing_revoked_at, valid_until
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetSecureCookieKey :many
SELECT * FROM sc_keys
ORDER BY created_at ASC;

-- name: DeleteOldestSCKey :one
DELETE FROM sc_keys
WHERE created_at = MIN(created_at)
RETURNING *;

-- name: RevokeSCSigning :one
UPDATE sc_keys
SET signing_revoked_at = $1
WHERE token = $1
RETURNING *;

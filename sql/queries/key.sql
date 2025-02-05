-- name: CreateKey :one
INSERT INTO keys(
    token, created_at, signing_revoked_at, valid_until, type
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetSecureCookieKey :many
SELECT * FROM keys
WHERE type = "secure_cookie";

-- name: GetJWTKey :many
SELECT * FROM keys
WHERE type = "jwt";

-- name: DeleteOldestKeyOfType :one
DELETE FROM keys
WHERE created_at = MIN(created_at)
AND type = $1
RETURNING *;

-- name: GetKeysOfType :many
SELECT * FROM keys 
WHERE type = $1
LIMIT 2;

-- name: RevokeSigningOfType :one
UPDATE keys
SET signing_revoked_at = $1
WHERE token = $1 AND type = $2
RETURNING *;

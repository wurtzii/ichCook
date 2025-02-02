-- name: IsTokenRevoked :one
SELECT * FROM refresh_tokens 
WHERE token = $1 and revoked_at <> NULL;

-- name: RevokeToken :exec
UPDATE refresh_tokens SET revoked_at = $1
WHERE token = $2;

-- name: CreateToken :one
INSERT INTO refresh_tokens(token, user_id, valid_from, valid_until)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: RevokeAllTokens :many
UPDATE refresh_tokens SET revoked_at = $1
WHERE user_id = $1
RETURNING *;

-- name: RemoveRevokedTokens :many
DELETE FROM refresh_tokens
WHERE user_id = $1 AND revoked_at <> NULL
RETURNING *;

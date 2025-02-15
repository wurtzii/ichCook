// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: refresh_token.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createToken = `-- name: CreateToken :one
INSERT INTO refresh_tokens(token, user_id, valid_from, valid_until)
VALUES ($1, $2, $3, $4)
RETURNING token, valid_from, valid_until, revoked_at, user_id
`

type CreateTokenParams struct {
	Token      string
	UserID     int32
	ValidFrom  time.Time
	ValidUntil time.Time
}

func (q *Queries) CreateToken(ctx context.Context, arg CreateTokenParams) (RefreshToken, error) {
	row := q.db.QueryRowContext(ctx, createToken,
		arg.Token,
		arg.UserID,
		arg.ValidFrom,
		arg.ValidUntil,
	)
	var i RefreshToken
	err := row.Scan(
		&i.Token,
		&i.ValidFrom,
		&i.ValidUntil,
		&i.RevokedAt,
		&i.UserID,
	)
	return i, err
}

const isTokenRevoked = `-- name: IsTokenRevoked :one
SELECT token, valid_from, valid_until, revoked_at, user_id FROM refresh_tokens 
WHERE token = $1 and revoked_at <> NULL
`

func (q *Queries) IsTokenRevoked(ctx context.Context, token string) (RefreshToken, error) {
	row := q.db.QueryRowContext(ctx, isTokenRevoked, token)
	var i RefreshToken
	err := row.Scan(
		&i.Token,
		&i.ValidFrom,
		&i.ValidUntil,
		&i.RevokedAt,
		&i.UserID,
	)
	return i, err
}

const removeRevokedTokens = `-- name: RemoveRevokedTokens :many
DELETE FROM refresh_tokens
WHERE user_id = $1 AND revoked_at <> NULL
RETURNING token, valid_from, valid_until, revoked_at, user_id
`

func (q *Queries) RemoveRevokedTokens(ctx context.Context, userID int32) ([]RefreshToken, error) {
	rows, err := q.db.QueryContext(ctx, removeRevokedTokens, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RefreshToken
	for rows.Next() {
		var i RefreshToken
		if err := rows.Scan(
			&i.Token,
			&i.ValidFrom,
			&i.ValidUntil,
			&i.RevokedAt,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const revokeAllTokens = `-- name: RevokeAllTokens :many
UPDATE refresh_tokens SET revoked_at = $1
WHERE user_id = $2
RETURNING token, valid_from, valid_until, revoked_at, user_id
`

type RevokeAllTokensParams struct {
	RevokedAt sql.NullTime
	UserID    int32
}

func (q *Queries) RevokeAllTokens(ctx context.Context, arg RevokeAllTokensParams) ([]RefreshToken, error) {
	rows, err := q.db.QueryContext(ctx, revokeAllTokens, arg.RevokedAt, arg.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RefreshToken
	for rows.Next() {
		var i RefreshToken
		if err := rows.Scan(
			&i.Token,
			&i.ValidFrom,
			&i.ValidUntil,
			&i.RevokedAt,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const revokeToken = `-- name: RevokeToken :exec
UPDATE refresh_tokens SET revoked_at = $1
WHERE token = $2
`

type RevokeTokenParams struct {
	RevokedAt sql.NullTime
	Token     string
}

func (q *Queries) RevokeToken(ctx context.Context, arg RevokeTokenParams) error {
	_, err := q.db.ExecContext(ctx, revokeToken, arg.RevokedAt, arg.Token)
	return err
}

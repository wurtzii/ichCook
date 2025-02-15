// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"time"
)

type Recipe struct {
	ID           int32
	Ingredients  sql.NullString
	Instructions sql.NullString
}

type RecipeFollow struct {
	ID       int32
	RecipeID int32
	UserID   int32
}

type RefreshToken struct {
	Token      string
	ValidFrom  time.Time
	ValidUntil time.Time
	RevokedAt  sql.NullTime
	UserID     int32
}

type ScKey struct {
	HashKey          []byte
	BlockKey         []byte
	CreatedAt        time.Time
	ValidUntil       time.Time
	SigningRevokedAt sql.NullTime
}

type User struct {
	ID       int32
	Username sql.NullString
	Password sql.NullString
}

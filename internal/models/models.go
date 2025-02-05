package models

import (
    "github.com/golang-jwt/jwt/v5"
    "time"
    "github.com/gorilla/securecookie"
)

type CustomClaims struct {
    Value interface{} `json:"value"`
    jwt.RegisteredClaims
}

type RefreshKey struct {
    Token       securecookie.SecureCookie
    ValidUntil  time.Time
    ValidFrom   time.Time
    SigningRevoked bool
}

type SCKey struct {
    Token            interface{} `json:"token"`
    Type             string     `json:"type"`
    CreatedAt        time.Time  `json:"created_at"`
	ValidUntil       time.Time  `json:"valid_until"`
    SigningRevokedAt time.Time  `json:"signing_revoked_at"`
}

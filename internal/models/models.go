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

type TimedSC struct {
    SC              *securecookie.SecureCookie
    CreatedAt       time.Time  `json:"created_at"`
	ValidUntil      time.Time  `json:"valid_until"`
}

package handlers

import(
    "ichCook/internal/database"
    "github.com/gorilla/securecookie"
)

type ApiConfig struct {
    DQ      *database.Queries
    OldSC   securecookie.SecureCookie
    CurrSC  securecookie.SecureCookie
}

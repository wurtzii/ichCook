package main

import (
    "ichCook/internal/database"
    "ichCook/internal/models"
    _ "github.com/lib/pq"
    "bytes"
    "github.com/joho/godotenv"
    "time"
    "os"
    "log"
    "encoding/gob"
    "context"
    "ichCook/internal/handlers"
    "github.com/gorilla/securecookie"
    "database/sql"
)

func main() {
    godotenv.Load()
    conn_str, ok := os.LookupEnv("CONN_STR")
    if !ok {
        log.Fatal("database connection string not found")
    }

    dbtx, err := sql.Open("postgres", conn_str)
    if err != nil {
        log.Fatal(err)
    }
    queries := database.New(dbtx)
    sc_keys, err := queries.GetSecureCookieKey(context.Background());

    if err != nil {
    }
    gob.Register(securecookie.SecureCookie{})
    
    queries.GetKeysOfType()
    for len(sc_keys) < 2 {
        // create refresh keys until we have 2
        rk := models.RefreshKey{
            Token: securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32)),
            ValidFrom: time.Now(),
            ValidUntil: time.Now().AddDate(0, 0, 7),
            SigningRevoked: false,
        }
    }

    decoder := gob.NewDecoder()
    decoder.Decode(sc_keys[0].)

    _ = handlers.ApiConfig{ 
        DQ:     queries,
        OldSC: sc_keys[0].,
        CurrSC: sc_keys[1],
    }
}

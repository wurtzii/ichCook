package main

import (
    "ichCook/internal/database"
    "ichCook/internal/models"
    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
    "time"
    "os"
    "log"
    "encoding/gob"
    "context"
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
    keys, err := queries.GetSecureCookieKey(context.Background());
    tsc := []models.TimedSC{}
    gob.Register(&securecookie.SecureCookie{})
    for _, key := range keys {
        sc := securecookie.New(key.HashKey, key.BlockKey)
        tsc = append(tsc, models.TimedSC{
            SC: sc,
            CreatedAt: key.CreatedAt,
            ValidUntil: key.ValidUntil,
        })
    }

    if len(tsc) < 2 { // create a new sc key if available
        hashKey := securecookie.GenerateRandomKey(64)
        blockKey := securecookie.GenerateRandomKey(32)
        sc := securecookie.New(hashKey, blockKey)
        dckp := database.CreateSCKeyParams{
            HashKey: hashKey,
            BlockKey: blockKey,
            CreatedAt: time.Now(),
            ValidUntil: time.Now().AddDate(0, 0, 7),
        }

        _, err := queries.CreateSCKey(context.Background(), dckp)

        if err != nil {
            log.Fatal(err)
        }

        tsc = append(tsc, models.TimedSC{
            SC: sc,
            CreatedAt: dckp.CreatedAt,
            ValidUntil: dckp.ValidUntil,
        })
    }
}

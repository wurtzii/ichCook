package lib

import (
    "context"
    "encoding/gob"
    "time"
    "bytes"
    "ichCook/internal/database"
    "ichCook/internal/handlers"
    "github.com/gorilla/securecookie"
)

func Encode(v interface{}) ([]byte, error) {
    var buf bytes.Buffer
    enc := gob.NewEncoder(&buf)
    err := enc.Encode(v)
    return buf.Bytes(), err
}

func Decode(dat []byte, t interface{}) error {
    buf := bytes.NewBuffer(dat)
    dec := gob.NewDecoder(buf)
    err := dec.Decode(&t)
    return err
}

func RotateSCKeys(cfg *handlers.ApiConfig, q *database.Queries) (database.ScKey, error) {
    key, err := q.DeleteOldestSCKey(context.Background())
    if err != nil {
        return database.ScKey{}, err
    }

    hashKey := securecookie.GenerateRandomKey(64)
    blockKey := securecookie.GenerateRandomKey(32)

    params := database.CreateSCKeyParams{
        HashKey: hashKey,
        BlockKey: blockKey,
        CreatedAt: time.Now(),
        ValidUntil: time.Now().AddDate(0, 0, 7),
    }
    
    key, err = q.CreateSCKey(context.Background(), params)
    if err != nil {
        return key, err
    }
    sc := securecookie.New(hashKey, blockKey)
    cfg.OldSC = cfg.CurrSC
    cfg.CurrSC = sc

    return key, nil
}

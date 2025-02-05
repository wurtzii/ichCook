package handlers 

import (
    "net/http"
    "ichCook/internal/auth"
    "ichCook/internal/models"
    "time"
    "strconv"
)

type authedHandler func(w http.ResponseWriter, r *http.Request, userid int)

func MiddlewareAuthenticate(handler authedHandler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        authToken, err := auth.GetAuthorizationToken(r)
        if err != nil {
            http.Error(w, "not authorized", http.StatusUnauthorized)
            return
        }

        claims := models.CustomClaims{}
        auth.ReadJWT(authToken, &claims)

        expTime, err := claims.GetExpirationTime()
        if err != nil {
            http.Error(w, "not authorized", http.StatusUnauthorized)
            return
        }

        if expTime.Time.After(time.Now()) {
            http.Error(w, "jwt expired", http.StatusUnauthorized)
            return
        }

        useridStr, err := claims.GetSubject()
        if err != nil {
            http.Error(w, "user id could not be fetched", http.StatusBadRequest)
            return
        }

        userid, err := strconv.Atoi(useridStr)
        if err != nil {
            http.Error(w, "invalid user id", http.StatusBadRequest)
            return
        }

        handler(w, r, userid)
    }
}

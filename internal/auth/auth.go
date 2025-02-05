package auth

import (
    "github.com/gorilla/securecookie"
    "net/http"
    "errors"
    "github.com/golang-jwt/jwt/v5"
    "ichCook/internal/models"
    "time"
)

var (
    ErrInvalidNumSCCookies = errors.New("invalid number of secure cookies")
    ErrInvalidClaims = errors.New("invalid claims")
)

func WriteSCookie(w http.ResponseWriter, name string, value interface{}, sc *securecookie.SecureCookie) (error) {
    encoded, err := securecookie.EncodeMulti(name, value, sc)
    if err != nil {
        return err
    }

    cookie := &http.Cookie{
        Name: name,
        Value: encoded,
        Path: "/",
        HttpOnly: true,
        SameSite: http.SameSiteStrictMode,
    }

    http.SetCookie(w, cookie)
    return nil
}

func ReadSCookie(w http.ResponseWriter, r *http.Request, name string, cookies []*securecookie.SecureCookie) (map[string]string, error) {
    value := make(map[string]string)
    if len(cookies) != 2 {
        return  value, ErrInvalidNumSCCookies
    }

	cookie, err := r.Cookie(name);
    if err != nil {
        return value, err
    }

    err = securecookie.DecodeMulti(name, cookie.Value, &value, cookies[0], cookies[1])
    return value, err
}

func CreateJWT(jwt_secret []byte, value interface{}, subject string, expiry time.Time) (string, error) {
    claims := models.CustomClaims{
        value,
        jwt.RegisteredClaims{
            Issuer:     "ichCook",
            Subject:    subject,
            IssuedAt:   jwt.NewNumericDate(time.Now()),
            ExpiresAt:  jwt.NewNumericDate(expiry),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SigningString()
}

func ReadJWT(token_string string, value interface{}) (error) {
    claims := models.CustomClaims{}
    token, err := jwt.ParseWithClaims(token_string, &claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(token_string), nil
    }, jwt.WithLeeway(5*time.Second))

    if err != nil {
        return err
    } 

    value = token.Claims
    return  nil
}

func CreateNewSecureCookie() *securecookie.SecureCookie {
    return securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))
}

func GetAuthorizationToken(r *http.Request) (string, error) {
    return "", nil
}

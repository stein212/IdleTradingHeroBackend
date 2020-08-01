package utils

import (
	"net/http"

	"github.com/IdleTradingHeroServer/constants"
	"github.com/gorilla/securecookie"
)

func CreateJWTCookieExtractor(secureCookie *securecookie.SecureCookie) func(r *http.Request) (string, error) {
	return func(r *http.Request) (string, error) {
		jwtCookie, err := r.Cookie(constants.CookieAuthName)

		if err != nil {
			return "", nil // no error just no token
		}

		var jwt string

		err = secureCookie.Decode(constants.CookieAuthName, jwtCookie.Value, &jwt)

		if err != nil {
			return "", err
		}

		return jwt, nil
	}
}

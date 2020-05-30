package middlewares

import (
	"log"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/julienschmidt/httprouter"
)

func CheckJWT(jwtMiddleware *jwtmiddleware.JWTMiddleware, next func(http.ResponseWriter, *http.Request, httprouter.Params)) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		err := jwtMiddleware.CheckJWT(w, r)

		if err != nil {
			log.Println(err)
			return
		}

		next(w, r, params)
	}
}

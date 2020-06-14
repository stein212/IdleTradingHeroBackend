package handlers

import (
	"database/sql"

	idletradinghero "github.com/IdleTradingHeroServer/route/github.com/idletradinghero/v2"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/gorilla/securecookie"
)

type ControllerConfig struct {
	DB                  *sql.DB
	JWTSecretKey        string
	JWTMiddleware       *jwtmiddleware.JWTMiddleware
	SecureCookie        *securecookie.SecureCookie
	Domain              string
	PythonServiceClient idletradinghero.RouteClient
}

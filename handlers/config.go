package handlers

import (
	"database/sql"

	strategy_proto "github.com/IdleTradingHeroServer/pb/strategy"
	"github.com/IdleTradingHeroServer/repositories"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/gorilla/securecookie"
)

type ControllerConfig struct {
	JWTSecretKey  string
	JWTMiddleware *jwtmiddleware.JWTMiddleware
	SecureCookie  *securecookie.SecureCookie
	Domain        string

	StrategyClient strategy_proto.StrategyServiceClient

	DB *sql.DB

	StrategyRepository repositories.StrategyRepository
}

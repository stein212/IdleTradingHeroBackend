package main

import (
	"database/sql"

	"github.com/IdleTradingHeroServer/handlers"
	"github.com/IdleTradingHeroServer/middlewares"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/julienschmidt/httprouter"
)

type RouterConfig struct {
	DB            *sql.DB
	JWTSecretKey  string
	JWTMiddleware *jwtmiddleware.JWTMiddleware
}

func SetupRoutes(router *httprouter.Router, config *RouterConfig) {
	router.GET("/", handlers.Index)

	authController := handlers.AuthController{}
	authController.Init(config.DB, config.JWTSecretKey)
	router.POST("/login", authController.GetAccessTokenByPassword)
	router.POST("/register", authController.Register)

	router.GET("/admin/users", handlers.GetUsers(config.DB))

	// test jwt
	router.GET("/users", middlewares.CheckJWT(config.JWTMiddleware, handlers.GetUsers(config.DB)))

}

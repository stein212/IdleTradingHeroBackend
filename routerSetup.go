package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/IdleTradingHeroServer/handlers"
	"github.com/IdleTradingHeroServer/middlewares"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

type RouterSetup struct {
	config *handlers.ControllerConfig
}

func NewRouterSetup(config *handlers.ControllerConfig) *RouterSetup {
	return &RouterSetup{
		config: config,
	}
}

func (rs *RouterSetup) SetupRoutes(router *httprouter.Router) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   getAllowedOrigins(),
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	router.GET("/", handlers.Index)

	// auth
	authController := handlers.NewAuthController(rs.config)
	router.POST("/login", authController.GetAccessTokenByPassword)
	router.POST("/loginc", authController.GetCookieAuthByPassword)
	router.POST("/register", authController.Register)

	router.GET("/admin/users", handlers.GetUsers(rs.config.DB))

	// user
	userController := handlers.NewUserController(rs.config)
	router.GET("/UserInfo", rs.protect(userController.GetUserInfo))

	// strategy
	strategyController := handlers.NewStrategyController(rs.config)
	router.POST("/strategies/macd", rs.protect(strategyController.CreateMacdStrategy))

	router.GET("/strategies", rs.protect(strategyController.GetStrategies))
	router.GET("/strategies/info/:strategyID", rs.protect(strategyController.GetMacdStrategy))
	// router.POST("/strategies/macd/backtest", rs.protect(strategyController.GetStrategyPerformance))
	router.POST("/strategies/initialise/:strategyID", rs.protect(strategyController.InitialiseStrategy))
	router.POST("/strategies/start/:strategyID", rs.protect(strategyController.StartStrategy))
	router.POST("/strategies/pause/:strategyID", rs.protect(strategyController.PauseStrategy))
	router.GET("/strategies/getdata/:strategyID/:length", rs.protect(strategyController.GetStrategyData))

	// test jwt
	router.GET("/users", rs.protect(handlers.GetUsers(rs.config.DB)))

	return c.Handler(router)
}

func (rs *RouterSetup) protect(handler func(http.ResponseWriter, *http.Request, httprouter.Params)) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return middlewares.CheckJWT(rs.config.JWTMiddleware, handler)
}

func getAllowedOrigins() []string {
	return strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
}

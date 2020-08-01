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
		AllowedMethods:   getAllowedMethods(),
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
		// AllowedHeaders: getAllowedHeaders(),
	})

	router.GET("/admin/users", handlers.GetUsers(rs.config.DB))
	// reset all strategies to notDeployed
	router.GET("/admin/resetStrategiesStatus", handlers.ResetStrategiesStatus(rs.config.DB))

	router.GET("/", handlers.Index)

	// auth
	authController := handlers.NewAuthController(rs.config)
	router.POST("/login", authController.GetAccessTokenByPassword)
	router.POST("/loginc", authController.GetCookieAuthByPassword)
	router.POST("/register", authController.Register)

	// user
	userController := handlers.NewUserController(rs.config)
	router.GET("/UserInfo", rs.protect(userController.GetUserInfo))

	// strategy
	strategyController := handlers.NewStrategyController(rs.config)
	router.POST("/strategies/macd", rs.protect(strategyController.CreateMacdStrategy))
	router.POST("/strategies/mfi", rs.protect(strategyController.CreateMfiStrategy))
	router.POST("/strategies/rsi", rs.protect(strategyController.CreateRsiStrategy))

	router.GET("/strategies", rs.protect(strategyController.GetStrategies))
	router.GET("/strategies/macd/info/:strategyID", rs.protect(strategyController.GetMacdStrategy))
	router.GET("/strategies/mfi/info/:strategyID", rs.protect(strategyController.GetMfiStrategy))
	router.GET("/strategies/rsi/info/:strategyID", rs.protect(strategyController.GetRsiStrategy))

	router.POST("/strategies/initialise/:strategyID", rs.protect(strategyController.InitialiseStrategy))
	router.POST("/strategies/start/:strategyType/:strategyID", rs.protect(strategyController.StartStrategy))
	router.POST("/strategies/pause/:strategyType/:strategyID", rs.protect(strategyController.PauseStrategy))
	router.GET("/strategies/getData/:strategyID/:length", rs.protect(strategyController.GetStrategyData))
	router.GET("/strategies/getIndicatorData/:strategyID/:length", rs.protect(strategyController.GetIndicatorData))
	router.GET("/strategies/getPerformanceData/:strategyID/:length", rs.protect(strategyController.GetPerformanceData))

	router.DELETE("/strategies/:strategyType/:strategyID", rs.protect(strategyController.DeleteStrategy))

	strategyEventHookController := handlers.NewStrategyEventHookController(rs.config)
	router.POST("/strategies/events", strategyEventHookController.CreateStrategyEvent)
	router.GET("/strategies/events/:strategyID/:length/:to", rs.protect(strategyEventHookController.GetStrategyEvents))
	router.GET("/strategies/allEvents/:length/:to", rs.protect(strategyEventHookController.GetUserStrategiesEvents))

	// test jwt
	router.GET("/users", rs.protect(handlers.GetUsers(rs.config.DB)))

	return c.Handler(router)
}

func (rs *RouterSetup) protect(handler func(http.ResponseWriter, *http.Request, httprouter.Params)) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return middlewares.CheckJWT(rs.config.JWTMiddleware, handler)
}

func getAllowedOrigins() []string {
	return strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
	// return []string{"*"}
}

func getAllowedMethods() []string {
	return []string{"GET", "POST", "PUT", "DELETE"}
}

func getAllowedHeaders() []string {
	return []string{"*"}
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/IdleTradingHeroServer/constants"
	"github.com/IdleTradingHeroServer/handlers"
	strategy_proto "github.com/IdleTradingHeroServer/pb/strategy"
	"github.com/IdleTradingHeroServer/repositories"
	"github.com/IdleTradingHeroServer/utils"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/securecookie"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	// Database config
	connectionString := os.Getenv(constants.EnvPsqlConnectionString)
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	// Cookie config
	cookieHashKey := []byte(os.Getenv(constants.EnvCookieHashKey))
	cookieBlockKey := []byte(os.Getenv(constants.EnvCookieBlockKey))
	sc := securecookie.New(cookieHashKey, cookieBlockKey)

	// JWT config
	jwtSecretKey := os.Getenv(constants.EnvJWTSecretKey)
	fromCookieAuth := utils.CreateJWTCookieExtractor(sc)

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecretKey), nil
		},
		// When set, the middleware verifies that tokens are signed with the specific signing algorithm
		// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
		// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		SigningMethod: jwt.SigningMethodHS256,
		Extractor: jwtmiddleware.FromFirst(
			fromCookieAuth,
			jwtmiddleware.FromAuthHeader),
	})

	// Get domain
	domain := os.Getenv(constants.EnvDomain)

	// Setup Python service
	strategyGRPCHost := os.Getenv(constants.EnvStrategyGRPCHost)
	dialOption := grpc.WithInsecure()
	fmt.Println(strategyGRPCHost)
	conn, err := grpc.Dial(strategyGRPCHost, dialOption)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	strategyClient := strategy_proto.NewStrategyServiceClient(conn)

	// Setup repositories
	strategyRepo := repositories.NewStrategyRepository(db, strategyClient)

	router := httprouter.New()
	routerConfig := &handlers.ControllerConfig{
		JWTSecretKey:  jwtSecretKey,
		JWTMiddleware: jwtMiddleware,
		SecureCookie:  sc,
		Domain:        domain,

		StrategyClient: strategyClient,

		DB: db,

		StrategyRepository: strategyRepo,
	}
	routerSetup := NewRouterSetup(routerConfig)
	handler := routerSetup.SetupRoutes(router)

	log.Println("Starting server at port 3000")
	http.ListenAndServe(":3000", handler)
}

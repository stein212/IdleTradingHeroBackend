package main

import (
	"fmt"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/IdleTradingHeroServer/utils"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func main() {
	// Database config
	connectionString := os.Getenv("PSQL_CONNECTION_STRING")
	fmt.Println(connectionString)
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// JWT config
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecretKey), nil
		},
		// When set, the middleware verifies that tokens are signed with the specific signing algorithm
		// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
		// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		SigningMethod: jwt.SigningMethodHS256,
	})

	utils.GetValidator()

	router := httprouter.New()
	routerConfig := &RouterConfig{
		DB:            db,
		JWTSecretKey:  jwtSecretKey,
		JWTMiddleware: jwtMiddleware,
	}
	SetupRoutes(router, routerConfig)

	log.Println("Starting server at port 8080")
	http.ListenAndServe(":8080", router)
}

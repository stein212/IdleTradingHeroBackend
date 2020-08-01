package handlers

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/IdleTradingHeroServer/models"
	"github.com/IdleTradingHeroServer/utils"
	viewmodels "github.com/IdleTradingHeroServer/viewModels"
	"github.com/julienschmidt/httprouter"
)

func GetUsers(db *sql.DB) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		users, err := models.Users().All(context.Background(), db)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}
		payload := utils.MapUsersToGetUserReponses(users, viewmodels.UserToGetUserResponse)

		respondJSON(w, payload)
	}
}

func ResetStrategiesStatus(db *sql.DB) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// update macd_strategies
		_, err := db.Exec("update macd_strategies set status = $1", "notDeployed")
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		// update mfi_strategies
		_, err = db.Exec("update mfi_strategies set status = $1", "notDeployed")
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		// update rsi_strategies
		_, err = db.Exec("update rsi_strategies set status = $1", "notDeployed")
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}
	}
}

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
		}
		payload := utils.MapUsersToGetUserReponses(users, viewmodels.UserToGetUserResponse)

		respondJSON(w, payload)
	}
}

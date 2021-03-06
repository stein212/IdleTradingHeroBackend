package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/IdleTradingHeroServer/models"
	routehelpers "github.com/IdleTradingHeroServer/routeHelpers"
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
		userResponses := utils.MapUsersToGetUserReponses(users, viewmodels.UserToGetUserResponse)

		jsonPayload, err := json.Marshal(userResponses)

		if err != nil {
			log.Println(err)
		}

		routehelpers.RespondJSON(w, jsonPayload)
	}
}

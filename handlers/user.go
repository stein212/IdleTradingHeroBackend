package handlers

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/IdleTradingHeroServer/models"
	viewmodels "github.com/IdleTradingHeroServer/viewModels"
	"github.com/julienschmidt/httprouter"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// var (
// 	logger = log.New(os.Stdout, "userController: ", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)
// )

type UserController struct {
	db *sql.DB
}

func NewUserController(config *ControllerConfig) *UserController {
	return &UserController{
		db: config.DB,
	}
}

func (uc *UserController) GetUserInfo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userID := getUserIDFromJWT(r)
	user, err := models.Users(qm.Where("id = ?", userID)).One(context.Background(), uc.db)

	if err != nil {
		logger.Println(err)
		w.WriteHeader(401)
		return
	}

	payload := viewmodels.UserToGetUserResponse(user)
	respondJSON(w, payload)
}

package handlers

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/IdleTradingHeroServer/models"
	routehelpers "github.com/IdleTradingHeroServer/routeHelpers"
	viewmodels "github.com/IdleTradingHeroServer/viewModels"
	"github.com/dgrijalva/jwt-go"
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
	j := r.Context().Value("user")
	accessToken := j.(*jwt.Token)
	userId := accessToken.Claims.(jwt.MapClaims)["userId"].(string)
	user, err := models.Users(qm.Where("id = ?", userId)).One(context.Background(), uc.db)

	if err != nil {
		logger.Println(err)
		w.WriteHeader(401)
		return
	}

	payload := viewmodels.UserToGetUserResponse(user)
	routehelpers.RespondJSON(w, payload)
}

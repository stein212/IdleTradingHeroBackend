package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/IdleTradingHeroServer/auth"
	"github.com/IdleTradingHeroServer/constants"
	"github.com/IdleTradingHeroServer/models"
	routehelpers "github.com/IdleTradingHeroServer/routeHelpers"
	"github.com/IdleTradingHeroServer/utils"
	viewmodels "github.com/IdleTradingHeroServer/viewModels"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

const (
	// epoch seconds
	jwtTokenDuration = 30 * 60
	jwtIssuer        = "IdleTradingHero"
)

var (
	logger       = log.New(os.Stdout, "authController: ", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)
	validate     = utils.GetValidator()
	enTranslator = utils.GetEnTranslator()
)

type AuthController struct {
	db           *sql.DB
	jwtSecretKey []byte
}

func (controller *AuthController) Init(db *sql.DB, jwtSecretKey string) {
	controller.db = db
	controller.jwtSecretKey = []byte(jwtSecretKey)
}

func (controller *AuthController) GetAccessTokenByPassword(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	loginUser := &viewmodels.LoginUser{}

	errPayload := routehelpers.DecodeJSONBody(w, r, loginUser)

	if errPayload != nil {
		logger.Println(errPayload)
		routehelpers.RespondWithErrorPayload(logger, w, errPayload)
		return
	}

	// check request
	if errs := validate.Struct(loginUser); errs != nil {
		validationErrs, _ := errs.(validator.ValidationErrors)

		errorResponses := make([]*routehelpers.ErrorResponse, len(validationErrs))

		for i, validationErr := range validationErrs {
			errorResponses[i] = &routehelpers.ErrorResponse{
				Code:    constants.ErrorCodeInvalidField,
				Message: validationErr.Translate(enTranslator),
			}
		}

		routehelpers.RespondWithErrorPayloadFromErrorResponses(logger, w, http.StatusUnprocessableEntity, errorResponses)
		return
	}

	username := strings.ToUpper(loginUser.Username)
	user, err := models.Users(qm.Where("username = ?", username)).One(context.Background(), controller.db)

	if err != nil {
		errMessage := fmt.Sprintf("Failed login attempt: Non existent user (%s)\n", user.Username)
		logger.Println(errMessage)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	logger.Println(user.Password, loginUser.Password)
	err = auth.CheckPassword(user, loginUser.Password)

	if err != nil {
		logger.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, err := controller.generateAccessToken(user)

	if err != nil {
		logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	payload := viewmodels.GetAccessToken{
		AccessToken: token,
	}

	jsonPayload, err := json.Marshal(payload)

	if err != nil {
		logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	routehelpers.RespondJSON(w, jsonPayload)
}

func (controller *AuthController) Register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	registerUser := &viewmodels.RegisterUser{}

	errPayload := routehelpers.DecodeJSONBody(w, r, registerUser)

	if errPayload != nil {
		logger.Println(errPayload)
		routehelpers.RespondWithErrorPayload(logger, w, errPayload)
		return
	}

	// check request
	if errs := validate.Struct(registerUser); errs != nil {
		validationErrs, _ := errs.(validator.ValidationErrors)

		errorResponses := make([]*routehelpers.ErrorResponse, len(validationErrs))

		for i, validationErr := range validationErrs {
			errorResponses[i] = &routehelpers.ErrorResponse{
				Code:    constants.ErrorCodeInvalidField,
				Message: validationErr.Translate(enTranslator),
			}
		}

		routehelpers.RespondWithErrorPayloadFromErrorResponses(logger, w, http.StatusUnprocessableEntity, errorResponses)
		return
	}

	// check if username is taken
	username := strings.ToUpper(registerUser.Username)
	isUsernameTaken, err := models.Users(qm.Where("username = ?", username)).Exists(context.Background(), controller.db)

	if err != nil {
		logger.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if isUsernameTaken {
		errMsg := fmt.Sprintf("Username Taken (%s)", registerUser.Username)
		logger.Println(errMsg)

		routehelpers.RespondWithErrorPayloadFromString(logger, w, http.StatusUnprocessableEntity, constants.ErrorCodeUsernameTaken, errMsg)
		return
	}

	// hash password
	passwordHash, _ := auth.HashPassword(registerUser.Password)

	// create user
	userID, _ := uuid.NewRandom()
	newUser := models.User{
		ID:        userID.String(),
		Username:  username,
		Password:  string(passwordHash),
		Email:     registerUser.Email,
		FirstName: registerUser.FirstName,
		LastName:  registerUser.LastName,
		CreatedOn: time.Now().UTC(),
	}

	err = newUser.Insert(context.Background(), controller.db, boil.Infer())
	if err != nil {
		logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: send email verification

	w.WriteHeader(http.StatusCreated)
}

type IdleTradingHeroClaims struct {
	UserID string `json:"userId"`
	jwt.StandardClaims
}

func (controller *AuthController) generateAccessToken(user *models.User) (string, error) {
	now := time.Now().Unix()

	claims := IdleTradingHeroClaims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: now + jwtTokenDuration,
			Issuer:    jwtIssuer,
			IssuedAt:  now,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(controller.jwtSecretKey)
}

package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/IdleTradingHeroServer/auth"
	"github.com/IdleTradingHeroServer/constants"
	"github.com/IdleTradingHeroServer/models"
	"github.com/IdleTradingHeroServer/utils"
	viewmodels "github.com/IdleTradingHeroServer/viewModels"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/securecookie"
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
	secureCookie *securecookie.SecureCookie
	domain       string
}

func NewAuthController(config *ControllerConfig) *AuthController {
	return &AuthController{
		db:           config.DB,
		jwtSecretKey: []byte(config.JWTSecretKey),
		secureCookie: config.SecureCookie,
		domain:       config.Domain,
	}
}

func (controller *AuthController) GetAccessTokenByPassword(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	user, err := controller.authByPassword(w, r)

	if err != nil {
		// response handled by authByPassword
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

	respondJSON(w, payload)
}

func (controller *AuthController) GetCookieAuthByPassword(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	user, err := controller.authByPassword(w, r)

	if err != nil {
		// response handled by authByPassword
		return
	}

	cookie, err := controller.generateCookieAuth(user)

	if err != nil {
		logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, cookie)

	w.WriteHeader(200)
}

// Returns a user if correct username and password. Will handle writing response if return err
func (controller *AuthController) authByPassword(w http.ResponseWriter, r *http.Request) (*models.User, error) {
	loginUser := &viewmodels.LoginUser{}

	errPayload := decodeJSONBody(w, r, loginUser)

	if errPayload != nil {
		logger.Println(errPayload)
		respondWithErrorPayload(logger, w, errPayload)
		return nil, errPayload
	}

	// check request
	if errs := validate.Struct(loginUser); errs != nil {
		validationErrs, _ := errs.(validator.ValidationErrors)

		errorResponses := make([]*errorResponse, len(validationErrs))

		for i, validationErr := range validationErrs {
			errorResponses[i] = &errorResponse{
				Code:    constants.ErrorCodeInvalidField,
				Message: validationErr.Translate(enTranslator),
			}
		}

		respondWithErrorPayloadFromErrorResponses(logger, w, http.StatusUnprocessableEntity, errorResponses)
		return nil, errs
	}

	username := strings.ToUpper(loginUser.Username)
	user, err := models.Users(qm.Where("username = ?", username)).One(context.Background(), controller.db)

	if err != nil {
		errMessage := fmt.Sprintf("Failed login attempt: Non existent user (%s)\n", user.Username)
		logger.Println(errMessage)
		w.WriteHeader(http.StatusUnauthorized)
		return nil, err
	}

	err = auth.CheckPassword(user, loginUser.Password)

	if err != nil {
		logger.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return nil, err
	}

	return user, nil
}

func (controller *AuthController) Register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	registerUser := &viewmodels.RegisterUser{}

	errPayload := decodeJSONBody(w, r, registerUser)

	if errPayload != nil {
		logger.Println(errPayload)
		respondWithErrorPayload(logger, w, errPayload)
		return
	}

	// check request
	isValidPayload := validateStruct(registerUser, w, logger)
	if !isValidPayload {
		// handled by validateStruct
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

		respondWithErrorPayloadFromString(logger, w, http.StatusUnprocessableEntity, constants.ErrorCodeUsernameTaken, errMsg)
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

func (controller *AuthController) generateCookieAuth(user *models.User) (*http.Cookie, error) {
	accessToken, err := controller.generateAccessToken(user)

	if err != nil {
		return nil, err
	}

	encoded, err := controller.secureCookie.Encode(constants.CookieAuthName, accessToken)

	if err != nil {
		return nil, err
	}

	cookie := &http.Cookie{
		Name:   constants.CookieAuthName,
		Value:  encoded,
		Domain: controller.domain,
		// Secure: true, // TODO: make development environment https
		HttpOnly: true,
	}

	return cookie, nil
}

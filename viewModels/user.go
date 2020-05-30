package viewmodels

import (
	"github.com/IdleTradingHeroServer/models"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type GetUserResponse struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	TelegramID *int   `json:"telegramId"`
	CreatedOn  int64  `json:"createdOn"`
}

func UserToGetUserResponse(user *models.User) *GetUserResponse {
	telegramId := &user.TelegramId.Int
	if !user.TelegramId.Valid {
		telegramId = nil
	}
	return &GetUserResponse{
		ID:         user.ID,
		Username:   user.Username,
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		TelegramID: telegramId,
		CreatedOn:  user.CreatedOn.Unix(),
	}
}

type LoginUser struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterUser struct {
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required,passwd"`
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

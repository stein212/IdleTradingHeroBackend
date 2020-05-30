package auth

import (
	"errors"
	"fmt"

	"github.com/IdleTradingHeroServer/models"
	"golang.org/x/crypto/bcrypt"
)

const (
	saltRounds = 13
)

func CheckPassword(user *models.User, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		errMessage := fmt.Sprintf("Invalid password (%s)", user.Username)
		return errors.New(errMessage)
	}

	return nil
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), saltRounds)
}

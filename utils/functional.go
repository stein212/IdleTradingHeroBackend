package utils

import (
	"github.com/IdleTradingHeroServer/models"
	viewmodels "github.com/IdleTradingHeroServer/viewModels"
)

func MapUsersToGetUserReponses(users models.UserSlice, f func(*models.User) *viewmodels.GetUserResponse) []*viewmodels.GetUserResponse {
	result := make([]*viewmodels.GetUserResponse, len(users))
	for i, user := range users {
		result[i] = f(user)
	}

	return result
}

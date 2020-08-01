package viewmodels

import "github.com/IdleTradingHeroServer/models"

type CreateStrategyEvent struct {
	StrategyID string `json:"strategyId" validate:"required"`
	Action     string `json:"action" validate:"required"`
	Amount     int    `json:"amount" validate:"required"`
	EventOn    int64  `json:"eventOn" validate:"required"`
}

type StrategyEventResponse struct {
	ID         string `json:"id"`
	StrategyID string `json:"strategyId"`
	Action     string `json:"action"`
	Amount     int    `json:"amount"`
	EventOn    int64  `json:"eventOn"`
	CreatedOn  int64  `json:"createdOn"`
}

func StrategyEventToStrategyEventResponse(strategyEvent *models.StrategyEvent) *StrategyEventResponse {
	return &StrategyEventResponse{
		ID:         strategyEvent.ID,
		StrategyID: strategyEvent.StrategyID,
		Action:     strategyEvent.StrategyAction,
		Amount:     strategyEvent.Amount,
		EventOn:    strategyEvent.EventOn.Unix(),
		CreatedOn:  strategyEvent.CreatedOn.Unix(),
	}
}

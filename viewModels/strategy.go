package viewmodels

import (
	"github.com/IdleTradingHeroServer/constants"
	"github.com/IdleTradingHeroServer/models"
)

type InitialiseStrategyConfig struct {
	StrategyType string `json:"strategyType" validate:"required"`
	Capital      int    `json:"capital" validate:"required"`
}

type MacdParams struct {
	Ema26 int
	Ema12 int
	Ema9  int
}

type BacktestMacdConfig struct {
	Asset    string  `json:"asset" validate:"required"`
	Strategy string  `json:"strategy" validate:"required"`
	Capital  float64 `json:"capital" validate:"required"`
	Ema26    int     `json:"ema26" validate:"required"`
	Ema12    int     `json:"ema12" validate:"required"`
	Ema9     int     `json:"ema9" validate:"required"`
}

type CreateMacdConfig struct {
	Name        string `json:"name" validate:"required"`
	Instrument  string `json:"instrument" validate:"required"`
	Granularity string `json:"granularity" validate:"required"`
	Ema26       int    `json:"ema26" validate:"required"`
	Ema12       int    `json:"ema12" validate:"required"`
	Ema9        int    `json:"ema9" validate:"required"`
}

type CreateMacdResponse struct {
	ID string `json:"id"`
}

type StrategyResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type MacdStrategyResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Instrument   string `json:"instrument"`
	Granularity  string `json:"granularity"`
	Ema26        int    `json:"ema26"`
	Ema12        int    `json:"ema12"`
	Ema9         int    `json:"ema9"`
	Status       string `json:"status"`
	CreatedOn    int64  `json:"createdOn"`
	LastEditedOn int64  `json:"lastEditedOn"`
}

func MacdStrategyToStrategyResponse(strategy *models.MacdStrategy) *StrategyResponse {
	return &StrategyResponse{
		ID:   strategy.ID,
		Name: strategy.Name,
		Type: constants.StrategyTypeMacd,
	}
}

func MacdStrategyToMacdStrategyResponse(strategy *models.MacdStrategy) *MacdStrategyResponse {
	return &MacdStrategyResponse{
		ID:           strategy.ID,
		Name:         strategy.Name,
		Instrument:   strategy.Instrument,
		Granularity:  strategy.Granularity,
		Ema26:        strategy.Ema26,
		Ema12:        strategy.Ema12,
		Ema9:         strategy.Ema9,
		Status:       strategy.Status,
		CreatedOn:    strategy.CreatedOn.Unix(),
		LastEditedOn: strategy.LastEditedOn.Unix(),
	}
}

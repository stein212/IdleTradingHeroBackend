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

type CreateMacdConfig struct {
	Name        string `json:"name" validate:"required"`
	Instrument  string `json:"instrument" validate:"required"`
	Granularity string `json:"granularity" validate:"required"`
	Ema26       int    `json:"ema26" validate:"required"`
	Ema12       int    `json:"ema12" validate:"required"`
	Ema9        int    `json:"ema9" validate:"required"`
}

type CreateMfiConfig struct {
	Name            string `json:"name" validate:"required"`
	Instrument      string `json:"instrument" validate:"required"`
	Granularity     string `json:"granularity" validate:"required"`
	OverboughtLevel int    `json:"overboughtLevel" validate:"required"`
	OversoldLevel   int    `json:"oversoldLevel" validate:"required"`
}

type CreateRsiConfig struct {
	Name            string `json:"name" validate:"required"`
	Instrument      string `json:"instrument" validate:"required"`
	Granularity     string `json:"granularity" validate:"required"`
	OverboughtLevel int    `json:"overboughtLevel" validate:"required"`
	OversoldLevel   int    `json:"oversoldLevel" validate:"required"`
}

type CreateStrategyResponse struct {
	ID string `json:"id"`
}

type StrategyResponse struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Type    string  `json:"type"`
	Status  string  `json:"status"`
	Balance float64 `json:"balance"`
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

type MfiStrategyResponse struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Instrument      string `json:"instrument"`
	Granularity     string `json:"granularity"`
	OverboughtLevel int    `json:"overboughtLevel"`
	OversoldLevel   int    `json:"oversoldLevel"`
	Status          string `json:"status"`
	CreatedOn       int64  `json:"createdOn"`
	LastEditedOn    int64  `json:"lastEditedOn"`
}

type RsiStrategyResponse struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Instrument      string `json:"instrument"`
	Granularity     string `json:"granularity"`
	OverboughtLevel int    `json:"overboughtLevel"`
	OversoldLevel   int    `json:"oversoldLevel"`
	Status          string `json:"status"`
	CreatedOn       int64  `json:"createdOn"`
	LastEditedOn    int64  `json:"lastEditedOn"`
}

func MacdStrategyToStrategyResponse(strategy *models.MacdStrategy) *StrategyResponse {
	return &StrategyResponse{
		ID:     strategy.ID,
		Name:   strategy.Name,
		Type:   constants.StrategyTypeMacd,
		Status: strategy.Status,
	}
}

func MfiStrategyToStrategyResponse(strategy *models.MfiStrategy) *StrategyResponse {
	return &StrategyResponse{
		ID:     strategy.ID,
		Name:   strategy.Name,
		Type:   constants.StrategyTypeMfi,
		Status: strategy.Status,
	}
}

func RsiStrategyToStrategyResponse(strategy *models.RsiStrategy) *StrategyResponse {
	return &StrategyResponse{
		ID:     strategy.ID,
		Name:   strategy.Name,
		Type:   constants.StrategyTypeRsi,
		Status: strategy.Status,
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

func MfiStrategyToMfiStrategyResponse(strategy *models.MfiStrategy) *MfiStrategyResponse {
	return &MfiStrategyResponse{
		ID:              strategy.ID,
		Name:            strategy.Name,
		Instrument:      strategy.Instrument,
		Granularity:     strategy.Granularity,
		OverboughtLevel: strategy.OverboughtLevel,
		OversoldLevel:   strategy.OversoldLevel,
		Status:          strategy.Status,
		CreatedOn:       strategy.CreatedOn.Unix(),
		LastEditedOn:    strategy.LastEditedOn.Unix(),
	}
}

func RsiStrategyToRsiStrategyResponse(strategy *models.RsiStrategy) *RsiStrategyResponse {
	return &RsiStrategyResponse{
		ID:              strategy.ID,
		Name:            strategy.Name,
		Instrument:      strategy.Instrument,
		Granularity:     strategy.Granularity,
		OverboughtLevel: strategy.OverboughtLevel,
		OversoldLevel:   strategy.OversoldLevel,
		Status:          strategy.Status,
		CreatedOn:       strategy.CreatedOn.Unix(),
		LastEditedOn:    strategy.LastEditedOn.Unix(),
	}
}

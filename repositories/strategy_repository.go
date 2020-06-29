package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/IdleTradingHeroServer/constants"
	"github.com/IdleTradingHeroServer/models"
	strategy_proto "github.com/IdleTradingHeroServer/pb/strategy"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/boil"
	. "github.com/volatiletech/sqlboiler/queries/qm"
)

type StrategyRepository interface {
	CreateMacdStrategy(ctx context.Context, config *MacdConfig, userID string) (*models.MacdStrategy, error)

	GetMacdStrategies(ctx context.Context, userID string) ([]*models.MacdStrategy, error)
	GetMacdStrategy(ctx context.Context, userID string, macdStrategyID string) (*models.MacdStrategy, error)

	// StrategyService
	InitialiseStrategy(ctx context.Context, strategyID string, strategyType string, capital int) error
	StartStrategy(ctx context.Context, strategyID string) error
	PauseStrategy(ctx context.Context, strategyID string) error
	ActStrategy(ctx context.Context, strategyID string) error
	GetStrategyStatistics(ctx context.Context, strategyID string) (*strategy_proto.Statistics, error)
	GetStrategyData(ctx context.Context, strategyID string, length int) (map[string][]float32, error)
}

type MacdConfig struct {
	Name        string
	Instrument  string
	Granularity string
	Ema26       int
	Ema12       int
	Ema9        int
}

type strategyRepository struct {
	db             *sql.DB
	strategyClient strategy_proto.StrategyServiceClient
}

func NewStrategyRepository(db *sql.DB, strategyClient strategy_proto.StrategyServiceClient) *strategyRepository {
	return &strategyRepository{
		db:             db,
		strategyClient: strategyClient,
	}
}

func (r *strategyRepository) CreateMacdStrategy(ctx context.Context, config *MacdConfig, userID string) (*models.MacdStrategy, error) {
	macdStrategyID, _ := uuid.NewRandom()
	newMacdStrategy := &models.MacdStrategy{
		ID:           macdStrategyID.String(),
		UserID:       userID,
		Name:         config.Name,
		Instrument:   config.Instrument,
		Granularity:  config.Granularity,
		Ema26:        config.Ema26,
		Ema12:        config.Ema12,
		Ema9:         config.Ema9,
		Status:       constants.StrategyNotDeployed,
		CreatedOn:    time.Now().UTC(),
		LastEditedOn: time.Now().UTC(),
	}

	err := newMacdStrategy.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return newMacdStrategy, nil
}

func (r *strategyRepository) GetMacdStrategies(ctx context.Context, userID string) ([]*models.MacdStrategy, error) {
	return models.MacdStrategies(
		Where("user_id = ?", userID),
	).All(ctx, r.db)
}

func (r *strategyRepository) GetMacdStrategy(ctx context.Context, userID string, macdStrategyID string) (*models.MacdStrategy, error) {
	return models.MacdStrategies(
		Where("user_id = ?", userID),
		And("id = ?", macdStrategyID),
	).One(ctx, r.db)
}

func (r *strategyRepository) InitialiseStrategy(ctx context.Context, strategyID string, strategyType string, capital int) error {
	var instrument string
	var granularity string
	var params *structpb.Struct

	var macdStrategy *models.MacdStrategy
	var err error

	switch strategyType {
	case constants.StrategyTypeMacd:
		macdStrategy, err = models.MacdStrategies(Where("id = ?", strategyID)).One(ctx, r.db)
		if err != nil {
			return err
		}

		instrument = macdStrategy.Instrument
		granularity = macdStrategy.Granularity

		params = &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"ema26": {
					Kind: &structpb.Value_NumberValue{
						NumberValue: float64(macdStrategy.Ema26),
					},
				},
				"ema12": {
					Kind: &structpb.Value_NumberValue{
						NumberValue: float64(macdStrategy.Ema12),
					},
				},
				"ema9": {
					Kind: &structpb.Value_NumberValue{
						NumberValue: float64(macdStrategy.Ema9),
					},
				},
			},
		}
	}

	_, err = r.strategyClient.InitialiseAlgorithm(ctx, &strategy_proto.Selection{
		ID:          strategyID,
		Instrument:  instrument,
		Granularity: granularity,
		Capital:     int32(capital),
		Strategy:    strategyType,
		Parameters:  params,
	})

	if err != nil {
		return err
	}

	switch strategyType {
	case constants.StrategyTypeMacd:
		macdStrategy.Status = constants.StrategyDeployed
		macdStrategy.Update(ctx, r.db, boil.Infer())
	}

	return nil
}

func (r *strategyRepository) StartStrategy(ctx context.Context, strategyID string) error {
	startRes, err := r.strategyClient.StartAlgorithm(ctx, &strategy_proto.StartAlgorithmParam{
		ID: strategyID,
	})

	if err != nil {
		return err
	}

	if startRes.Error != "" {
		return errors.New(startRes.Error)
	}

	return nil
}

func (r *strategyRepository) PauseStrategy(ctx context.Context, strategyID string) error {
	stopRes, err := r.strategyClient.StopAlgorithm(ctx, &strategy_proto.StopAlgorithmParam{
		ID: strategyID,
	})

	if err != nil {
		return err
	}

	if stopRes.Error != "" {
		return errors.New(stopRes.Error)
	}

	return nil
}

func (r *strategyRepository) ActStrategy(ctx context.Context, strategyID string) error {
	_, err := r.strategyClient.Act(ctx, &strategy_proto.Tmp{
		ID:  strategyID,
		Tmp: 1,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *strategyRepository) GetStrategyStatistics(ctx context.Context, strategyID string) (*strategy_proto.Statistics, error) {
	statistics, err := r.strategyClient.GetStatistics(ctx, &strategy_proto.Tmp{
		ID:  strategyID,
		Tmp: 1,
	})

	if err != nil {
		return nil, err
	}

	return statistics, nil
}

func (r *strategyRepository) GetStrategyData(ctx context.Context, strategyID string, length int) (map[string][]float32, error) {
	history, err := r.strategyClient.GetData(ctx, &strategy_proto.HistoryParams{
		ID:     strategyID,
		Length: int32(length),
	})

	if err != nil {
		return nil, err
	}

	data := make(map[string][]float32)

	for _, ts := range history.GetTS() {
		data[ts.Key] = ts.Value
	}

	return data, nil
}

package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/IdleTradingHeroServer/constants"
	"github.com/IdleTradingHeroServer/models"
	strategy_proto "github.com/IdleTradingHeroServer/pb/strategy"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/google/uuid"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	. "github.com/volatiletech/sqlboiler/queries/qm"
)

type StrategyRepository interface {
	CreateMacdStrategy(ctx context.Context, config *MacdConfig, userID string) (*models.MacdStrategy, error)
	CreateMfiStrategy(ctx context.Context, config *MfiConfig, userID string) (*models.MfiStrategy, error)
	CreateRsiStrategy(ctx context.Context, config *RsiConfig, userID string) (*models.RsiStrategy, error)

	GetMacdStrategies(ctx context.Context, userID string) ([]*models.MacdStrategy, error)
	GetMacdStrategy(ctx context.Context, userID string, macdStrategyID string) (*models.MacdStrategy, error)
	GetMfiStrategies(ctx context.Context, userID string) ([]*models.MfiStrategy, error)
	GetMfiStrategy(ctx context.Context, userID string, mfiStrategyID string) (*models.MfiStrategy, error)
	GetRsiStrategies(ctx context.Context, userID string) ([]*models.RsiStrategy, error)
	GetRsiStrategy(ctx context.Context, userID string, rsiStrategyID string) (*models.RsiStrategy, error)

	GetBalance(ctx context.Context, strategyID string) (float64, error)

	UpdateStrategyStatus(ctx context.Context, strategyID string, strategyType string, status string) error

	// StrategyService
	InitialiseStrategy(ctx context.Context, strategyID string, strategyType string, capital int) error
	StartStrategy(ctx context.Context, strategyID string, strategyType string) error
	PauseStrategy(ctx context.Context, strategyID string, strategyType string) error
	ActStrategy(ctx context.Context, strategyID string) error
	GetStrategyStatistics(ctx context.Context, strategyID string) (*strategy_proto.Statistics, error)
	GetStrategyData(ctx context.Context, strategyID string, length int) (map[string][]float64, error)
	GetStrategyIndicatorData(ctx context.Context, strategyID string, length int) (map[string][]float64, error)
	GetStrategyPerformanceData(ctx context.Context, strategyID string, length int) (map[string][]float64, error)

	DeleteStrategy(ctx context.Context, strategyID string, strategyType string) error
}

type MacdConfig struct {
	Name        string
	Instrument  string
	Granularity string
	Ema26       int
	Ema12       int
	Ema9        int
}

type MfiConfig struct {
	Name            string
	Instrument      string
	Granularity     string
	OverboughtLevel int
	OversoldLevel   int
}

type RsiConfig struct {
	Name            string
	Instrument      string
	Granularity     string
	OverboughtLevel int
	OversoldLevel   int
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

func (r *strategyRepository) CreateMfiStrategy(ctx context.Context, config *MfiConfig, userID string) (*models.MfiStrategy, error) {
	mfiStrategyID, _ := uuid.NewRandom()
	newMfiStrategy := &models.MfiStrategy{
		ID:              mfiStrategyID.String(),
		UserID:          userID,
		Name:            config.Name,
		Instrument:      config.Instrument,
		Granularity:     config.Granularity,
		OverboughtLevel: config.OverboughtLevel,
		OversoldLevel:   config.OversoldLevel,
		Status:          constants.StrategyNotDeployed,
		CreatedOn:       time.Now().UTC(),
		LastEditedOn:    time.Now().UTC(),
	}

	err := newMfiStrategy.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return newMfiStrategy, nil
}

func (r *strategyRepository) CreateRsiStrategy(ctx context.Context, config *RsiConfig, userID string) (*models.RsiStrategy, error) {
	rsiStrategyID, _ := uuid.NewRandom()
	newRsiStrategy := &models.RsiStrategy{
		ID:              rsiStrategyID.String(),
		UserID:          userID,
		Name:            config.Name,
		Instrument:      config.Instrument,
		Granularity:     config.Granularity,
		OverboughtLevel: config.OverboughtLevel,
		OversoldLevel:   config.OversoldLevel,
		Status:          constants.StrategyNotDeployed,
		CreatedOn:       time.Now().UTC(),
		LastEditedOn:    time.Now().UTC(),
	}

	err := newRsiStrategy.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return newRsiStrategy, nil
}

func (r *strategyRepository) GetMacdStrategies(ctx context.Context, userID string) ([]*models.MacdStrategy, error) {
	return models.MacdStrategies(
		Where("user_id = ?", userID),
		And("deleted_on is null"),
	).All(ctx, r.db)
}

func (r *strategyRepository) GetMacdStrategy(ctx context.Context, userID string, macdStrategyID string) (*models.MacdStrategy, error) {
	return models.MacdStrategies(
		Where("user_id = ?", userID),
		And("id = ?", macdStrategyID),
		And("deleted_on is null"),
	).One(ctx, r.db)
}

func (r *strategyRepository) GetMfiStrategies(ctx context.Context, userID string) ([]*models.MfiStrategy, error) {
	return models.MfiStrategies(
		Where("user_id = ?", userID),
		And("deleted_on is null"),
	).All(ctx, r.db)
}

func (r *strategyRepository) GetMfiStrategy(ctx context.Context, userID string, MfiStrategyID string) (*models.MfiStrategy, error) {
	return models.MfiStrategies(
		Where("user_id = ?", userID),
		And("id = ?", MfiStrategyID),
		And("deleted_on is null"),
	).One(ctx, r.db)
}

func (r *strategyRepository) GetRsiStrategies(ctx context.Context, userID string) ([]*models.RsiStrategy, error) {
	return models.RsiStrategies(
		Where("user_id = ?", userID),
		And("deleted_on is null"),
	).All(ctx, r.db)
}

func (r *strategyRepository) GetRsiStrategy(ctx context.Context, userID string, RsiStrategyID string) (*models.RsiStrategy, error) {
	return models.RsiStrategies(
		Where("user_id = ?", userID),
		And("id = ?", RsiStrategyID),
		And("deleted_on is null"),
	).One(ctx, r.db)
}

func (r *strategyRepository) GetBalance(ctx context.Context, strategyID string) (float64, error) {
	res, err := r.strategyClient.GetBalance(ctx, &strategy_proto.GetBalanceParam{
		ID: strategyID,
	})

	if err != nil {
		return 0, err
	}

	return res.Balance, nil
}

func (r *strategyRepository) UpdateStrategyStatus(ctx context.Context, strategyID string, strategyType string, status string) error {
	switch strategyType {
	case constants.StrategyTypeMacd:
		macdStrategy, err := models.MacdStrategies(Where("id = ?", strategyID)).One(ctx, r.db)
		if err != nil {
			return err
		}

		macdStrategy.Status = status
		macdStrategy.LastEditedOn = time.Now().UTC()
		macdStrategy.Update(ctx, r.db, boil.Infer())

		return nil

	case constants.StrategyTypeMfi:
		mfiStrategy, err := models.MfiStrategies(Where("id = ?", strategyID)).One(ctx, r.db)
		if err != nil {
			return err
		}

		mfiStrategy.Status = status
		mfiStrategy.LastEditedOn = time.Now().UTC()
		mfiStrategy.Update(ctx, r.db, boil.Infer())

		return nil
	case constants.StrategyTypeRsi:
		rsiStrategy, err := models.RsiStrategies(Where("id = ?", strategyID)).One(ctx, r.db)
		if err != nil {
			return err
		}

		rsiStrategy.Status = status
		rsiStrategy.LastEditedOn = time.Now().UTC()
		rsiStrategy.Update(ctx, r.db, boil.Infer())

		return nil
	default:
		return fmt.Errorf("Unsupported strategyType: %s", strategyType)
	}
}

func (r *strategyRepository) InitialiseStrategy(ctx context.Context, strategyID string, strategyType string, capital int) error {
	var instrument string
	var granularity string
	var params *structpb.Struct

	var macdStrategy *models.MacdStrategy
	var mfiStrategy *models.MfiStrategy
	var rsiStrategy *models.RsiStrategy
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
	case constants.StrategyTypeMfi:
		mfiStrategy, err = models.MfiStrategies(Where("id = ?", strategyID)).One(ctx, r.db)
		if err != nil {
			return err
		}

		instrument = mfiStrategy.Instrument
		granularity = mfiStrategy.Granularity

		params = &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"mfi80": {
					Kind: &structpb.Value_NumberValue{
						NumberValue: float64(mfiStrategy.OverboughtLevel),
					},
				},
				"mfi20": {
					Kind: &structpb.Value_NumberValue{
						NumberValue: float64(mfiStrategy.OversoldLevel),
					},
				},
			},
		}
	case constants.StrategyTypeRsi:
		rsiStrategy, err = models.RsiStrategies(Where("id = ?", strategyID)).One(ctx, r.db)
		if err != nil {
			return err
		}

		instrument = rsiStrategy.Instrument
		granularity = rsiStrategy.Granularity

		params = &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"rsi70": {
					Kind: &structpb.Value_NumberValue{
						NumberValue: float64(rsiStrategy.OverboughtLevel),
					},
				},
				"rsi30": {
					Kind: &structpb.Value_NumberValue{
						NumberValue: float64(rsiStrategy.OversoldLevel),
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
	case constants.StrategyTypeMfi:
		mfiStrategy.Status = constants.StrategyDeployed
		mfiStrategy.Update(ctx, r.db, boil.Infer())
	case constants.StrategyTypeRsi:
		rsiStrategy.Status = constants.StrategyDeployed
		rsiStrategy.Update(ctx, r.db, boil.Infer())
	}

	return nil
}

func (r *strategyRepository) StartStrategy(ctx context.Context, strategyID string, strategyType string) error {
	startRes, err := r.strategyClient.StartAlgorithm(ctx, &strategy_proto.StartAlgorithmParam{
		ID: strategyID,
	})

	if err != nil {
		return err
	}

	if startRes.Error != "" {
		return errors.New(startRes.Error)
	}

	err = r.UpdateStrategyStatus(ctx, strategyID, strategyType, constants.StrategyLive)
	if err != nil {
		return err
	}

	return nil
}

func (r *strategyRepository) PauseStrategy(ctx context.Context, strategyID string, strategyType string) error {
	stopRes, err := r.strategyClient.StopAlgorithm(ctx, &strategy_proto.StopAlgorithmParam{
		ID: strategyID,
	})

	if err != nil {
		return err
	}

	if stopRes.Error != "" {
		return errors.New(stopRes.Error)
	}

	err = r.UpdateStrategyStatus(ctx, strategyID, strategyType, constants.StrategyPaused)
	if err != nil {
		return err
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

func (r *strategyRepository) GetStrategyData(ctx context.Context, strategyID string, length int) (map[string][]float64, error) {
	history, err := r.strategyClient.GetData(ctx, &strategy_proto.HistoryParams{
		ID:     strategyID,
		Length: int32(length),
	})

	if err != nil {
		return nil, err
	}

	data := make(map[string][]float64)

	for _, ts := range history.GetTS() {
		data[ts.Key] = ts.Value
	}

	return data, nil
}

func (r *strategyRepository) GetStrategyIndicatorData(ctx context.Context, strategyID string, length int) (map[string][]float64, error) {
	history, err := r.strategyClient.GetIndicators(ctx, &strategy_proto.HistoryParams{
		ID:     strategyID,
		Length: int32(length),
	})

	if err != nil {
		return nil, err
	}

	data := make(map[string][]float64)

	for _, ts := range history.GetTS() {
		data[ts.Key] = ts.Value
	}

	return data, nil
}

func (r *strategyRepository) GetStrategyPerformanceData(ctx context.Context, strategyID string, length int) (map[string][]float64, error) {
	history, err := r.strategyClient.GetPerformances(ctx, &strategy_proto.HistoryParams{
		ID:     strategyID,
		Length: int32(length),
	})

	if err != nil {
		return nil, err
	}

	data := make(map[string][]float64)

	for _, ts := range history.GetTS() {
		data[ts.Key] = ts.Value
	}

	return data, nil
}

func (r *strategyRepository) DeleteStrategy(ctx context.Context, strategyID string, strategyType string) error {
	// attempt to stop if there is a deployed strategy
	_, _ = r.strategyClient.StopAlgorithm(ctx, &strategy_proto.StopAlgorithmParam{
		ID: strategyID,
	})

	// if err != nil {
	// 	//return err
	// }

	// if stopRes.Error != "" {
	// 	return errors.New(stopRes.Error)
	// }

	switch strategyType {
	case constants.StrategyTypeMacd:
		macdStrategy, err := models.MacdStrategies(Where("id = ?", strategyID)).One(ctx, r.db)
		if err != nil {
			return err
		}

		macdStrategy.Status = constants.StrategyDeleted
		macdStrategy.DeletedOn = null.TimeFrom(time.Now().UTC())
		macdStrategy.LastEditedOn = time.Now().UTC()
		macdStrategy.Update(ctx, r.db, boil.Infer())

		return nil

	case constants.StrategyTypeMfi:
		mfiStrategy, err := models.MfiStrategies(Where("id = ?", strategyID)).One(ctx, r.db)
		if err != nil {
			return err
		}

		mfiStrategy.Status = constants.StrategyDeleted
		mfiStrategy.DeletedOn = null.TimeFrom(time.Now().UTC())
		mfiStrategy.LastEditedOn = time.Now().UTC()
		mfiStrategy.Update(ctx, r.db, boil.Infer())

		return nil
	case constants.StrategyTypeRsi:
		rsiStrategy, err := models.RsiStrategies(Where("id = ?", strategyID)).One(ctx, r.db)
		if err != nil {
			return err
		}

		rsiStrategy.Status = constants.StrategyDeleted
		rsiStrategy.DeletedOn = null.TimeFrom(time.Now().UTC())
		rsiStrategy.LastEditedOn = time.Now().UTC()
		rsiStrategy.Update(ctx, r.db, boil.Infer())

		return nil
	default:
		return fmt.Errorf("Unsupported strategyType: %s", strategyType)
	}
}

package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/IdleTradingHeroServer/models"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/boil"
	. "github.com/volatiletech/sqlboiler/queries/qm"
)

type StrategyEventRepository interface {
	CreateStrategyEvent(ctx context.Context, strategyID string, action string, amount int, eventOn time.Time) error

	GetStrategyEvents(ctx context.Context, strategyID string, length int, to time.Time) ([]*models.StrategyEvent, error)
	GetUserStrategiesEvents(ctx context.Context, userID string, length int, to time.Time) ([]*models.StrategyEvent, error)
}

type strategyEventRepository struct {
	db *sql.DB
}

func NewStrategyEventRepository(db *sql.DB) *strategyEventRepository {
	return &strategyEventRepository{
		db: db,
	}
}

func (r *strategyEventRepository) CreateStrategyEvent(
	ctx context.Context,
	strategyID string,
	action string,
	amount int,
	eventOn time.Time,
) error {
	id, _ := uuid.NewRandom()
	event := models.StrategyEvent{
		ID:             id.String(),
		StrategyID:     strategyID,
		StrategyAction: action,
		Amount:         amount,
		EventOn:        eventOn,
		CreatedOn:      time.Now().UTC(),
	}

	fmt.Println(event)
	err := event.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

func (r *strategyEventRepository) GetStrategyEvents(ctx context.Context, strategyID string, length int, to time.Time) ([]*models.StrategyEvent, error) {
	return models.StrategyEvents(Where("strategy_id = ?", strategyID), And("event_on <= ?", to), OrderBy("event_on desc"), Limit(length)).All(ctx, r.db)
}

func (r *strategyEventRepository) GetUserStrategiesEvents(ctx context.Context, userID string, length int, to time.Time) ([]*models.StrategyEvent, error) {
	return models.StrategyEvents(
		LeftOuterJoin("macd_strategies macds on strategy_id = macds.id"),
		LeftOuterJoin("mfi_strategies mfis on strategy_id = mfis.id"),
		LeftOuterJoin("rsi_strategies rsis on strategy_id = rsis.id"),
		Where("macds.user_id = ?", userID),
		And("macds.deleted_on is null"),
		Or("mfis.user_id = ?", userID),
		And("mfis.deleted_on is null"),
		Or("rsis.user_id = ?", userID),
		And("rsis.deleted_on is null"),
		OrderBy("event_on desc"),
		Limit(length),
	).All(ctx, r.db)
}

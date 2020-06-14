package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	idletradinghero "github.com/IdleTradingHeroServer/route/github.com/idletradinghero/v2"
	routehelpers "github.com/IdleTradingHeroServer/routeHelpers"
	viewmodels "github.com/IdleTradingHeroServer/viewModels"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/protobuf/types/known/structpb"
)

var (
	strategyLogger = log.New(os.Stdout, "strategyController: ", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)
)

type StrategyController struct {
	pythonServiceClient idletradinghero.RouteClient
}

func NewStrategyController(config *ControllerConfig) *StrategyController {
	return &StrategyController{
		pythonServiceClient: config.PythonServiceClient,
	}
}

func (c *StrategyController) GetStrategyPerformance(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	postMACDConfig := &viewmodels.PostMACDConfig{}

	errPayload := routehelpers.DecodeJSONBody(w, r, postMACDConfig)

	if errPayload != nil {
		strategyLogger.Println(errPayload)
		routehelpers.RespondWithErrorPayload(logger, w, errPayload)
		return
	}

	intCapital := int32(postMACDConfig.Capital)

	selection := idletradinghero.Selection{
		Asset:    postMACDConfig.Asset,
		Strategy: postMACDConfig.Strategy,
		Capital:  intCapital,
		Parameters: &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"ema26": {
					Kind: &structpb.Value_NumberValue{
						NumberValue: float64(postMACDConfig.Ema26),
					},
				},
				"ema12": {
					Kind: &structpb.Value_NumberValue{
						NumberValue: float64(postMACDConfig.Ema12),
					},
				},
				"ema9": {
					Kind: &structpb.Value_NumberValue{
						NumberValue: float64(postMACDConfig.Ema9),
					},
				},
			},
		},
	}

	statistics, err := c.pythonServiceClient.InitialiseAlgorithm(context.Background(), &selection)

	if err != nil {
		strategyLogger.Println(err)
		routehelpers.RespondWithErrorPayloadFromString(logger, w, http.StatusInternalServerError, "unknown", err.Error())
		return
	}

	json.NewEncoder(w).Encode(statistics)
}

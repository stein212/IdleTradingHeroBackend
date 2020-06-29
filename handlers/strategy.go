package handlers

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/IdleTradingHeroServer/repositories"
	viewmodels "github.com/IdleTradingHeroServer/viewModels"
	"github.com/julienschmidt/httprouter"
)

var (
	strategyLogger = log.New(os.Stdout, "strategyController: ", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)
)

type StrategyController struct {
	strategyRepository repositories.StrategyRepository
}

func NewStrategyController(config *ControllerConfig) *StrategyController {
	return &StrategyController{
		strategyRepository: config.StrategyRepository,
	}
}

func (c *StrategyController) CreateMacdStrategy(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userID := getUserIDFromJWT(r)

	createMACDConfig := &viewmodels.CreateMacdConfig{}

	errPayload := decodeJSONBody(w, r, createMACDConfig)

	if errPayload != nil {
		strategyLogger.Println(errPayload)
		respondWithErrorPayload(logger, w, errPayload)
		return
	}

	// check request
	isValidPayload := validateStruct(createMACDConfig, w, logger)
	if !isValidPayload {
		// handled by validateStruct
		return
	}

	macdConfig := &repositories.MacdConfig{
		Name:        createMACDConfig.Name,
		Instrument:  createMACDConfig.Instrument,
		Granularity: createMACDConfig.Granularity,
		Ema26:       createMACDConfig.Ema26,
		Ema12:       createMACDConfig.Ema12,
		Ema9:        createMACDConfig.Ema9,
	}

	macdStrategy, err := c.strategyRepository.CreateMacdStrategy(r.Context(), macdConfig, userID)
	if err != nil {
		strategyLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	respondJSON(w, &viewmodels.CreateMacdResponse{
		ID: macdStrategy.ID,
	})
}

func (c *StrategyController) GetStrategies(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userID := getUserIDFromJWT(r)

	macdStrategies, err := c.strategyRepository.GetMacdStrategies(r.Context(), userID)
	if err != nil {
		strategyLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	payload := make([]*viewmodels.StrategyResponse, len(macdStrategies))
	for i, macdStrategy := range macdStrategies {
		payload[i] = viewmodels.MacdStrategyToStrategyResponse(macdStrategy)
	}

	respondJSON(w, payload)
}

func (c *StrategyController) GetMacdStrategy(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userID := getUserIDFromJWT(r)
	strategyID := params.ByName("strategyID")

	macdStrategy, err := c.strategyRepository.GetMacdStrategy(r.Context(), userID, strategyID)
	if err != nil {
		strategyLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	payload := viewmodels.MacdStrategyToMacdStrategyResponse(macdStrategy)

	respondJSON(w, payload)
}

func (c *StrategyController) InitialiseStrategy(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	strategyID := params.ByName("strategyID")

	initialiseStrategyConfig := &viewmodels.InitialiseStrategyConfig{}

	errPayload := decodeJSONBody(w, r, initialiseStrategyConfig)

	if errPayload != nil {
		strategyLogger.Println(errPayload)
		respondWithErrorPayload(logger, w, errPayload)
		return
	}

	// check request
	isValidPayload := validateStruct(initialiseStrategyConfig, w, logger)
	if !isValidPayload {
		// handled by validateStruct
		return
	}

	err := c.strategyRepository.InitialiseStrategy(
		r.Context(),
		strategyID,
		initialiseStrategyConfig.StrategyType,
		initialiseStrategyConfig.Capital)

	if err != nil {
		strategyLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// call act once to really make it ready
	err = c.strategyRepository.ActStrategy(r.Context(), strategyID)
	if err != nil {
		strategyLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (c *StrategyController) StartStrategy(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	strategyID := params.ByName("strategyID")

	err := c.strategyRepository.StartStrategy(r.Context(), strategyID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (c *StrategyController) PauseStrategy(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	strategyID := params.ByName("strategyID")

	err := c.strategyRepository.PauseStrategy(r.Context(), strategyID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (c *StrategyController) GetStrategyData(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	strategyID := params.ByName("strategyID")
	dataLength, err := strconv.ParseInt(params.ByName("length"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	data, err := c.strategyRepository.GetStrategyData(r.Context(), strategyID, int(dataLength))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respondJSON(w, data)
}

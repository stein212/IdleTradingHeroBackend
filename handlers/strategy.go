package handlers

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/IdleTradingHeroServer/constants"
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
	respondJSON(w, &viewmodels.CreateStrategyResponse{
		ID: macdStrategy.ID,
	})
}

func (c *StrategyController) CreateMfiStrategy(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userID := getUserIDFromJWT(r)

	createMfiConfig := &viewmodels.CreateMfiConfig{}

	errPayload := decodeJSONBody(w, r, createMfiConfig)

	if errPayload != nil {
		strategyLogger.Println(errPayload)
		respondWithErrorPayload(logger, w, errPayload)
		return
	}

	// check request
	isValidPayload := validateStruct(createMfiConfig, w, logger)
	if !isValidPayload {
		// handled by validateStruct
		return
	}

	mfiConfig := &repositories.MfiConfig{
		Name:            createMfiConfig.Name,
		Instrument:      createMfiConfig.Instrument,
		Granularity:     createMfiConfig.Granularity,
		OverboughtLevel: createMfiConfig.OverboughtLevel,
		OversoldLevel:   createMfiConfig.OversoldLevel,
	}

	mfiStrategy, err := c.strategyRepository.CreateMfiStrategy(r.Context(), mfiConfig, userID)
	if err != nil {
		strategyLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	respondJSON(w, &viewmodels.CreateStrategyResponse{
		ID: mfiStrategy.ID,
	})
}

func (c *StrategyController) CreateRsiStrategy(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userID := getUserIDFromJWT(r)

	createRsiConfig := &viewmodels.CreateMfiConfig{}

	errPayload := decodeJSONBody(w, r, createRsiConfig)

	if errPayload != nil {
		strategyLogger.Println(errPayload)
		respondWithErrorPayload(logger, w, errPayload)
		return
	}

	// check request
	isValidPayload := validateStruct(createRsiConfig, w, logger)
	if !isValidPayload {
		// handled by validateStruct
		return
	}

	rsiConfig := &repositories.RsiConfig{
		Name:            createRsiConfig.Name,
		Instrument:      createRsiConfig.Instrument,
		Granularity:     createRsiConfig.Granularity,
		OverboughtLevel: createRsiConfig.OverboughtLevel,
		OversoldLevel:   createRsiConfig.OversoldLevel,
	}

	rsiStrategy, err := c.strategyRepository.CreateRsiStrategy(r.Context(), rsiConfig, userID)
	if err != nil {
		strategyLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	respondJSON(w, &viewmodels.CreateStrategyResponse{
		ID: rsiStrategy.ID,
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

	mfiStrategies, err := c.strategyRepository.GetMfiStrategies(r.Context(), userID)
	if err != nil {
		strategyLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rsiStrategies, err := c.strategyRepository.GetRsiStrategies(r.Context(), userID)
	if err != nil {
		strategyLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	payload := make([]*viewmodels.StrategyResponse, len(macdStrategies)+len(mfiStrategies)+len(rsiStrategies))
	i := 0
	for _, macdStrategy := range macdStrategies {
		payload[i] = viewmodels.MacdStrategyToStrategyResponse(macdStrategy)
		i++
	}
	for _, mfiStrategy := range mfiStrategies {
		payload[i] = viewmodels.MfiStrategyToStrategyResponse(mfiStrategy)
		i++
	}
	for _, rsiStrategy := range rsiStrategies {
		payload[i] = viewmodels.RsiStrategyToStrategyResponse(rsiStrategy)
		i++
	}

	for _, strategyRes := range payload {
		if strategyRes.Status == constants.StrategyNotDeployed ||
			strategyRes.Status == constants.StrategyDeleted {
			strategyRes.Balance = 0
			continue
		}

		balance, err := c.strategyRepository.GetBalance(r.Context(), strategyRes.ID)
		if err != nil {
			strategyRes.Balance = 0
			continue
		}

		strategyRes.Balance = balance
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

func (c *StrategyController) GetMfiStrategy(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userID := getUserIDFromJWT(r)
	strategyID := params.ByName("strategyID")

	mfiStrategy, err := c.strategyRepository.GetMfiStrategy(r.Context(), userID, strategyID)
	if err != nil {
		strategyLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	payload := viewmodels.MfiStrategyToMfiStrategyResponse(mfiStrategy)

	respondJSON(w, payload)
}

func (c *StrategyController) GetRsiStrategy(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userID := getUserIDFromJWT(r)
	strategyID := params.ByName("strategyID")

	rsiStrategy, err := c.strategyRepository.GetRsiStrategy(r.Context(), userID, strategyID)
	if err != nil {
		strategyLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	payload := viewmodels.RsiStrategyToRsiStrategyResponse(rsiStrategy)

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
	strategyType := params.ByName("strategyType")

	err := c.strategyRepository.StartStrategy(r.Context(), strategyID, strategyType)
	if err != nil {
		strategyLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (c *StrategyController) PauseStrategy(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	strategyID := params.ByName("strategyID")
	strategyType := params.ByName("strategyType")

	err := c.strategyRepository.PauseStrategy(r.Context(), strategyID, strategyType)
	if err != nil {
		strategyLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (c *StrategyController) GetStrategyData(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	strategyID := params.ByName("strategyID")
	dataLength, err := strconv.ParseInt(params.ByName("length"), 10, 64)
	if err != nil {
		strategyLogger.Println(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	data, err := c.strategyRepository.GetStrategyData(r.Context(), strategyID, int(dataLength))

	if err != nil {
		strategyLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respondJSON(w, data)
}

func (c *StrategyController) GetIndicatorData(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	strategyID := params.ByName("strategyID")
	dataLength, err := strconv.ParseInt(params.ByName("length"), 10, 64)
	if err != nil {
		strategyLogger.Println(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	data, err := c.strategyRepository.GetStrategyIndicatorData(r.Context(), strategyID, int(dataLength))

	if err != nil {
		strategyLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respondJSON(w, data)
}

func (c *StrategyController) GetPerformanceData(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	strategyID := params.ByName("strategyID")
	dataLength, err := strconv.ParseInt(params.ByName("length"), 10, 64)
	if err != nil {
		strategyLogger.Println(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	data, err := c.strategyRepository.GetStrategyPerformanceData(r.Context(), strategyID, int(dataLength))

	if err != nil {
		strategyLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respondJSON(w, data)
}

func (c *StrategyController) DeleteStrategy(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	strategyType := params.ByName("strategyType")
	strategyID := params.ByName("strategyID")

	err := c.strategyRepository.DeleteStrategy(r.Context(), strategyID, strategyType)
	if err != nil {
		strategyLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

package handlers

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/IdleTradingHeroServer/repositories"
	viewmodels "github.com/IdleTradingHeroServer/viewModels"
	"github.com/julienschmidt/httprouter"
)

var (
	strategyEventHookLogger = log.New(os.Stdout, "strategyEventHookController: ", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)
)

type StrategyEventHookController struct {
	strategyEventRepository repositories.StrategyEventRepository
}

func NewStrategyEventHookController(config *ControllerConfig) *StrategyEventHookController {
	return &StrategyEventHookController{
		strategyEventRepository: config.StrategyEventRepository,
	}
}

func (c *StrategyEventHookController) CreateStrategyEvent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	createStrategyEvent := &viewmodels.CreateStrategyEvent{}

	errPayload := decodeJSONBody(w, r, createStrategyEvent)

	if errPayload != nil {
		strategyLogger.Println(errPayload)
		respondWithErrorPayload(logger, w, errPayload)
		return
	}

	// check request
	isValidPayload := validateStruct(createStrategyEvent, w, strategyEventHookLogger)
	if !isValidPayload {
		// handled by validateStruct
		return
	}

	eventOn := time.Unix(createStrategyEvent.EventOn, 0)

	err := c.strategyEventRepository.CreateStrategyEvent(
		r.Context(),
		createStrategyEvent.StrategyID,
		createStrategyEvent.Action,
		createStrategyEvent.Amount,
		eventOn,
	)
	if err != nil {
		strategyLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *StrategyEventHookController) GetStrategyEvents(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	strategyID := params.ByName("strategyID")

	length, err := strconv.ParseInt(params.ByName("length"), 10, 64)
	if err != nil {
		strategyEventHookLogger.Println(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	to, err := strconv.ParseInt(params.ByName("to"), 10, 64)
	if err != nil {
		strategyEventHookLogger.Println(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	toTime := time.Unix(to, 0)

	events, err := c.strategyEventRepository.GetStrategyEvents(r.Context(), strategyID, int(length), toTime)

	if err != nil {
		strategyEventHookLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := make([]*viewmodels.StrategyEventResponse, len(events))

	for i, event := range events {
		data[i] = viewmodels.StrategyEventToStrategyEventResponse(event)
	}

	respondJSON(w, data)
}

func (c *StrategyEventHookController) GetUserStrategiesEvents(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userID := getUserIDFromJWT(r)

	length, err := strconv.ParseInt(params.ByName("length"), 10, 64)
	if err != nil {
		strategyEventHookLogger.Println(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	to, err := strconv.ParseInt(params.ByName("to"), 10, 64)
	if err != nil {
		strategyEventHookLogger.Println(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	toTime := time.Unix(to, 0)

	events, err := c.strategyEventRepository.GetUserStrategiesEvents(r.Context(), userID, int(length), toTime)

	if err != nil {
		strategyEventHookLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := make([]*viewmodels.StrategyEventResponse, len(events))

	for i, event := range events {
		data[i] = viewmodels.StrategyEventToStrategyEventResponse(event)
	}

	respondJSON(w, data)
}

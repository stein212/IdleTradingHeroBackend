package viewmodels

type StrategyConfig struct {
	Asset    string
	Strategy string
	Capital  float64
	Params   interface{}
}

type MACDParams struct {
	Ema26 int
	Ema12 int
	Ema9  int
}

type PostMACDConfig struct {
	Asset    string  `json:"asset" validate:"required"`
	Strategy string  `json:"strategy" validate:"required"`
	Capital  float64 `json:"capital" validate:"required"`
	Ema26    int     `json:"ema26" validate:"required"`
	Ema12    int     `json:"ema12" validate:"required"`
	Ema9     int     `json:"ema9" validate:"required"`
}

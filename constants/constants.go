package constants

// environmant variables
const (
	EnvJWTSecretKey         = "JWT_SECRET_KEY"
	EnvCookieHashKey        = "COOKIE_HASH_KEY"
	EnvCookieBlockKey       = "COOKIE_BLOCK_KEY"
	EnvPsqlConnectionString = "PSQL_CONNECTION_STRING"
	EnvAllowedOrigins       = "ALLOWED_ORIGINS"
	EnvDomain               = "DOMAIN"
	EnvStrategyGRPCHost     = "STRATEGY_GRPC_HOST"
)

// errors constants
const (
	ErrorCodeInvalidField    = "invalidField"
	ErrorCodeRequiredField   = "requiredField"
	ErrorCodeInvalidEmail    = "invalidEmail"
	ErrorCodeInvalidPassword = "invalidPassword"
	ErrorCodeUsernameTaken   = "usernameTaken"

	ErrorCodeInvalidContentType = "invalidContentType"
	ErrorCodeMalformedJSON      = "malformedJSON"

	ErrorCodeRequestTooLarge = "requestTooLarge"
)

// strategy statuses
const (
	StrategyNotDeployed = "notDeployed"
	StrategyDeployed    = "deployed"
	StrategyLive        = "live"
	StrategyPaused      = "paused"
)

const (
	StrategyTypeMacd = "macd"
)

var StrategyTypes = []string{StrategyTypeMacd}

// others
const (
	CookieAuthName = "jwt"
)

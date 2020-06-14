package constants

// environmant variables
const (
	EnvJWTSecretKey         = "JWT_SECRET_KEY"
	EnvCookieHashKey        = "COOKIE_HASH_KEY"
	EnvCookieBlockKey       = "COOKIE_BLOCK_KEY"
	EnvPsqlConnectionString = "PSQL_CONNECTION_STRING"
	EnvAllowedOrigins       = "ALLOWED_ORIGINS"
	EnvDomain               = "DOMAIN"
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

// Others
const (
	CookieAuthName = "jwt"
)

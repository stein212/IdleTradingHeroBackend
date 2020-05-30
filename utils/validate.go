package utils

import (
	"regexp"
	"sync"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	enTranslator ut.Translator
)

var (
	doOnce       = sync.Once{}
	validate     = validator.New()
	matchLower   = regexp.MustCompile(`[a-z]`)
	matchUpper   = regexp.MustCompile(`[A-Z]`)
	matchNumber  = regexp.MustCompile(`[0-9]`)
	matchSpecial = regexp.MustCompile(`[\!\@\#\$\%\^\&\*\(\\\)\-_\=\+\,\.\?\/\:\;\{\}\[\]~]`)
)

func GetValidator() *validator.Validate {
	doOnce.Do(func() {
		translator := en.New()
		universalTranslator := ut.New(translator, translator)
		enTranslator, _ = universalTranslator.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(validate, enTranslator)
		validate.RegisterTranslation("required", enTranslator, registerRequiredTranslation, requiredTranslation)
		validate.RegisterTranslation("email", enTranslator, registerEmailTranslation, emailTranslation)
		validate.RegisterTranslation("passwd", enTranslator, registerPasswordTranslation, passwordTranslation)

		validate.RegisterValidation("passwd", passwordValidation)
	})
	return validate
}

func GetEnTranslator() ut.Translator {
	return enTranslator
}

func registerRequiredTranslation(trans ut.Translator) error {
	return trans.Add("required", "{0} is required", true)
}

func requiredTranslation(trans ut.Translator, fe validator.FieldError) string {
	t, _ := trans.T("required", fe.Field())
	return t
}

func registerEmailTranslation(trans ut.Translator) error {
	return trans.Add("email", "Invalid email ({0})", true)
}

func emailTranslation(trans ut.Translator, fe validator.FieldError) string {
	t, _ := trans.T("email", fe.Field())
	return t
}

func passwordValidation(fieldLevel validator.FieldLevel) bool {
	field := fieldLevel.Field()

	password := field.String()

	if len(password) < 8 {
		return false
	}

	if !matchLower.MatchString(password) {
		return false
	}

	if !matchUpper.MatchString(password) {
		return false
	}

	if !matchNumber.MatchString(password) {
		return false
	}

	if !matchSpecial.MatchString(password) {
		return false
	}

	return true
}

func registerPasswordTranslation(trans ut.Translator) error {
	return trans.Add("passwd", "{0} requires at least 8 characters, 1 upper, 1 lower, 1 number, and 1 special", true)
}

func passwordTranslation(trans ut.Translator, fieldErr validator.FieldError) string {
	t, _ := trans.T("passwd", fieldErr.Field())
	return t
}

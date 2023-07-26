package validator

import (
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enMsg "github.com/go-playground/validator/v10/translations/en"
)

var (
	validate   *validator.Validate
	translator ut.Translator
)

func init() {
	validate = validator.New()
	en := en.New()
	translator, _ = ut.New(en, en).GetTranslator("en")
	enMsg.RegisterDefaultTranslations(validate, translator)
}

// Struct is ...
func Struct(input interface{}) map[string]string {
	if err := validate.Struct(input); err != nil {
		return buildTranslatedErrorMessages(err.(validator.ValidationErrors))
	}
	return nil
}

func buildTranslatedErrorMessages(err validator.ValidationErrors) map[string]string {
	errors := make(map[string]string)
	for _, err := range err {
		errors[strings.ToLower(err.Field())] = err.Translate(translator)
	}
	return errors
}

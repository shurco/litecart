package validator

import (
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enMsg "github.com/go-playground/validator/v10/translations/en"
)

// Struct is ...
func Struct(input interface{}) *map[string]string {
	validate := validator.New()

	en := en.New()
	translator, _ := ut.New(en, en).GetTranslator("en")
	enMsg.RegisterDefaultTranslations(validate, translator)

	if err := validate.Struct(input); err != nil {
		return buildTranslatedErrorMessages(err.(validator.ValidationErrors), translator)
	}
	return nil
}

func buildTranslatedErrorMessages(err validator.ValidationErrors, translator ut.Translator) *map[string]string {
	errors := make(map[string]string)
	for _, err := range err {
		errors[strings.ToLower(err.Field())] = err.Translate(translator)
	}
	return &errors
}

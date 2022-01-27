package val

import (
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni        *ut.UniversalTranslator
	validate   *validator.Validate
	english    locales.Translator
	translator ut.Translator
)

func init() {
	// 1. Register the validator
	validate = validator.New()

	// 2. Create the english locale
	english = en.New()

	// 3. Create the translator
	uni = ut.New(english, english)

	// 4. Add english language to translator
	translator, _ = uni.GetTranslator("en")

	// 5. Register english translations to the validator
	_ = en_translations.RegisterDefaultTranslations(validate, translator)
}

func Run(v interface{}) map[string]string {
	if e := validate.Struct(v); e != nil {
		if _, ok := e.(*validator.InvalidValidationError); ok {
			return map[string]string{
				"InvalidValidationError": e.Error(),
			}
		}

		return (e.(validator.ValidationErrors)).Translate(translator)
	}

	return nil
}

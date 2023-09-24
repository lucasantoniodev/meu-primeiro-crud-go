package validation

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	validatorLanguage "github.com/go-playground/validator/v10/translations/en"
	"github.com/lucasantoniodev/meu-primeiro-crud-go/src/configuration/rest_err"
)

var (
	Validate   = validator.New()
	translator ut.Translator
)

/* Configuração de idioma/tradução do validator */
func init() {
	// Capturando a instância do validator do GinGonic e fazendo um casting para o tipo Validate do pacote playground
	validatorGinGonic, ok := binding.Validator.Engine().(*validator.Validate)

	if ok {
		// Translator
		language := en.New()

		// Gerenciador de translators
		universalTranslator := ut.New(language)

		// Obtendo um translator de um idioma e atribuindo a variavel global
		translator, _ = universalTranslator.GetTranslator(language.Locale())

		// Registra o tradutor default do validator
		if err := validatorLanguage.RegisterDefaultTranslations(validatorGinGonic, translator); err != nil {
			return
		}

	}
}

/* 		Método para personalizar o erro de acordo com um erro causado no request	*/

func ValidateUserError(validationErr error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	/* Verificando se o tipo de erro é igual ao erro de Tipo de valor de um campo (Error de parser)*/
	if errors.As(validationErr, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
	} else if errors.As(validationErr, &jsonValidationError) {
		var errorsCauses []rest_err.Causes

		// Iterando sobre cada uma das causas do erro, necessário o casting para identificação do objeto
		for _, e := range validationErr.(validator.ValidationErrors) {
			errorsCauses = append(errorsCauses, rest_err.Causes{
				Field:   e.Field(),
				Message: e.Translate(translator),
			})
		}

		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	}

	return rest_err.NewBadRequestError("Error trying to convert fields")
}

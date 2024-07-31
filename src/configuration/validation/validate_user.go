package validation

import (
	"encoding/json"
	"errors"

	rest_err "github.com/DavibernardesA/CRUD-Go/src/configuration/api_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	validate = validator.New()
	transL   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		uni := ut.New(en, en)
		transL, _ = uni.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transL)
	}
}

func ValidateUserError(validationErr error) *rest_err.ApiErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type.")
	} else if errors.As(validationErr, &jsonValidationError) {
		errorsCauses := []rest_err.Causes{}

		for _, e := range validationErr.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(transL),
				Field:   e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}
		return rest_err.NewBadRequestValidationError("Some fields are invalids.", errorsCauses)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields.")
	}
}

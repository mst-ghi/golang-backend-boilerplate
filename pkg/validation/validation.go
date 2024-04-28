package validation

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

var errorsList = make(map[string]string)

func Handle(err error) map[string]string {
	errorsList = map[string]string{}

	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			Add(fieldError.Field(), GetErrorMsg(fieldError.Tag()))
		}
	}

	return errorsList
}

func Add(key string, value string) {
	errorsList[strings.ToLower(key)] = value
}

func GetErrorMsg(tag string) string {
	return ErrorMessages()[tag]
}

func Get() map[string]string {
	return errorsList
}

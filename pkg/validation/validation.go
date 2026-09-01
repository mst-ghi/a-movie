package validation

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func Handle(err error) map[string]string {
	result := map[string]string{}

	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			result[strings.ToLower(fieldError.Field())] = GetErrorMsg(fieldError.Tag())
		}
	}

	return result
}

func GetErrorMsg(tag string) string {
	return ErrorMessages()[tag]
}

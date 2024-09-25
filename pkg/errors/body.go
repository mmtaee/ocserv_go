package errors

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"strings"
)

func InvalidBodyError(err error) map[string]interface{} {
	if !errors.As(err, &validator.ValidationErrors{}) {
		errSplit := strings.Split(err.Error(), ".")
		return map[string]interface{}{
			"error": errSplit[len(errSplit)-1],
		}
	}

	var validationErrors []string
	for _, e := range err.(validator.ValidationErrors) {
		validationErrors = append(validationErrors, formatError(e))
	}
	return map[string]interface{}{
		"error": validationErrors,
	}
}

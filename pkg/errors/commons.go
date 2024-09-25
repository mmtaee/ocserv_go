package errors

import (
	"github.com/go-playground/validator/v10"
	"unicode"
)

func formatSnakeCase(s string) string {
	var result []rune
	for i, char := range s {
		if unicode.IsUpper(char) && i > 0 {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(char))
	}
	return string(result)
}

func formatError(err validator.FieldError) string {
	field := formatSnakeCase(err.Field())
	switch err.Tag() {
	case "required":
		return field + " is required"
	case "min":
		return field + " must be at least " + err.Param() + " characters long"
	case "max":
		return field + " must be at most " + err.Param() + " characters long"
	case "oneof":
		return field + " can be one of " + err.Param()
	case "gt":
		return field + " must be greater than " + err.Param()
	default:
		return "Invalid input"
	}
}

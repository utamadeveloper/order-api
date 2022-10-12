package config

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetErrorMsg(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "alphanum":
		return "This field must be alphanumeric"
	case "unique":
		return "This field must be unique"
	default:
		return "This field is invalid"
	}
}

func GetErrorValidator(err error) []FieldError {
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		out := make([]FieldError, len(ve))
		for i, fe := range ve {
			out[i] = FieldError{fe.Field(), GetErrorMsg(fe.Tag())}
		}
		return out
	}
	return nil
}

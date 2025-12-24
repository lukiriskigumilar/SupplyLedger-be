package utils

import validation "github.com/go-ozzo/ozzo-validation/v4"

func ParseValidationError(err error) map[string]string {
	errors := make(map[string]string)

	if validationErrors, ok := err.(validation.Errors); ok {
		for fields, fieldErr := range validationErrors {
			errors[fields] = fieldErr.Error()
		}
	}
	return errors
}

package validator

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

func getTagMessage(errTag string) (string, error) {
	switch errTag {
	case "required":
		return "field is required", nil
	case "email":
		return "invalid email", nil
	}
	return "", fmt.Errorf("can't find message for %s tag", errTag)
}

func buildFieldError(e validator.FieldError) error {
	errMsg, err := getTagMessage(e.Tag())
	if err != nil {
		errMsg = e.Error()
	}
	return fmt.Errorf("%s: %s", e.Field(), errMsg)
}

func BuildValidationErrors(vErr error) (errs error) {
	var vErrs validator.ValidationErrors
	if errors.As(vErr, &vErrs) {
		for _, e := range vErrs {
			errs = errors.Join(errs, fmt.Errorf("%w", buildFieldError(e)))
		}
	}
	return errs
}

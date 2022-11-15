package query

import (
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// Define a single validator to do all of the validations for us.
var v = validator.New()

// ValidatedQueryParam extracts a query parameter and validates it.
func ValidatedQueryParam(ctx echo.Context, name, validationTag string) (string, error) {
	value := ctx.QueryParam(name)

	// Validate the value.
	if err := v.Var(value, validationTag); err != nil {
		return "", err
	}

	return value, nil
}

// ValidateBooleanQueryParam extracts a Boolean query parameter and validates it.
func ValidateBooleanQueryParam(ctx echo.Context, name string, defaultValue *bool) (bool, error) {
	errMsg := fmt.Sprintf("invalid query parameter: %s", name)
	value := ctx.QueryParam(name)

	// Assume that the parameter is required if there's no default.
	if defaultValue == nil {
		if err := v.Var(value, "required"); err != nil {
			return false, fmt.Errorf("missing required query parameter: %s", name)
		}
	}

	// If no value was provided at this point then the prameter is optional; return the default value.
	if value == "" {
		return *defaultValue, nil
	}

	// Parse the parameter value and return the result.
	result, err := strconv.ParseBool(value)
	if err != nil {
		return false, errors.Wrap(err, errMsg)
	}
	return result, nil
}

package helper

import (
	"errors"
)

var (
	LogStrKeyLevel = "cadana"
)

// CustomError() is used return custom errors
func CustomError(value string) error {
	if len(value) == 0 {
		return ErrDefaultError
	}

	return errors.New(value)
}

// GetStringPointerValue returns the value of the string pointer or an empty string if it is nil
func GetStringPointerValue(str *string) string {
	var val string
	if str != nil {
		return *str
	}

	return val
}

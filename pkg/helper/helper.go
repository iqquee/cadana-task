package helper

import (
	"errors"
)

// CustomError() is used return custom errors
func CustomError(value string) error {
	if len(value) == 0 {
		return ErrDefaultError
	}

	return errors.New(value)
}

// GetStringPointerValue returns string from pointer
func GetStringPointerValue(str *string) string {
	var val string
	if str != nil {
		return *str
	}

	return val
}

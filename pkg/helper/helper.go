package helper

import (
	"errors"
)

const (
	// LogStrKeyModule log service name value
	LogStrKeyModule = "ser_name"
	// LogStrPartnerLevel log partner name value
	LogStrPartnerLevel = "partner_name"
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

// GetStringPointer returns a string pointer
func GetStringPointer(val string) *string {
	return &val
}

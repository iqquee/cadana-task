package helper

import "errors"

var (
	// ErrDefaultError is the default error
	ErrDefaultError = errors.New("an error occured, please try again")
	// ErrInvalidCurrency is the error for an invalid currency type
	ErrInvalidCurrency = errors.New("invalid currency type")
	// ErrExchangeRateEmpty is the error for an empty exchange rate value
	ErrExchangeRateEmpty = errors.New("exchange rate value is empty")
)

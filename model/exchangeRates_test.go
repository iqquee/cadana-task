package model

import (
	"cadana/pkg/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestValidateCurrencyTypes() tests the ValidateCurrencyTypes method
func TestValidateCurrencyTypes(t *testing.T) {
	ex := ExchangeRateReq{}

	// Test case 1 -  Valid currency type
	value1 := []string{"USD", "EUR"}
	expected1 := error(nil)
	result1 := ex.ValidateCurrencyTypes(value1)
	assert.Equal(t, expected1, result1)

	// Test case 2 -  Invalid currency type
	value2 := []string{"USD", "CADANA"}
	expected2 := helper.ErrInvalidCurrency
	result2 := ex.ValidateCurrencyTypes(value2)
	assert.Equal(t, expected2, result2)
}

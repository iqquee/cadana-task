package rates

import (
	"cadana/model"
	"cadana/pkg/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCurrencyServerA tests the CurrencyServerA function
func TestCurrencyServerA(t *testing.T) {
	var rates Rates

	testCases := []struct {
		name          string
		from          string
		to            string
		rate          float64
		expectedError error
	}{
		{
			name:          "empty from value",
			from:          "",
			to:            "USD",
			expectedError: helper.ErrConvertCurrencyValueMissing,
		},
		{
			name:          "empty to value",
			from:          "GBP",
			to:            "",
			expectedError: helper.ErrConvertCurrencyValueMissing,
		},
		{
			name:          "valid values",
			from:          "GBP",
			to:            "USD",
			expectedError: nil,
			rate:          0.92,
		},
	}

	validResonse := &model.ExchangeRateServerResponse{}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.name != "valid values" {
				_, err := rates.CurrencyServerA(testCase.from, testCase.to)
				assert.Error(t, err)
				assert.Equal(t, testCase.expectedError, err)
			} else {
				response, err := rates.CurrencyServerA(testCase.from, testCase.to)
				assert.NoError(t, err)
				assert.Equal(t, testCase.expectedError, err)

				validResonse.From = testCase.from
				validResonse.To = testCase.to
				validResonse.Rate = testCase.rate
				assert.Equal(t, validResonse, response)
			}
		})
	}

}

// TestCurrencyServerB tests the CurrencyServerB function
func TestCurrencyServerB(t *testing.T) {
	var rates Rates

	testCases := []struct {
		name          string
		from          string
		to            string
		rate          float64
		expectedError error
	}{
		{
			name:          "empty from value",
			from:          "",
			to:            "EUR",
			expectedError: helper.ErrConvertCurrencyValueMissing,
		},
		{
			name:          "empty to value",
			from:          "USD",
			to:            "",
			expectedError: helper.ErrConvertCurrencyValueMissing,
		},
		{
			name:          "valid values",
			from:          "USD",
			to:            "EUR",
			expectedError: nil,
			rate:          1.27,
		},
	}

	validResonse := &model.ExchangeRateServerResponse{}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.name != "valid values" {
				_, err := rates.CurrencyServerB(testCase.from, testCase.to)
				assert.Error(t, err)
				assert.Equal(t, testCase.expectedError, err)
			} else {
				response, err := rates.CurrencyServerB(testCase.from, testCase.to)
				assert.NoError(t, err)
				assert.Equal(t, testCase.expectedError, err)

				validResonse.From = testCase.from
				validResonse.To = testCase.to
				validResonse.Rate = testCase.rate
				assert.Equal(t, validResonse, response)
			}
		})
	}

}

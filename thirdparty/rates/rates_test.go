package rates

import (
	"cadana/pkg/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestServerA tests the ServerA function
func TestServerA(t *testing.T) {
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

	validResonse := &ExchangeRateServerResponse{}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.name != "valid values" {
				_, err := rates.ServerA(testCase.from, testCase.to)
				assert.Error(t, err)
				assert.Equal(t, testCase.expectedError, err)
			} else {
				response, err := rates.ServerA(testCase.from, testCase.to)
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

// TestServerB tests the ServerB function
func TestServerB(t *testing.T) {
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

	validResonse := &ExchangeRateServerResponse{}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.name != "valid values" {
				_, err := rates.ServerB(testCase.from, testCase.to)
				assert.Error(t, err)
				assert.Equal(t, testCase.expectedError, err)
			} else {
				response, err := rates.ServerB(testCase.from, testCase.to)
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

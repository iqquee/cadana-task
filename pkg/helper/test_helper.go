package helper

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCustomError() tests the CustomError function
func TestCustomError(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expectedError error
	}{
		{
			name:          "Value Empty",
			input:         "",
			expectedError: ErrDefaultError,
		},
		{
			name:          "Value not empty",
			input:         "custom error",
			expectedError: errors.New("custom error"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := CustomError(testCase.input)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

// TestGetStringPointerValue() tests the GetStringPointerValue function
func TestGetStringPointerValue(t *testing.T) {
	// Test case 1 - String pointer is not nil
	str1 := "test"
	expected1 := "test"
	result1 := GetStringPointerValue(&str1)
	assert.Equal(t, expected1, result1)

	// Test case 2 - String pointer is nil
	var str2 *string
	expected2 := ""
	result2 := GetStringPointerValue(str2)
	assert.Equal(t, expected2, result2)
}

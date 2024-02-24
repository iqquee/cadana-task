package model

import (
	"cadana/pkg/helper"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var (
	// errorOpeningCurrencyTypeFile is the string value for error opening currency json file
	errorOpeningCurrencyTypeFile = "error opening JSON file"
	// errorDecodingJsonFile is the string value for error decoding json file
	errorDecodingJsonFile = "error decoding JSON file"
	// currencyTypesFileDir is the file location to the currency json file
	currencyTypesFileDir = "./pkg/currency/currency.json"
)

type (
	// ExchangeRate is request object for an exchange rate
	ExchangeRate struct {
		CurrencyPair string `json:"currency-pair"`
	}

	// currencyType is an object for the currency types
	currencyType struct {
		ID       string `json:"ID"`
		Currency string `json:"Currency"`
	}
)

// Validate() validates the ExchangeRate object request
func (ex ExchangeRate) Validate() (*ExchangeRate, error) {
	// check if the CurrencyPair value is not empty
	if len(ex.CurrencyPair) == 0 {
		return nil, helper.ErrExchangeRateEmpty
	}

	// check if the CurrencyPair value is in this format e.g USD-EUR
	if !strings.Contains(ex.CurrencyPair, "-") || strings.Contains(ex.CurrencyPair, " ") {
		return nil, helper.ErrInvalidCurrency
	} else {
		// check if the currencies are valid currency types
		currencyTypeValues := strings.Split(ex.CurrencyPair, "-")
		if err := ex.ValidateCurrencyTypes(currencyTypeValues); err != nil {
			return nil, err
		}
	}

	return &ex, nil
}

// ValidateCurrencyTypes() validates the currency types from the ExchangeRate request object
func (ex ExchangeRate) ValidateCurrencyTypes(value []string) error {
	var currencyTypes []currencyType
	file, err := os.Open(currencyTypesFileDir)
	if err != nil {
		return helper.CustomError(fmt.Sprintf("%s ::: error message: %v", errorOpeningCurrencyTypeFile, err))
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&currencyTypes); err != nil {
		return helper.CustomError(fmt.Sprintf("%s ::: error message: %s", errorDecodingJsonFile, err))
	}

	for _, currencyType := range currencyTypes {
		if strings.ToUpper(value[0]) != currencyType.Currency || strings.ToUpper(value[1]) != currencyType.Currency {
			return helper.ErrInvalidCurrency
		}
	}

	return nil
}

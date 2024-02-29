package model

import (
	"cadana/pkg/helper"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
)

// ExchangeRateJSONFileLocation is the file location to the currency.json file
const ExchangeRateJSONFileLocation = "/pkg/currency/currency.json"

var (
	// errorUnmarshalingJsonFile is the string value for error unmarshaling json file
	errorUnmarshalingJsonFile = "error unmarshal JSON file"
	// errorGettingWorkingDir is the string value for error getting current directory
	errorGettingWorkingDir = "error getting current directory"
	// errorReadingJson is the string value for error rading json file
	errorReadingJson = "error reading JSON file"
)

type (
	// ExchangeRateReq is request object for an exchange rate
	ExchangeRateReq struct {
		CurrencyPair string `json:"currency-pair"`
	}

	// ExchangeRateRes is response object for an exchange rate
	ExchangeRateRes struct {
		From string  `json:"from"`
		To   string  `json:"to"`
		Rate float64 `json:"rate"`
	}

	// currencyType is an object for the currency types
	currencyType struct {
		ID       string `json:"ID"`
		Currency string `json:"Currency"`
	}

	// ExchangeRateServerResponse is response object from the server for an exchange rate
	ExchangeRateServerResponse struct {
		From  string  `json:"from"`
		To    string  `json:"to"`
		Rate  float64 `json:"rate"`
		Error error   `json:"error"`
	}
)

// Validate() validates the ExchangeRate object request
func (ex ExchangeRateReq) Validate() ([]string, error) {
	log.Info().Msg("Validation started...")
	// check if the CurrencyPair value is not empty
	if len(ex.CurrencyPair) == 0 {
		return nil, helper.ErrExchangeRateEmpty
	}

	var currencyTypes []string
	// check if the CurrencyPair value is in this format e.g USD-EUR
	if !strings.Contains(ex.CurrencyPair, "-") || strings.Contains(ex.CurrencyPair, " ") {
		return nil, helper.ErrInvalidCurrency
	} else {
		// check if the currencies are valid currency types
		currencyTypeValues := strings.Split(ex.CurrencyPair, "-")
		if err := ex.ValidateCurrencyTypes(currencyTypeValues); err != nil {
			return nil, err
		}

		currencyTypes = currencyTypeValues
	}

	return currencyTypes, nil
}

// ValidateCurrencyTypes() validates the currency types from the ExchangeRate request object
func (ex ExchangeRateReq) ValidateCurrencyTypes(value []string) error {
	var currencyTypes []currencyType

	currentDir, err := os.Getwd()
	if err != nil {
		return helper.CustomError(fmt.Sprintf("%s ::: error message: %v", errorGettingWorkingDir, err))
	}

	filePath := filepath.Join(currentDir, ExchangeRateJSONFileLocation)

	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return helper.CustomError(fmt.Sprintf("%s ::: error message: %v", errorReadingJson, err))
	}

	err = json.Unmarshal(jsonData, &currencyTypes)
	if err != nil {
		return helper.CustomError(fmt.Sprintf("%s ::: error message: %v", errorUnmarshalingJsonFile, err))
	}

	channelErrNil := make(chan bool)
	done := make(chan bool)
	var currencyErr bool

	if len(value) == 2 {
		go func() {
			for _, currencyType := range currencyTypes {
				if strings.ToUpper(value[0]) == currencyType.Currency {
					channelErrNil <- true
				}

				if strings.ToUpper(value[1]) == currencyType.Currency {
					channelErrNil <- true
				}

			}

			done <- true
		}()

		go func() {
			for {
				select {
				case err := <-channelErrNil:
					currencyErr = err
				case <-done:
					return
				}
			}
		}()
		<-done

	} else {
		for _, currencyType := range currencyTypes {
			if strings.ToUpper(value[0]) == currencyType.Currency {
				return nil
			}
		}
	}

	if !currencyErr {
		return helper.ErrInvalidCurrency
	}

	return nil
}

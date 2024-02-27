package model

import (
	"cadana/pkg/helper"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	// errorOpeningCurrencyTypeFile is the string value for error opening currency json file
	errorOpeningCurrencyTypeFile = "error opening JSON file"
	// errorDecodingJsonFile is the string value for error decoding json file
	errorDecodingJsonFile = "error decoding JSON file"
	// currencyTypesFileDir is the file location to the currency json file
	currencyTypesFileDir = "pkg/currency/currency.json"
	// errorGettingWorkingDir is the string value for error getting current directory
	errorGettingWorkingDir = "error getting current directory"
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
)

// Validate() validates the ExchangeRate object request
func (ex ExchangeRateReq) Validate() (*ExchangeRateReq, error) {
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
func (ex ExchangeRateReq) ValidateCurrencyTypes(value []string) error {
	var currencyTypes []currencyType
	currentDir, err := os.Getwd()
	if err != nil {
		return helper.CustomError(fmt.Sprintf("%s ::: error message: %v", errorGettingWorkingDir, err))
	}

	var projectDir string
	for {
		if filepath.Base(currentDir) == "cadana-task" {
			projectDir = currentDir
			break
		}

		currentDir = filepath.Dir(currentDir)
	}

	filePath := filepath.Join(projectDir, currencyTypesFileDir)
	file, err := os.Open(filePath)
	if err != nil {
		return helper.CustomError(fmt.Sprintf("%s ::: error message: %v", errorOpeningCurrencyTypeFile, err))
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&currencyTypes); err != nil {
		return helper.CustomError(fmt.Sprintf("%s ::: error message: %s", errorDecodingJsonFile, err))
	}

	channelErrNil := make(chan bool)
	done := make(chan bool)

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

	var currencyErr bool
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

	if !currencyErr {
		return helper.ErrInvalidCurrency
	}

	return nil
}

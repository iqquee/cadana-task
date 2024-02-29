package person

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"cadana/controller"
	"cadana/model"
	"cadana/pkg/helper"
)

// PersonJSONFileLocation is the file location to the person.json file
const PersonJSONFileLocation = "/pkg/person/person.json"

type (
	// Person is the object of a person.
	Person struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Salary Salary `json:"salary"`
	}

	// Salary is the object of salary of a person
	Salary struct {
		Balance  float64 `json:"balance"`
		Currency string  `json:"currency"`
	}

	// Persons is a collection of Persons objects.
	Persons struct {
		Data []Person `json:"data"`
	}
)

// GroupByCurrency groups persons by salary currency type into hash maps.
func (p Persons) GroupByCurrency() map[string][]Person {
	personsGroup := make(map[string][]Person)

	for _, person := range p.Data {
		personsGroup[person.Salary.Currency] = append(personsGroup[person.Salary.Currency], person)
	}

	return personsGroup
}

// FilterByCurrency filters persons by salary currency and converts the salaries to USD.
func (p Persons) FilterByCurrency(currency string, c controller.Operations) ([]Person, error) {
	var filteredPersons []Person

	currencyExObj := model.ExchangeRateReq{
		CurrencyPair: currency,
	}

	var currencyType []string
	currencyType = append(currencyType, currency)

	if err := currencyExObj.ValidateCurrencyTypes(currencyType); err != nil {
		return nil, err
	}

	for _, person := range p.Data {
		// if currency type is already in USD then ignore
		if person.Salary.Currency == currency {
			filteredPersons = append(filteredPersons, person)
		} else {
			// get the currency exchange rates for currencies that are not in USD to USD
			exchangeRate, err := c.ServerResponse(person.Salary.Currency, currency)
			if err != nil {
				fmt.Printf("Failed to fetch exchange rate for %s: %v\n", person.Salary.Currency, err)
				continue
			}

			person.Salary.Balance = person.Salary.Balance * exchangeRate.Rate
			person.Salary.Currency = currency
			filteredPersons = append(filteredPersons, person)
		}
	}

	return filteredPersons, nil
}

// UnmarshalPersonJSON unmarshal the datamanipulation.json file into Persons struct{}
func (p Persons) UnmarshalPersonJSON() (Persons, error) {
	var persons Persons

	currentDir, err := os.Getwd()
	if err != nil {
		return Persons{}, helper.CustomError(fmt.Sprintf("os.Getwd ::: %v", err))
	}

	filePath := filepath.Join(currentDir, PersonJSONFileLocation)

	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return Persons{}, helper.CustomError(fmt.Sprintf("os.ReadFile ::: %v", err))
	}

	if err = json.Unmarshal(jsonData, &persons.Data); err != nil {
		return Persons{}, helper.CustomError(fmt.Sprintf("json.Unmarshal ::: %v", err))
	}

	return persons, nil
}

// FilterFromAscToDesc filters persons salary from ascending to descending or reverse
func (p Persons) FilterFromAscToDesc(sortDir string) (Persons, error) {

	if strings.ToUpper(sortDir) == "ASC" {
		sort.Slice(p.Data, func(i, j int) bool {
			return p.Data[i].Salary.Balance > p.Data[j].Salary.Balance
		})
	} else if strings.ToUpper(sortDir) == "DESC" {
		sort.Slice(p.Data, func(i, j int) bool {
			return p.Data[i].Salary.Balance < p.Data[j].Salary.Balance
		})
	} else {
		return Persons{}, helper.CustomError("Invalid sort direction")
	}

	return p, nil
}

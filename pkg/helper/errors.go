package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	// ErrDefaultError is the default error
	ErrDefaultError = errors.New("an error occured, please try again")
	// ErrInvalidCurrency is the error for an invalid currency type
	ErrInvalidCurrency = errors.New("invalid currency type")
	// ErrExchangeRateEmpty is the error for an empty exchange rate value
	ErrExchangeRateEmpty = errors.New("exchange rate value is empty")
	// ErrConvertCurrencyValueMissing is the error for empty currency(s) params
	ErrConvertCurrencyValueMissing = errors.New("you must pass currencies values to convert")
	// ErrBothServicesResponse is the error when both services returned an error
	ErrBothServicesResponse = errors.New("both services returned errors")
)

// ExchangeRateResponse represents the JSON response structure
type ExchangeRateResponse struct {
	Rate  float64 `json:"rate"`
	Error string  `json:"error,omitempty"`
}

// APIService simulates an external API service
type APIService struct {
	Name string
}

// GetExchangeRate simulates fetching exchange rates from an external API
func (s *APIService) GetExchangeRate(apiKey string) ExchangeRateResponse {
	// Simulate delay
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	// Simulate response
	return ExchangeRateResponse{Rate: rand.Float64() * 10}
}

// AWSKeyManager simulates the secure retrieval of API keys from AWS Secrets Manager
type AWSKeyManager struct{}

// GetAPIKey simulates fetching an API key from AWS Secrets Manager
func (k *AWSKeyManager) GetAPIKey(serviceName string) (string, error) {
	// Simulate fetching the API key from AWS Secrets Manager
	// In a real scenario, you would use AWS SDK to interact with AWS Secrets Manager
	// For the sake of simplicity, we return a hardcoded key for simulation purposes
	return "dummy-api-key", nil
}

func v8() {
	// Create a mock AWSKeyManager
	keyManager := &AWSKeyManager{}

	// Simulate fetching API keys for Service A and Service B
	apiKeyA, err := keyManager.GetAPIKey("ServiceA")
	if err != nil {
		fmt.Println("Error fetching API key for Service A:", err)
		return
	}

	apiKeyB, err := keyManager.GetAPIKey("ServiceB")
	if err != nil {
		fmt.Println("Error fetching API key for Service B:", err)
		return
	}

	// Create instances of the external API services
	serviceA := &APIService{Name: "ServiceA"}
	serviceB := &APIService{Name: "ServiceB"}

	// Use wait groups to synchronize goroutines
	var wg sync.WaitGroup
	wg.Add(2)

	// Channel to receive the first response
	responseChan := make(chan ExchangeRateResponse)

	// Call Service A concurrently
	go func() {
		defer wg.Done()
		response := serviceA.GetExchangeRate(apiKeyA)
		responseChan <- response
	}()

	// Call Service B concurrently
	go func() {
		defer wg.Done()
		response := serviceB.GetExchangeRate(apiKeyB)
		responseChan <- response
	}()

	// Wait for both goroutines to finish
	wg.Wait()

	// Close the channel to avoid deadlock
	close(responseChan)

	// Retrieve the first response from the channel
	response := <-responseChan

	// Convert the response to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Print the JSON response
	fmt.Println("First response:", string(jsonResponse))
}

package rates

import (
	"cadana/pkg/environment"
	"cadana/pkg/helper"
	"math/rand"
	"time"

	"github.com/rs/zerolog"
)

const (
	packageName = "rates"
)

// ExchangeRateService enlist all possible operations for exchange rates in the platform
type ExchangeRateService interface {
	ServerA(from, to string) (*ExchangeRateServerResponse, error)
	ServerB(from, to string) (*ExchangeRateServerResponse, error)
}

type (
	// ExchangeRateServerResponse is response object from the server for an exchange rate
	ExchangeRateServerResponse struct {
		From string  `json:"from"`
		To   string  `json:"to"`
		Rate float64 `json:"rate"`
	}
)

// Rates represents rates instance
type Rates struct {
	APIKey string `json:"apiKey"`
	env    *environment.Env
	logger zerolog.Logger
}

// New creates a an instance of ExchangeRateService and returns an error if any occurs
func New(z zerolog.Logger, ev *environment.Env) *ExchangeRateService {
	l := z.With().Str(helper.LogStrKeyLevel, packageName).Logger()

	r := &Rates{
		env:    ev,
		logger: l,
	}
	ex := ExchangeRateService(r)
	return &ex
}

// ServerA is a mock server response
func (r Rates) ServerA(from, to string) (*ExchangeRateServerResponse, error) {
	if len(from) == 0 || len(to) == 0 {
		return nil, helper.ErrConvertCurrencyValueMissing
	}

	// Simulatate server processing delay
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	return &ExchangeRateServerResponse{
		From: from,
		To:   to,
		Rate: 0.92,
	}, nil
}

// ServerB is a mock server response
func (r Rates) ServerB(from, to string) (*ExchangeRateServerResponse, error) {
	if len(from) == 0 || len(to) == 0 {
		return nil, helper.ErrConvertCurrencyValueMissing
	}

	// Simulatate server processing delay
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	return &ExchangeRateServerResponse{
		From: from,
		To:   to,
		Rate: 1.27,
	}, nil
}

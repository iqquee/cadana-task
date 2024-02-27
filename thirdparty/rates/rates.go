package rates

import (
	"cadana/model"
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
	CurrencyServerA(from, to string) (*model.ExchangeRateServerResponse, error)
	CurrencyServerB(from, to string) (*model.ExchangeRateServerResponse, error)
}

// Rates represents rates instance
type Rates struct {
	env    *environment.Env
	logger zerolog.Logger
}

// New creates a an instance of ExchangeRateService and returns an error if any occurs
func New(z zerolog.Logger, ev *environment.Env) *ExchangeRateService {
	l := z.With().Str(helper.LogStrPartnerLevel, packageName).Logger()

	r := &Rates{
		env:    ev,
		logger: l,
	}

	ex := ExchangeRateService(r)
	return &ex
}

// ServerA is a mock server response
func (r Rates) CurrencyServerA(from, to string) (*model.ExchangeRateServerResponse, error) {
	if len(from) == 0 || len(to) == 0 {
		return nil, helper.ErrConvertCurrencyValueMissing
	}

	// gets the API key for this request and log the value since er are not actually using it- we are just mocking
	apiKey := r.env.MockGet("API_KEY")
	r.logger.Info().Msgf("API key value ::: %s", apiKey)

	// Simulatate server processing delay
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	// server response
	return &model.ExchangeRateServerResponse{
		From: from,
		To:   to,
		Rate: 0.92,
	}, nil
}

// ServerB is a mock server response
func (r Rates) CurrencyServerB(from, to string) (*model.ExchangeRateServerResponse, error) {
	if len(from) == 0 || len(to) == 0 {
		return nil, helper.ErrConvertCurrencyValueMissing
	}

	// gets the API key for this request and log the value since er are not actually using it- we are just mocking
	apiKey := r.env.MockGet("API_KEY")
	r.logger.Info().Msgf("API key value ::: %s", apiKey)

	// Simulatate server processing delay
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	// server response
	return &model.ExchangeRateServerResponse{
		From: from,
		To:   to,
		Rate: 1.27,
	}, nil
}

// Package controller defines implementation that exposes logics of the app
package controller

import (
	"github.com/rs/zerolog"

	"cadana/model"
	"cadana/pkg/environment"
	"cadana/pkg/helper"
	"cadana/thirdparty/rates"
)

const packageName = "controller"

// Operations enlist all possible operations for this controller across all modules
type Operations interface {
	CurrencyServerA(from, to string) (*model.ExchangeRateServerResponse, error)
	CurrencyServerB(from, to string) (*model.ExchangeRateServerResponse, error)

	ServerResponse(from, to string) (*model.ExchangeRateServerResponse, error)
}

// Controller object to hold necessary reference to other dependencies
type Controller struct {
	logger zerolog.Logger
	env    *environment.Env

	// rates service layer
	ratesService rates.ExchangeRateService
}

// New creates a new instance of Controller
func New(z zerolog.Logger) *Operations {
	l := z.With().Str(helper.LogStrKeyModule, packageName).Logger()

	// init all package layer dependencies under here
	env, err := environment.New(l)
	if err != nil {
		l.Fatal().Err(err)
		panic(err)
	}

	envVars := make(map[string]string)
	envVars["API_KEY"] = "test_API_KEY"
	envVars["AWS_REGION"] = "region"

	// set mock environmental variable(s)
	env.MockEnv(envVars)

	currencyRate := rates.New(l, env)

	ctrl := &Controller{
		logger:       l,
		env:          env,
		ratesService: *currencyRate,
	}

	op := Operations(ctrl)
	return &op
}

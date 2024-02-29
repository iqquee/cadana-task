package controller

import (
	"cadana/model"
	"cadana/pkg/helper"
)

// CurrencyServerA is a mock server response
func (c *Controller) CurrencyServerA(from, to string) (*model.ExchangeRateServerResponse, error) {
	c.logger.Info().Msg("Currency Server A started...")
	return c.ratesService.CurrencyServerA(from, to)
}

// CurrencyServerB is a mock server response
func (c *Controller) CurrencyServerB(from, to string) (*model.ExchangeRateServerResponse, error) {
	c.logger.Info().Msg("Currency Server B started...")
	return c.ratesService.CurrencyServerB(from, to)
}

// ServerResponse makes a concurrent request to both of the mock server: Controller.CurrencyServerA() and Controller.CurrencyServerB() and returns a response of model.ExchangeRateServerResponse{} pointer
func (c *Controller) ServerResponse(from, to string) (*model.ExchangeRateServerResponse, error) {
	c.logger.Info().Msg("Server Response started...")

	responseChan := make(chan *model.ExchangeRateServerResponse, 2)

	go func() {
		response, err := c.ratesService.CurrencyServerA(from, to)
		responseChan <- response
		if err != nil {
			response, err := c.ratesService.CurrencyServerB(from, to)
			responseChan <- response
			if err != nil {
				c.logger.Error().Msgf("CurrencyServerB ::: error: %v", err)
				responseChan <- &model.ExchangeRateServerResponse{Error: err}
			}
		}
	}()

	go func() {
		response, err := c.ratesService.CurrencyServerB(from, to)
		responseChan <- response
		if err != nil {
			response, err := c.ratesService.CurrencyServerA(from, to)
			responseChan <- response
			if err != nil {
				c.logger.Error().Msgf("CurrencyServerA ::: error: %v", err)
				responseChan <- &model.ExchangeRateServerResponse{Error: err}
			}
		}
	}()

	var successfulResponse *model.ExchangeRateServerResponse
	var err error
	for i := 0; i < 2; i++ {
		response := <-responseChan
		if response.Error == nil {
			successfulResponse = response
			break
		}

		err = helper.ErrBothServicesResponse
	}

	if err != nil {
		c.logger.Error().Msgf("ServerResponse ::: error: %v", err)
		return nil, err
	}

	return successfulResponse, nil
}

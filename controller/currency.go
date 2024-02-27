package controller

import (
	"cadana/model"
	"cadana/pkg/helper"
	"fmt"
)

func (c *Controller) CurrencyServerA(from, to string) (*model.ExchangeRateServerResponse, error) {
	fmt.Println("Server A started...")
	return c.ratesService.CurrencyServerA(from, to)
}

func (c *Controller) CurrencyServerB(from, to string) (*model.ExchangeRateServerResponse, error) {
	fmt.Println("Server B started...")
	return c.ratesService.CurrencyServerB(from, to)
}

func (c *Controller) ServerResponse(from, to string) (*model.ExchangeRateServerResponse, error) {
	fmt.Println("Server Response started...")
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

package exchangerates

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"cadana/controller"
	"cadana/model"
	restModel "cadana/model"
	"cadana/pkg/environment"
)

type exchangeHandler struct {
	logger      zerolog.Logger
	environment *environment.Env
	controller  controller.Operations
}

// New creates a new instance of the exchange rest handler
func New(r *gin.RouterGroup, l zerolog.Logger, env *environment.Env, c controller.Operations) {
	exchange := exchangeHandler{
		logger:      l,
		environment: env,
		controller:  c,
	}

	exchangeGroup := r.Group("/exchange")

	// Endpoints exposed under exchange rates
	exchangeGroup.POST("/rates", exchange.rates())

}

// exchange rates 	godoc
//
//	@Summary		exchange rates
//	@Description	this endpoint gets an exchange rate
//	@Tags			exchange
//	@Accept			json
//	@Produce		json
//	@Param			ExchangeRateRequest			body	model.ExchangeRateReq			true	"exchange rate request"
//	@Success		200
//	@Router			/exchange/rates [post]
func (ex *exchangeHandler) rates() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request model.ExchangeRateReq

		// run the validation first
		if err := c.ShouldBindJSON(&request); err != nil {
			ex.logger.Error().Msgf("%v", err)
			restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		currency, err := request.Validate()
		if err != nil {
			ex.logger.Error().Msgf("%v", err)
			restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		response, err := ex.controller.ServerResponse(currency[0], currency[1])
		if err != nil {
			ex.logger.Error().Msgf("%v", err)
			restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		var responseMap = make(map[string]interface{})
		responseMap[request.CurrencyPair] = response.Rate

		restModel.OkResponse(c, http.StatusOK, "successful currency conversion rate", responseMap)
	}

}

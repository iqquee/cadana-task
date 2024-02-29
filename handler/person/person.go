package person

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"cadana/controller"
	restModel "cadana/model"
	"cadana/person"
	"cadana/pkg/helper"
)

type personHandler struct {
	logger     zerolog.Logger
	controller controller.Operations
}

// New creates a new instance of the exchange rest handler
func New(r *gin.RouterGroup, l zerolog.Logger, c controller.Operations) {
	person := personHandler{
		logger:     l,
		controller: c,
	}

	exchangeGroup := r.Group("/persons")

	// Endpoints exposed under persons data manipulation
	exchangeGroup.GET("/filter/currency/group", person.groupByCurrency())
	exchangeGroup.GET("/filter/currency/:currency", person.filterByCurrency())
	exchangeGroup.GET("/currency/:sortDir", person.filterFromAscToDesc())

}

// groupByCurrency 	godoc
//
//	@Summary		groupByCurrency
//	@Description	this endpoint groups persons by currency
//	@Tags			person
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/persons/filter/currency/group [get]
func (p *personHandler) groupByCurrency() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request person.Persons

		persons, err := request.UnmarshalPersonJSON()
		if err != nil {
			p.logger.Error().Msgf("%v", err)
			restModel.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		response := persons.GroupByCurrency()

		restModel.OkResponse(c, http.StatusOK, "successful", response)
	}
}

// filterByCurrency 	godoc
//
//	@Summary		filterByCurrency
//	@Description	this endpoint converts currencies that are not in USD to USD with the balance converted
//	@Tags			person
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/persons/filter/currency/:currency [get]
func (p *personHandler) filterByCurrency() gin.HandlerFunc {
	return func(c *gin.Context) {
		currency := c.Param("currency")
		if len(currency) == 0 {
			p.logger.Error().Msgf("%v", helper.CustomError("currency is empty"))
			restModel.ErrorResponse(c, http.StatusInternalServerError, helper.CustomError("currency is empty").Error())
			return
		}

		var request person.Persons

		persons, err := request.UnmarshalPersonJSON()
		if err != nil {
			p.logger.Error().Msgf("%v", err)
			restModel.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		response, err := persons.FilterByCurrency(currency, p.controller)
		if err != nil {
			p.logger.Error().Msgf("%v", err)
			restModel.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		restModel.OkResponse(c, http.StatusOK, "successful", response)
	}
}

// filterByCurrency 	godoc
//
//	@Summary		filterByCurrency
//	@Description	this endpoint filters persons salary from ascending to descending or reverse
//	@Tags			person
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/persons/currency/:sortDir [get]
func (p *personHandler) filterFromAscToDesc() gin.HandlerFunc {
	return func(c *gin.Context) {
		sortDir := c.Param("sortDir")
		if len(sortDir) == 0 {
			p.logger.Error().Msgf("%v", helper.CustomError("sort direction is not set"))
			restModel.ErrorResponse(c, http.StatusInternalServerError, helper.CustomError("sort direction is not set").Error())
			return
		}

		var request person.Persons

		persons, err := request.UnmarshalPersonJSON()
		if err != nil {
			p.logger.Error().Msgf("%v", err)
			restModel.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		response, err := persons.FilterFromAscToDesc(sortDir)
		if err != nil {
			p.logger.Error().Msgf("%v", err)
			restModel.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		restModel.OkResponse(c, http.StatusOK, "successful", response)
	}
}

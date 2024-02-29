package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"cadana/controller"
	"cadana/handler/docs"
	exchangerates "cadana/handler/exchangeRates"
	"cadana/handler/person"
	"cadana/pkg/environment"
	"cadana/pkg/helper"
)

const (
	packageName = "handler"
)

// Handler object
type Handler struct {
	logger      *zerolog.Logger
	env         *environment.Env
	api         *gin.RouterGroup
	application *controller.Operations
}

// New creates a new instance of Handler
func New(z zerolog.Logger, ev *environment.Env, engine *gin.Engine, a *controller.Operations) *Handler {
	log := z.With().Str(helper.LogStrPartnerLevel, packageName).Logger()
	apiGroup := engine.Group("/api")
	return &Handler{
		logger:      &log,
		env:         ev,
		api:         apiGroup,
		application: a,
	}
}

// Build setups the APi endpoints
func (h *Handler) Build() {
	v1 := h.api.Group("/v1")
	// register the exchange rates endpoints
	exchangerates.New(v1, *h.logger, h.env, *h.application)

	person.New(v1, *h.logger, *h.application)
	// register the docs endpoint
	docs.New(v1)
}

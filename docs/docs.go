// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/exchange/rates": {
            "post": {
                "description": "this endpoint gets an exchange rate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "exchange"
                ],
                "summary": "exchange rates",
                "parameters": [
                    {
                        "description": "exchange rate request",
                        "name": "ExchangeRateRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ExchangeRateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/persons/filter/currency/:currency": {
            "get": {
                "description": "this endpoint converts currencies that are not in USD to USD with the balance converted",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person"
                ],
                "summary": "filterByCurrency",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/persons/filter/currency/group": {
            "get": {
                "description": "this endpoint groups persons by currency",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person"
                ],
                "summary": "groupByCurrency",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "model.ExchangeRateReq": {
            "type": "object",
            "properties": {
                "currency-pair": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "currency-exchange-g8p8.onrender.com",
	BasePath:         "/api/v1",
	Schemes:          []string{"https"},
	Title:            "Cadana API",
	Description:      "This is the API for Cadana interview task",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

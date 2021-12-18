// Package classification of Product API
//
// Documentation for Product API
//
// Schemes: http
// Host: localhost
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// -application/json
// swagger:meta
package handlers

import (
	"github.com/2001adarsh/Introduction-to-Microservices/data"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// Products is a http.Handler
type Products struct {
	logger    *log.Logger
	validator *data.Validation
}

// A list of Products returns in response
// swagger:response productResponse
type productResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// swagger:parameters deleteProduct
type productIdParameterWrapper struct {
	// The ID of the product to delete from database
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:response noContent
type noContentWrapper struct {
}

// NewProducts creates a products handler with the given logger
func NewProducts(logger *log.Logger, v *data.Validation) *Products {
	return &Products{logger: logger, validator: v}
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

type KeyProduct struct{}

// getProductID returns the product ID from the URL
// Panics if cannot convert the id into an integer this should never happen as the router ensures that this is a valid number
func getProductID(request *http.Request) int {
	// parse the product id from the url
	vars := mux.Vars(request)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		//should never happen
		panic(err)
	}
	return id
}

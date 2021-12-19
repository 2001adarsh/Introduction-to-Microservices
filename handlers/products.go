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

// NewProducts creates a products handler with the given logger
func NewProducts(logger *log.Logger, v *data.Validation) *Products {
	return &Products{logger: logger, validator: v}
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

func (g GenericError) Error() string {
	return g.Message
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

type KeyProduct struct{}

// getProductID returns the product ID from the URL
// Panics if cannot convert the id into an integer this should never happen as the router ensures that this is a valid number
func getProductID(handler *Products, request *http.Request) int {
	// parse the product id from the url
	vars := mux.Vars(request)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		handler.logger.Println("[ERROR] Error while parsing id.", err)
		return -1
	}
	return id
}

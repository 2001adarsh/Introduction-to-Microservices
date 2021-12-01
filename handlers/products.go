package handlers

import (
	"github.com/Intro_to_Microservices/data"
	"log"
	"net/http"
)

// Products is a http.Handler
type Products struct {
	logger *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(logger *log.Logger) *Products {
	return &Products{logger: logger}
}

// ServeHTTP is the main entry point for the handler and satisfies the http.Handler interface
func (handler *Products) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		handler.getProducts(writer, request)
		return
	} else if request.Method == http.MethodPost {
		handler.createProduct(writer, request)
		return
	}

	//catch all
	writer.WriteHeader(http.StatusMethodNotAllowed)
}

// GET REQUEST - getProducts returns the products from the data store
func (handler *Products) getProducts(writer http.ResponseWriter, request *http.Request) {
	listOfProducts := data.GetProducts()
	err := listOfProducts.ToJSON(writer)
	if err != nil {
		http.Error(writer, "Unable to Marshal the data", http.StatusInternalServerError)
	}
}

// POST REQUEST - create a new Product and add to ProductList.
func (handler *Products) createProduct(writer http.ResponseWriter, request *http.Request) {
	handler.logger.Println("Handle POST Product")

	prod := &data.Product{}
	err := prod.FromJSON(request.Body)
	if err != nil {
		http.Error(writer, "Unable to parse body.", http.StatusBadRequest)
	}
	handler.logger.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}

package handlers

import (
	"context"
	"fmt"
	"github.com/2001adarsh/Introduction-to-Microservices/data"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// Products is a http.Handler
type Products struct {
	logger *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(logger *log.Logger) *Products {
	return &Products{logger: logger}
}

// GET REQUEST - getProducts returns the products from the data store
func (handler *Products) GetProducts(writer http.ResponseWriter, request *http.Request) {
	listOfProducts := data.GetProducts()
	err := listOfProducts.ToJSON(writer)
	if err != nil {
		http.Error(writer, "Unable to Marshal the data", http.StatusInternalServerError)
	}
}

// POST REQUEST - create a new Product and add to ProductList.
func (handler *Products) CreateProduct(writer http.ResponseWriter, request *http.Request) {
	handler.logger.Println("Handle POST Product")
	prod := request.Context().Value(KeyProduct{}).(*data.Product)
	handler.logger.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}

// PUT REQUEST - update a product with given id number
func (handler *Products) UpdateProduct(writer http.ResponseWriter, request *http.Request) {
	var header = mux.Vars(request)
	id, err := strconv.Atoi(header["id"])
	if err != nil {
		http.Error(writer, "Unable to get id from request", http.StatusBadRequest)
		return
	}
	handler.logger.Println("Handle PUT Product", id)
	prod := request.Context().Value(KeyProduct{}).(*data.Product)
	handler.logger.Printf("Prod: %#v", prod)
	err = data.UpdateProduct(id, prod)
	if err == data.ErrPositionNotFound {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	} else if err != nil {
		http.Error(writer, "Problems in getting error: "+err.Error(), http.StatusInternalServerError)
	}
}

type KeyProduct struct{}

func (handler *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		prod := &data.Product{}
		err := prod.FromJSON(request.Body)
		if err != nil {
			handler.logger.Println("[ERROR] deserializing product", err)
			http.Error(writer, "Unable to parse body.", http.StatusBadRequest)
			return
		}
		//validate the product
		if err = prod.ProductValidator(); err != nil {
			handler.logger.Println("[ERROR] validator failed", err)
			http.Error(
				writer,
				fmt.Sprintf("json validator failed, check again: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		// add the product to the context
		ctx := context.WithValue(request.Context(), KeyProduct{}, prod)
		request = request.WithContext(ctx)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(writer, request)
	})
}

package handlers

import (
	"github.com/2001adarsh/Introduction-to-Microservices/data"
	"net/http"
)

// CreateProduct handles POST requests to add new products
func (handler *Products) CreateProduct(writer http.ResponseWriter, request *http.Request) {
	handler.logger.Println("[DEBUG] Creating new product")

	//fetch product from context
	prod := request.Context().Value(KeyProduct{}).(*data.Product)

	handler.logger.Printf("[DEBUG] Inserting Product: %#v\n", prod)
	data.AddProduct(prod)
}

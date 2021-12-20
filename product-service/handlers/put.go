package handlers

import (
	"github.com/2001adarsh/Introduction-to-Microservices/product-service/data"
	"net/http"
)

// swagger:route PUT /products products updateProduct
// Update a products details
//
// responses:
//	201: noContent
//  404: errorResponse
//  422: errorValidation
// UpdateProduct handles PUT requests to update products
func (handler *Products) UpdateProduct(writer http.ResponseWriter, request *http.Request) {
	prod := request.Context().Value(KeyProduct{}).(*data.Product)
	handler.logger.Printf("[DEBUG] updating record id:", prod.ID)

	err := data.UpdateProduct(prod)
	if err == data.ErrProductNotFound {
		handler.logger.Println("[ERROR] product not found", err)
		writer.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product not found in database"}, writer)
		return
	}
	//write the no content success header
	writer.WriteHeader(http.StatusNoContent)
}

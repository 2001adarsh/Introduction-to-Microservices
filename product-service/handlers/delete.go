package handlers

import (
	"github.com/2001adarsh/Introduction-to-Microservices/product-service/data"
	"net/http"
)

// swagger:route DELETE /products/{id} products deleteProduct
// deletes a product from the database
// responses:
//	200: noContent
//  404: errorResponse
// 	500: internalServerError
// DeleteProduct handles DELETE request to remove items from database
func (handler *Products) DeleteProduct(writer http.ResponseWriter, request *http.Request) {
	id := getProductID(handler, request)
	var err error
	if id == -1 {
		err = GenericError{
			Message: "Id not parsable from request.",
		}
	} else {
		handler.logger.Println("[DEBUG] deleting record id", id)
		err = data.DeleteProduct(id)
	}

	if err == data.ErrProductNotFound {
		handler.logger.Println("[ERROR] deleting record id does not exist")
		writer.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, writer)
		return
	} else if err != nil {
		handler.logger.Println("[ERROR] deleting records", err)
		writer.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, writer)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}

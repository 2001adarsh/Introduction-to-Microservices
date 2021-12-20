package handlers

import (
	"github.com/2001adarsh/Introduction-to-Microservices/product-service/data"
	"net/http"
)

// swagger:route GET /products products listAllProducts
// Returns all products in DB as a list
// Responses:
// 	200: productsResponse
// 	500: internalServerError

// ListAll handles GET requests and returns all current products
func (handler *Products) ListAll(writer http.ResponseWriter, request *http.Request) {
	handler.logger.Println("[DEBUG] Get all records")
	listOfProducts := data.GetProducts()
	err := data.ToJSON(listOfProducts, writer)
	if err != nil {
		http.Error(writer, "Unable to Marshal the data", http.StatusInternalServerError)
	}
}

// swagger:route GET /products/{id} products singleProduct
// Return a list of products from the database
// responses:
//	200: productResponse
//	404: errorResponse
//  500: internalServerError
// ListSingle handles GET requests
func (handler *Products) ListSingle(writer http.ResponseWriter, request *http.Request) {
	id := getProductID(handler, request)
	var err error
	var prod interface{}
	if id == -1 {
		err = GenericError{Message: "Id not parsable from request."}
	} else {
		handler.logger.Println("[DEBUG] Getting records for id:", id)
		prod, err = data.GetProductByID(id)
	}
	switch err {
	case nil:
	case data.ErrProductNotFound:
		handler.logger.Println("[ERROR] fetching product", err)
		writer.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{
			Message: err.Error(),
		}, writer)
		return
	default:
		handler.logger.Println("[ERROR] fetching product", err)
		writer.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{
			Message: err.Error(),
		}, writer)
		return
	}
	err = data.ToJSON(prod, writer)
	if err != nil {
		// we should never be here, but log the error just in case
		handler.logger.Println("[ERROR] serializing product", err)
	}
}

package handlers

import (
	"context"
	"github.com/2001adarsh/Introduction-to-Microservices/product-service/data"
	"net/http"
)

// MiddlewareValidateProduct validates the product in the request and calls next if ok
func (handler *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		prod := &data.Product{}

		err := data.FromJSON(prod, request.Body)
		if err != nil {
			handler.logger.Println("[ERROR] deserializing product", err)

			writer.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, writer)
			return
		}

		//validate the product
		errs := handler.validator.Validate(prod)
		if len(errs) != 0 {
			handler.logger.Println("[ERROR] validating Product", errs)
			writer.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{errs.Errors()}, writer)
			return
		}

		// add the product to the context
		ctx := context.WithValue(request.Context(), KeyProduct{}, prod)
		request = request.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(writer, request)
	})
}

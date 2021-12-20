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

import "github.com/2001adarsh/Introduction-to-Microservices/product-service/data"

// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handlers

// A list of Products returns in response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// Data structure representing a single product
// swagger:response productResponse
type productResponseWrapper struct {
	// Newly created product
	// in: body
	Body data.Product
}

// swagger:parameters updateProduct createProduct
type productParamsWrapper struct {
	// Product data structure to Update or Create.
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body data.Product
}

// swagger:parameters deleteProduct
type productIdParameterWrapper struct {
	// The ID of the product from database
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:response noContent
type noContentWrapper struct {
}

// swagger:response internalServerError
type internalServerErrorWrapper struct {
}

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

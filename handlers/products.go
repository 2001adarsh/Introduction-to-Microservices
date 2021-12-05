package handlers

import (
	"github.com/Intro_to_Microservices/data"
	"log"
	"net/http"
	"regexp"
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

// ServeHTTP is the main entry point for the handler and satisfies the http.Handler interface
func (handler *Products) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		handler.getProducts(writer, request)
	} else if request.Method == http.MethodPost {
		handler.createProduct(writer, request)
	} else if request.Method == http.MethodPut { //update for data in store.
		id := handler.parseIdFromPath(writer, request)
		handler.updateProduct(id, writer, request)
	} else {
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
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

// PUT REQUEST - update a product with given id number
func (handler *Products) updateProduct(id int, writer http.ResponseWriter, request *http.Request) {
	handler.logger.Println("Handle PUT Product")
	prod := &data.Product{}
	err := prod.FromJSON(request.Body)
	if err != nil {
		http.Error(writer, "Unable to parse body.", http.StatusBadRequest)
	}
	handler.logger.Printf("Prod: %#v", prod)
	err = data.UpdateProduct(id, prod)
	if err == data.ErrPositionNotFound {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	} else if err != nil {
		http.Error(writer, "Problems in getting error: "+err.Error(), http.StatusInternalServerError)
	}
}

func (handler *Products) parseIdFromPath(writer http.ResponseWriter, request *http.Request) int {
	rexp := regexp.MustCompile(`/([0-9]+)`)
	path := request.URL.Path
	grp := rexp.FindAllStringSubmatch(path, -1)
	if len(grp) != 1 && len(grp[0]) != 2 {
		http.Error(writer, "Invalid URI", http.StatusBadRequest)
	}
	idString := grp[0][1]
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(writer, "Invalid URI", http.StatusBadRequest)
	}
	return id
}

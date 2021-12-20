package data

import (
	"fmt"
	"log"
	"time"
)

//Product defines the structure of an API Product
// swagger:meta
type Product struct {
	//id of the product
	// min:1
	ID int `json:"id"`
	// name of the product
	// required: true
	Name string `json:"name" validate:"required"`
	// description of the product
	Description string `json:"description"`
	// cost of the product
	// min: 0
	Price float32 `json:"price" validate:"gt=0"`
	// SKU of the product
	// required: true
	// example: abc-bcd-def
	SKU       string `json:"sku" validate:"required,sku"`
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}

// Products is a collection of Product
type Products []*Product

// GetProducts returns a list of products
func GetProducts() Products {
	return productList
}

func GetProductByID(id int) (*Product, error) {
	productIndex := getProductPosition(id)
	if productIndex == -1 {
		return nil, ErrProductNotFound
	}
	return productList[productIndex], nil
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProduct(p *Product) error {
	pos := getProductPosition(p.ID)
	if pos == -1 {
		return ErrProductNotFound
	}

	//update the product in the DB
	productList[pos] = p
	return nil
}

// DeleteProduct deletes a product from the database
func DeleteProduct(id int) error {
	pos := getProductPosition(id)
	log.Println("id,pos are: ", id, pos)
	if pos == -1 {
		return ErrProductNotFound
	}
	productList = append(productList[:pos], productList[pos+1:]...)
	return nil
}

// ErrProductNotFound is an error raised when a product can not be found in the database
var ErrProductNotFound = fmt.Errorf("Product not found")

func getProductPosition(id int) int {
	for idx, product := range productList {
		if product.ID == id {
			return idx
		}
	}
	return -1
}

func getNextID() int {
	lastProduct := productList[len(productList)-1]
	return lastProduct.ID + 1
}

// productList is a hard coded list of products for this example data source
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc232",
		CreatedOn:   time.Now().String(),
		UpdatedOn:   time.Now().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "efg313",
		CreatedOn:   time.Now().String(),
		UpdatedOn:   time.Now().String(),
	},
}

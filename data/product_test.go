package data

import "testing"

func TestProductValidator(t *testing.T) {
	p := &Product{
		Name:  "adarsh",
		Price: 1.50,
		SKU:   "abc-def-efg",
	}
	err := p.ProductValidator()
	if err != nil {
		t.Fatal("Failed to validation", err)
	}
}

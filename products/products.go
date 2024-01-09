package products

import (
	"encoding/json"
	"fmt"
)

// NewProduct is a factory function to create a new Product.
func NewProduct(name string, productType ProductType, price float64, seller string) *Product {
	return &Product{
		Name:   name,
		Type:   productType,
		Price:  price,
		Seller: seller,
	}
}

type Product struct {
	Name   string      `json:"name"`
	Type   ProductType `json:"type"  validate:"required,oneof=Grocery Electronics Kitchen"`
	Price  float64     `json:"price"`
	Seller string      `json:"seller"`
}

func (p *Product) Discount(discount float64) {
	p.Price = p.Price - (p.Price * discount / 100)
}
func (p *Product) PrintJson() {
	str, err := json.Marshal(p)
	if err != nil {
		fmt.Println("error in marshaling product:", err)
	}
	fmt.Printf("%s\n", str)
}

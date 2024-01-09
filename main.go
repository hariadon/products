// main.go
package main

import (
	"fmt"
	"os"

	"github.com/hariadon/products"
)

func main() {
	product := products.NewProduct("iPhone", products.Electronics, 100.0, "Reliance Digital")
	err := product.PrintJSON(os.Stdout)
	if err != nil {
		fmt.Println("Error printing Product:", err)
	}

	products.ProcessItem(product)

	err = product.PrintJSON(os.Stdout)
	if err != nil {
		fmt.Println("Error printing JSON:", err)
	}
}

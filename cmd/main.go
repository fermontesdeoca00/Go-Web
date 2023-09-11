package main

import (
	"github.com/gin-gonic/gin"
)

var nextID int
var Token = "123456"

type Product struct {
	ID           int     `json:"id,omitempty"`
	Name         string  `json:"name,omitempty"`
	Quantity     int     `json:"quantity,omitempty"`
	Code_Value   string  `json:"code_value,omitempty"`
	IS_Published bool    `json:"is_published"`
	Expiration   string  `json:"expiration,omitempty"`
	Price        float64 `json:"price,omitempty"`
}

type Request struct {
	ID           int     `json:"id,omitempty"`
	Name         string  `json:"name,omitempty"`
	Quantity     int     `json:"quantity,omitempty"`
	Code_Value   string  `json:"code_value,omitempty"`
	IS_Published bool    `json:"is_published"`
	Expiration   string  `json:"expiration,omitempty"`
	Price        float64 `json:"price,omitempty"`
}

func main() {

	server := gin.Default()

	// server.GET("/ping", ping)
	// server.GET("/products", allProducts)
	// server.GET("/products/:id", searchById)
	// server.GET("/products/search", getProductsByPrice)
	// server.POST("/products", addProduct)

	server.Run(":8080")

}

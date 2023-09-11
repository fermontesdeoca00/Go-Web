package main

import (
	"encoding/json"
	"os"

	"github.com/fermontesdeoca00/Go-Web/cmd/server/handler"
	"github.com/fermontesdeoca00/Go-Web/internal/domain"
	"github.com/fermontesdeoca00/Go-Web/internal/product"
	"github.com/gin-gonic/gin"
)

var nextID int
var Token = "123456"

func main() {

	var productsList []domain.Product
	loadProducts("products.json", &productsList)

	repo := product.NewRepository(productsList)
	service := product.NewService(repo)
	productHandler := handler.NewProductHandler(service)

	server := gin.Default()
	server.GET("ping", func(c *gin.Context) { c.String(200, "pong") })
	products := server.Group("/products")
	{
		products.GET("/", productHandler.GetAll())
		products.GET("/:id", productHandler.GetByID())
		products.GET("/search", productHandler.SearchByPrice())
		products.POST("/", productHandler.Post())
	}

	server.Run(":8080")

}

// load all the products from a json file
func loadProducts(path string, list *[]domain.Product) {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(file), &list)
	if err != nil {
		panic(err)
	}
}

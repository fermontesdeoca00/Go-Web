package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	loadJson "github.com/fermontesdeoca00/Go-Web/internal"
	"github.com/gin-gonic/gin"
)

type Product struct {
	ID           int     `json:"id,omitempty"`
	Name         string  `json:"name,omitempty"`
	Quantity     int     `json:"quantity,omitempty"`
	Code_Value   string  `json:"code_value,omitempty"`
	IS_Published bool    `json:"is_published,omitempty"`
	Expiration   string  `json:"expiration,omitempty"`
	Price        float64 `json:"price,omitempty"`
}

var DataSlice []Product

func main() {

	server := gin.Default()

	if err := loadJson.LoadDataFromFile("./products.json", &DataSlice); err != nil {
		fmt.Println("Error loading data", err)
		return
	}

	server.GET("/ping", ping)
	server.GET("/products", allProducts)
	server.GET("/products/:id", searchById)
	server.GET("/products/search", getProductsByPrice)

	server.Run(":8080")

}

func ping(ctxt *gin.Context) {
	ctxt.String(200, " pong")
}

func allProducts(ctx *gin.Context) {
	jsonData, err := json.Marshal(DataSlice)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.String(http.StatusOK, string(jsonData))
}

func searchById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	for _, product := range DataSlice {
		if product.ID == id {
			ctx.JSON(http.StatusOK, product)
			break
		}
	}

}

func getProductsByPrice(ctx *gin.Context) {
	productsList := []Product{}
	priceParam, _ := strconv.ParseFloat(ctx.Query("price"), 64)

	for _, product := range DataSlice {
		if product.Price > priceParam {
			productsList = append(productsList, product)
		}
	}
	ctx.JSON(http.StatusAccepted, productsList)
}

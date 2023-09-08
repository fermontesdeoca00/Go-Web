package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	loadJson "github.com/fermontesdeoca00/Go-Web/internal"
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

var DataSlice []Product

func main() {

	server := gin.Default()

	if err := loadJson.LoadDataFromFile("./products.json", &DataSlice); err != nil {
		fmt.Println("Error loading data", err)
		return
	}

	nextID = len(DataSlice)

	server.GET("/ping", ping)
	server.GET("/products", allProducts)
	server.GET("/products/:id", searchById)
	server.GET("/products/search", getProductsByPrice)
	server.POST("/products", addProduct)

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

func addProduct(ctx *gin.Context) {

	// parse the json request body into a Request struct
	var req Request
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate the requested data
	if req.Name == "" || req.Code_Value == "" || req.Quantity == 0 || req.Expiration == "" || req.Price == 0.0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Can't add empty data"})
		return
	}

	// check if code_value is unique
	for _, product := range DataSlice {
		if product.Code_Value == req.Code_Value {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Can't repeat a code value"})
			return
		}
	}

	//TODO:
	// 1) Los tipos de datos deben coincidir con los definidos en el planteo del problema.
	// 2) La fecha de vencimiento debe tener el formato: XX/XX/XXXX, además debemos verificar que día,
	// mes y año sean valores válidos.

	// generate an unique id
	nextID++
	req.ID = nextID

	// add the new product to the slice of json
	DataSlice = append(DataSlice, Product(req))

	//return the newly generated product as a response
	ctx.JSON(http.StatusCreated, req)

}

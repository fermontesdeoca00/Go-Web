package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/fermontesdeoca00/Go-Web/internal/domain"
	"github.com/fermontesdeoca00/Go-Web/internal/product"
	"github.com/gin-gonic/gin"
)

type productHandler struct {
	s product.Service
}

// NewProductHAndler creates a controller of products
func NewProductHandler(s product.Service) *productHandler {
	return &productHandler{
		s: s,
	}
}

// GetAll returns all products
func (h *productHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		products, _ := h.s.GetAll()
		c.JSON(http.StatusOK, products)
	}
}

// GetByID returns a product by ID
func (h *productHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, errors.New("invalid ID"))
			return
		}
		product, err := h.s.GetByID(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

// SearchByPrice returns a product by price greater than a given price
func (h *productHandler) SearchByPrice() gin.HandlerFunc {
	return func(c *gin.Context) {
		priceParam := c.Query("price")
		price, err := strconv.ParseFloat(priceParam, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, errors.New("invalid price"))
			return
		}
		products, err := h.s.SearchByPrice(price)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, products)
	}
}

// ValidateEmptyFields validates if the fields of the product are not empty
func ValidateEmpty(product *domain.Product) (bool, error) {
	switch {
	case product.Name == "" || product.Code_Value == "" || product.Expiration == "":
		return false, errors.New("fields cant't be empty")
	case product.Quantity <= 0 || product.Price <= 0:
		if product.Quantity <= 0 {
			return false, errors.New("quantity must be greater than 0")
		}
		if product.Price <= 0 {
			return false, errors.New("price must be greater than 0")
		}

	}
	return true, nil
}

// validate expiration date of the product
func ValidateExpirationDate(product *domain.Product) (bool, error) {
	date := strings.Split(product.Expiration, "/")
	list := []int{}
	if len(date) != 3 {
		return false, errors.New("invalid expiration date, must be dd/mm/yyyy")
	}
	for value := range date {
		num, err := strconv.Atoi(date[value])
		if err != nil {
			return false, errors.New("invalid expiration date, must be numbres")
		}
		list = append(list, num)
	}
	condition := (list[0] < 1 || list[0] > 31) && (list[1] < 1 || list[1] > 12) && (list[2] < 1 || list[2] > 9999)
	if condition {
		return false, errors.New("invalid expiration date, date must be between 1 and 31/12/9999")
	}
	return true, nil
}

// Post function that creates a new product with validations
func (h *productHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var product domain.Product
		err := c.ShouldBindJSON(&product)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product"})
			return
		}
		validate, err := ValidateEmpty(&product)
		if !validate {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validate, err = ValidateExpirationDate(&product)
		if !validate {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		p, err := h.s.Create(product)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, p)
	}
}

// Delete function that deletes a product by ID
func (h *productHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
	}
}

// Put function that updates a product by ID
func (h *productHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}
		var product domain.Product
		err = c.ShouldBindJSON(&product)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product"})
			return
		}
		validate, err := ValidateEmpty(&product)
		if !validate {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validate, err = ValidateExpirationDate(&product)
		if !validate {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		p, err := h.s.Update(id, product)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, p)
	}
}

// Patch function that updates a product by ID

func (h *productHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Name         string  `json:"name,omitempty"`
		Quantity     int     `json:"quantity,omitempty"`
		Code_Value   string  `json:"code_value,omitempty"`
		IS_Published bool    `json:"is_published"`
		Expiration   string  `json:"expiration,omitempty"`
		Price        float64 `json:"price,omitempty"`
	}
	return func(c *gin.Context) {
		var r Request
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product"})
			return
		}
		update := domain.Product{
			Name:         r.Name,
			Quantity:     r.Quantity,
			Code_Value:   r.Code_Value,
			IS_Published: r.IS_Published,
			Expiration:   r.Expiration,
			Price:        r.Price,
		}
		if update.Expiration != "" {
			validate, err := ValidateExpirationDate(&update)
			if !validate {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}
		p, err := h.s.Update(id, update)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, p)
	}
}

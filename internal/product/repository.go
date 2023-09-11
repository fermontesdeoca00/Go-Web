package product

import (
	"errors"

	"github.com/fermontesdeoca00/Go-Web/internal/domain"
)

type Repository interface {
	// declaration of the methods that will be implemented in the repository.go file
	GetAll() []domain.Product
	GetByID(id int) (domain.Product, error)
	SearchByPrice(price float64) []domain.Product
	Create(product domain.Product) (domain.Product, error)
}

// declaration of the struct repository that implements the interface Repository
type repository struct {
	// declaration of the struct repository
	list []domain.Product
}

// declaration of the function NewRepository that creates a new repository
func NewRepository(list []domain.Product) Repository {
	return &repository{list}
}

// declaration of the function GetAll that returns all products
func (r *repository) GetAll() []domain.Product {
	return r.list
}

// declaration of the function GetByID that searches for a product by ID
func (r *repository) GetByID(id int) (domain.Product, error) {
	for _, product := range r.list {
		if product.ID == id {
			return product, nil
		}
	}
	return domain.Product{}, errors.New("product not found")
}

// declaration of the function SearchByPrice that searches for a product by price
func (r *repository) SearchByPrice(price float64) []domain.Product {
	var products []domain.Product
	for _, product := range r.list {
		if product.Price <= price {
			products = append(products, product)
		}
	}
	return products
}

// declaration of the function Create that creates a new product
func (r *repository) Create(product domain.Product) (domain.Product, error) {
	product.ID = len(r.list) + 1
	r.list = append(r.list, product)
	return product, nil
}

// declaration of the function ValidateCodeValue that validates the code value of a product
func (r *repository) ValidateCodeValue(codeValue string) bool {
	for _, product := range r.list {
		if product.Code_Value == codeValue {
			return false
		}
	}
	return true
}

package product

import (
	"errors"

	"github.com/fermontesdeoca00/Go-Web/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Product, error)
	GetByID(id int) (domain.Product, error)
	SearchByPrice(pricce float64) ([]domain.Product, error)
	Create(product domain.Product) (domain.Product, error)
}

type service struct {
	r Repository
}

// declaration of the function NewService that creates a new service
// if all of the Service methods are not implemented it gives an error!
func NewService(r Repository) Service {
	return &service{r}
}

// declaration of the function GetAll that returns all products
func (s *service) GetAll() ([]domain.Product, error) {
	return s.r.GetAll(), nil
}

// declaration of the function GetByID that searches for a product by ID
func (s *service) GetByID(id int) (domain.Product, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

// declaration of the function SearchByPrice that searches for a product by price
func (s *service) SearchByPrice(price float64) ([]domain.Product, error) {
	products := s.r.SearchByPrice(price)
	if len(products) == 0 {
		return []domain.Product{}, errors.New("no products found")
	}
	return products, nil
}

// declaration of the function Create that creates a new product
func (s *service) Create(product domain.Product) (domain.Product, error) {
	p, err := s.r.Create(product)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

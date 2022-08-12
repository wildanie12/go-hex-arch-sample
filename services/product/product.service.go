package product

import (
	_entities "github.com/wildanie12/go-hex-arch-sample/entities"
	_productRepository "github.com/wildanie12/go-hex-arch-sample/repositories/product"
)

// Service struct
type Service struct {
	productRepository _productRepository.ProductInterface
}

// New Service instance
func New(productRepository _productRepository.ProductInterface) *Service {
	return &Service{
		productRepository: productRepository,
	}
}

// List of products
func (s Service) List() ([]_entities.Product, error) {
	// s.productRepository.Fil
	products, err := s.productRepository.FindAll()
	if err != nil {
		return []_entities.Product{}, nil
	}

	return products, nil
}

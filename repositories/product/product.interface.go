package repositories

import _entities "github.com/wildanie12/go-hex-arch-sample/entities"

// ProductInterface main declaration
type ProductInterface interface {

	// FindAll product data
	FindAll() ([]_entities.Product, error)

	// Find single product data based on filter
	Find(id int) (_entities.Product, error)

	// Store new product data
	Store(product _entities.Product) (_entities.Product, error)

	// Update existing product data
	Update(product _entities.Product) (_entities.Product, error)

	// Delete existing product data
	Delete(id int) error

}
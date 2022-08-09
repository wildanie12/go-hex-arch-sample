package repositories

import _entities "github.com/wildanie12/go-hex-arch-sample/entities"

// ProductInterface main declaration
type ProductInterface interface {
	
	FindAll() ([]_entities.Product, error)

	Find(id int) (_entities.Product, error)

	Store(product _entities.Product) (_entities.Product, error)

	Update(product _entities.Product) (_entities.Product, error)

	Delete(id int) error

}
package repositories

import (
	_entities "github.com/wildanie12/go-hex-arch-sample/entities"
	"gorm.io/gorm"
)

// Repository product main struct
type Repository struct {
	db *gorm.DB
}

// New product repository
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// FindAll product data
func (repo Repository) FindAll() ([]_entities.Product, error) {
	panic("implemented")
}

// Find single product data based on filter
func (repo Repository) Find(id int) (_entities.Product, error) {
	panic("implemented")
}

// Store new product data
func (repo Repository) Store(product _entities.Product) (_entities.Product, error) {
	panic("implemented")
}

// Update existing product data
func (repo Repository) Update(product _entities.Product) (_entities.Product, error) {
	panic("implemented")
}

// Delete existing product data
func (repo Repository) Delete(id int) error {
	panic("implemented")
}
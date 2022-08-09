package repositories

import (
	_entities "github.com/wildanie12/go-hex-arch-sample/entities"
	"gorm.io/gorm"
)


type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}
func (repo Repository) FindAll() ([]_entities.Product, error) {
	panic("implemented")
}
func (repo Repository) Find(id int) (_entities.Product, error) {
	panic("implemented")
}
func (repo Repository) Store(product _entities.Product) (_entities.Product, error) {
	panic("implemented")
}
func (repo Repository) Update(product _entities.Product) (_entities.Product, error) {
	panic("implemented")
}
func (repo Repository) Delete(id int) error {
	panic("implemented")
}
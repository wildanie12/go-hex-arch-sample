package repositories

import (
	"log"
	"net/http"

	_entities "github.com/wildanie12/go-hex-arch-sample/entities"
	_error "github.com/wildanie12/go-hex-arch-sample/entities/common/error"
	_color "github.com/wildanie12/go-hex-arch-sample/utils/color"
	"gorm.io/gorm"
)

// Repository product main struct
type Repository struct {
	db *gorm.DB
}

// NewMySQL returns product repository using mysql datasource
func NewMySQL(db *gorm.DB) *Repository {
	db = db.Model(&_entities.Product{})
	return &Repository{
		db: db,
	}
}

// FilterName return 
func (repo *Repository) FilterName(name string) {
	repo.db = repo.db.Where("name LIKE ?", "%" + name + "%")
}

// FindAll product data
func (repo *Repository) FindAll() ([]_entities.Product, error) {
	products := []_entities.Product{}
	tx := repo.db.Find(&products)
	if tx.Error != nil {
		log.Println(_color.ThisF("red", "[repository] ❌ mysql product repository error: %v", tx.Error))
		return []_entities.Product{}, _error.Make(
			http.StatusInternalServerError, 
			"data source error",
			"Product Data source error",
		)
	}
	return products, nil
}

// Find single product data based on filter
func (repo *Repository) Find(id int) (_entities.Product, error) {
	panic("implemented")
}

// Store new product data
func (repo *Repository) Store(product _entities.Product) (_entities.Product, error) {
	panic("implemented")
}

// Update existing product data
func (repo *Repository) Update(product _entities.Product) (_entities.Product, error) {
	panic("implemented")
}

// Delete existing product data
func (repo *Repository) Delete(id int) error {
	panic("implemented")
}
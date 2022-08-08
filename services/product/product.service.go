package product

import _models "github.com/wildanie12/go-hex-arch-sample/entities"

// Service struct
type Service struct {

}

// New Service instance
func New() *Service {
	return &Service{}
}

var products = []_models.Product{
	{
		ID: 1,
		Name: "Mouse keren",
		Description: "Mouse dengan presisi 6000 DPI yang bisa dipake buat mouse",
		Price: 30000,
		Category: "Aksesoris Komputer",
		Seller: "Wawan",
		Slug: "mouse-keren",
	},
	{
		ID: 2,
		Name: "Keyboard keren",
		Description: "keyboard disertai dengan pencetan untuk mengetik cocok sekali untuk tukang ketik",
		Price: 150000,
		Category: "Aksesoris Komputer",
		Seller: "Wawan",
		Slug: "keyboard-keren",
	},
	{
		ID: 3,
		Name: "Monitor keren",
		Description: "Monitor yang dapat menampilkan gambar untuk dilihat",
		Price: 1800000,
		Category: "Aksesoris Komputer",
		Seller: "Wawan",
		Slug: "monitor-keren",
	},
}

// List of products
func (s Service) List() ([]_models.Product, error) {
	return products, nil
}

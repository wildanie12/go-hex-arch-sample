package internal

import (
	"github.com/wildanie12/go-hex-arch-sample/internal/service/cart"
	"github.com/wildanie12/go-hex-arch-sample/internal/service/product"
)

// ServiceAPI contains all services.
type ServiceAPI struct {
	CartService    *cart.Service
	ProductService *product.Service
}

// AppPresenter interface holds contracts for valid app presentation layer.
type AppPresenter interface {
	// Start runs presenter for application.
	Start()

	// AttachServiceAPI attaches services to be used by presenter.
	AttachServiceAPI(*ServiceAPI)
}

// NewPresenter creates presenter to be used by application.
func NewPresenter(ap AppPresenter, sa *ServiceAPI) AppPresenter {
	ap.AttachServiceAPI(sa)
	return ap
}

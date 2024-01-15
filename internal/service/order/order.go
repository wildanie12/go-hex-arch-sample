package order

import (
	"context"

	orderRepo "github.com/wildanie12/go-hex-arch-sample/internal/repository/order"
)

// Service order struct.
type Service struct {
	orderRepository orderRepo.Interface
}

// New constructs order service.
func New(ordRepo orderRepo.Interface) *Service {
	return &Service{
		orderRepository: ordRepo,
	}
}

// Checkout existing card on particular user.
func (s *Service) Checkout(ctx context.Context) error {
	return nil
}

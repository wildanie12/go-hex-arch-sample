package handler

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	productSrv "github.com/wildanie12/go-hex-arch-sample/internal/service/product"
)

// ProductHandler main struct.
type ProductHandler struct {
	productService *productSrv.Service
}

// NewProduct constructs instance of ProductHandler.
func NewProduct(ps *productSrv.Service) *ProductHandler {
	return &ProductHandler{productService: ps}
}

// Event handles product changed message received from rabbitmq.
func (ph *ProductHandler) Event() OnIncomingMessage {
	return func(msg amqp.Delivery) {
		fmt.Println(" [x] Received", string(msg.Body))
		msg.Ack(false)
	}
}

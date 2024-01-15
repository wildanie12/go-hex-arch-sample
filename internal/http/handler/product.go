package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ProductHandler main struct.
type ProductHandler struct{}

// NewProduct constructs product handler.
func NewProduct() *ProductHandler {
	return &ProductHandler{}
}

// ListCarts provides list of cart with filter arguments.
func (ph *ProductHandler) ListCarts(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "OK",
		"message": "Success getting list of carts",
		"data": []map[string]any{
			{"id": 1, "name": "Something", "price": 99000},
		},
	})
}

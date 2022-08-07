package handlers

import (
	"context"

	_mappings "github.com/wildanie12/go-hex-arch-sample/app/grpc/mappings"
	_product "github.com/wildanie12/go-hex-arch-sample/app/grpc/models/product"
	_productService "github.com/wildanie12/go-hex-arch-sample/services/product"
)

// ProductHandler main struct
type ProductHandler struct {
	_product.UnimplementedProductServer
	productService *_productService.Service
}

// NewProductHandler instance
func NewProductHandler(service *_productService.Service) *ProductHandler {
	return &ProductHandler{
		productService: service,
	}
}

// List product
func (handler *ProductHandler) List(ctx context.Context, request *_product.Void) (*_product.ProductListResponse, error) {
	products, err := handler.productService.List()
	if err != nil {
		return &_product.ProductListResponse{}, nil
	}
	productsResponse := _mappings.ListProduct(products)
	return &_product.ProductListResponse{
		Success: true,
		Status: "OK",
		Code: 200,
		Data: productsResponse,
	}, nil
}

// Show product
func (handler *ProductHandler) Show(ctx context.Context, request *_product.ShowProductRequest) (*_product.ProductShowResponse, error) {
	panic("not implemented")
}

// Store product
func (handler *ProductHandler) Store(ctx context.Context, request *_product.StoreProductRequest) (*_product.WriteProductResponse, error) {
	panic("not implemented")
}

// Update product
func (handler *ProductHandler) Update(ctx context.Context, request *_product.UpdateProductRequest) (*_product.WriteProductResponse, error) {
	panic("not implemented")
}

// Delete product
func (handler *ProductHandler) Delete(ctx context.Context, request *_product.DeleteProductRequest) (*_product.WriteProductResponse, error) {
	panic("not implemented")
}

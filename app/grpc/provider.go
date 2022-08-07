package grpc

import (
	_handlers "github.com/wildanie12/go-hex-arch-sample/app/grpc/handlers"
	_productService "github.com/wildanie12/go-hex-arch-sample/services/product"
)

type gRPCProvider struct {
	handlers struct {
		product *_handlers.ProductHandler
	}
}

func (rpc *MainGRPC) provide() gRPCProvider {
	provider := gRPCProvider{}

	productService := _productService.New()
	provider.handlers.product = _handlers.NewProductHandler(productService)

	return provider	
}
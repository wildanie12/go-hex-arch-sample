package grpc

import (
	_handlers "github.com/wildanie12/go-hex-arch-sample/app/grpc/handlers"
	_config "github.com/wildanie12/go-hex-arch-sample/config"
	_mysql "github.com/wildanie12/go-hex-arch-sample/database/mysql"
	_productRepository "github.com/wildanie12/go-hex-arch-sample/repositories/product"
	_productService "github.com/wildanie12/go-hex-arch-sample/services/product"
)

type gRPCProvider struct {
	handlers struct {
		product *_handlers.ProductHandler
	}
}

func (rpc *MainGRPC) provide() gRPCProvider {
	provider := gRPCProvider{}

	config := _config.New()
	db := _mysql.New(config)
	productRepository := _productRepository.NewMySQL(db)

	productService := _productService.New(productRepository)
	provider.handlers.product = _handlers.NewProductHandler(productService)

	return provider	
}
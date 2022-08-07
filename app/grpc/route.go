package grpc

import (
	_product "github.com/wildanie12/go-hex-arch-sample/app/grpc/models/product"
)

func (rpc *MainGRPC) defineRoute(provider gRPCProvider) {
	_product.RegisterProductServer(rpc.server, provider.handlers.product)
}
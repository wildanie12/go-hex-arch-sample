package mappings

import (
	_grpcModel "github.com/wildanie12/go-hex-arch-sample/app/grpc/models/product"
	_appModel "github.com/wildanie12/go-hex-arch-sample/models"
)

// ListProduct service to grpc response mapping
func ListProduct(products []_appModel.Product) []*_grpcModel.ProductResponse {
	result := []*_grpcModel.ProductResponse{}
	for _, product := range products {
		result = append(result, &_grpcModel.ProductResponse{
			Id: int32(product.ID),
			Name: product.Name,
			Description: product.Description,
			Price: int32(product.Price),
			Category: product.Category,
			Seller: product.Seller,
			Slug: product.Slug,
		})
	}
	return result	
}
package http

import "github.com/wildanie12/go-hex-arch-sample/internal/http/handler"

func (ah *AppHTTP) initRoutes() {
	// constructs your http handler here...
	prdHdl := handler.NewProduct()

	// define your routes here...
	carts := ah.e.Group("/carts")
	carts.GET("", prdHdl.ListCarts)
}

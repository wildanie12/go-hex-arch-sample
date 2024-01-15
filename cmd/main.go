package main

import (
	"fmt"

	"github.com/wildanie12/go-hex-arch-sample/internal"
	appHttp "github.com/wildanie12/go-hex-arch-sample/internal/http"
	cartSrv "github.com/wildanie12/go-hex-arch-sample/internal/service/cart"
	productSrv "github.com/wildanie12/go-hex-arch-sample/internal/service/product"
)

func main() {
	// -- construct your services here...
	s := internal.ServiceAPI{
		CartService:    cartSrv.New(),
		ProductService: productSrv.New(),
	}

	// -- construct your presenters here...
	appHttp := appHttp.New(config.Get().HTTP.Port)

	// -- assign presenters
	frontendPresenter := internal.NewPresenter(appHttp, &s)

	// execute presenters
	forever := make(chan bool)
	fmt.Println("running enabled presenter...")
	frontendPresenter.Start()

	fmt.Println("App server is started...")
	<-forever
}

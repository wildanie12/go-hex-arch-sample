package main

import (
	"fmt"

	"github.com/wildanie12/go-hex-arch-sample/config"
	"github.com/wildanie12/go-hex-arch-sample/internal"
	appHttp "github.com/wildanie12/go-hex-arch-sample/internal/http"
	orderRepo "github.com/wildanie12/go-hex-arch-sample/internal/repository/order"
	cartSrv "github.com/wildanie12/go-hex-arch-sample/internal/service/cart"
	orderSrv "github.com/wildanie12/go-hex-arch-sample/internal/service/order"
	productSrv "github.com/wildanie12/go-hex-arch-sample/internal/service/product"
	"github.com/wildanie12/go-hex-arch-sample/lib"
)

func main() {
	// -- lib initializations
	di := lib.NewDatabaseInstance()
	di.ConnectOrderDB()
	defer di.CloseOrderDB()

	di.OrderDB.AutoMigrate(&orderRepo.Order{}, &orderRepo.Itemline{})

	// -- construct your repositories here...
	ordRepo := orderRepo.NewMySQL(di)

	// -- construct your services here...
	s := internal.ServiceAPI{
		CartService:    cartSrv.New(),
		ProductService: productSrv.New(),
		OrderService:   orderSrv.New(ordRepo),
	}

	// -- construct your presenters here...
	appHttp := appHttp.New(config.Get().HTTP.Port)

	// -- assign presenters
	frontendPresenter := internal.NewPresenter(appHttp, &s)

	forever := make(chan bool)
	fmt.Println("running enabled presenter...")

	// execute presenters
	frontendPresenter.Start()

	fmt.Println("App server is started...")
	<-forever
}

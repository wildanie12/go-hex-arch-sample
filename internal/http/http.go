package http

import (
	"github.com/labstack/echo/v4"
	"github.com/wildanie12/go-hex-arch-sample/config"
	"github.com/wildanie12/go-hex-arch-sample/internal"
)

// AppHTTP struct.
type AppHTTP struct {
	e        *echo.Echo
	services *internal.ServiceAPI
}

// New contructs HTTP application instance.
func New() *AppHTTP {
	return &AppHTTP{e: echo.New()}
}

// Start HTTP Application Server.
func (ah *AppHTTP) Start() {
	ah.initRoutes()
	ah.e.HideBanner = true
	go func() {
		ah.e.Start(":" + config.Get().HTTP.Port)
	}()
}

// AttachServiceAPI attaches services to be used by presenter.
func (ah *AppHTTP) AttachServiceAPI(s *internal.ServiceAPI) {
	ah.services = s
}

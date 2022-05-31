package infrastructure

import (
	"fmt"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"produx/infrastructure/container"
	"produx/infrastructure/http/echo/handler"
	"produx/infrastructure/http/echo/middleware"
	"produx/infrastructure/route"
)

// Start method is bootstrapping and starting the entire application
func Start(containerInstance container.Container) error {
	e := echo.New()
	config := containerInstance.GetConfig()

	e.HTTPErrorHandler = handler.HTTPErrorHandler

	e.Use(echoMiddleware.Recover())
	e.Use(middleware.Logger(containerInstance))
	e.Use(echoMiddleware.Gzip())

	route.PreparePublicRoutes(e, containerInstance)

	address := fmt.Sprintf("%s:%d", config.Interface, config.Port)
	err := e.Start(address)
	if err != nil {
		return err
	}

	return nil
}

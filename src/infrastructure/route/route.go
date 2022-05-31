package route

import (
	"github.com/labstack/echo/v4"

	"produx/infrastructure/container"
	"produx/infrastructure/controller"
)

// PreparePublicRoutes add the necessary public routes to echo
func PreparePublicRoutes(e *echo.Echo, containerInstance container.Container) {
	e.GET("/comment/:id", func(c echo.Context) error {
		commentController := controller.Comment{}

		return commentController.Get(c, containerInstance)
	})

	e.GET("/comment/list/:target", func(c echo.Context) error {
		commentController := controller.Comment{}

		return commentController.List(c, containerInstance)
	})

	e.POST("/comment", func(c echo.Context) error {
		ctrl := controller.Comment{}

		return ctrl.Create(c, containerInstance)
	})

	// -------------------- USER -------------------------
	e.GET("/user/:id", func(c echo.Context) error {
		userController := controller.User{}

		return userController.Get(c, containerInstance)
	})

	e.PUT("/user/:id", func(c echo.Context) error {
		userController := controller.User{}

		return userController.Update(c, containerInstance)
	})

	e.POST("/user", func(c echo.Context) error {
		userController := controller.User{}

		return userController.Create(c, containerInstance)
	})

	e.GET("/user", func(c echo.Context) error {
		userController := controller.User{}

		return userController.List(c, containerInstance)
	})

	e.DELETE("/user/:id", func(c echo.Context) error {
		userController := controller.User{}

		return userController.Delete(c, containerInstance)
	})

	// -------------------- TARGET -------------------------
	e.GET("/target/:id", func(c echo.Context) error {
		targetController := controller.Target{}

		return targetController.Get(c, containerInstance)
	})

	e.PUT("/target/:id", func(c echo.Context) error {
		targetController := controller.Target{}

		return targetController.Update(c, containerInstance)
	})

	e.POST("/target", func(c echo.Context) error {
		targetController := controller.Target{}

		return targetController.Create(c, containerInstance)
	})

	e.GET("/target", func(c echo.Context) error {
		targetController := controller.Target{}

		return targetController.List(c, containerInstance)
	})

	e.DELETE("/target/:id", func(c echo.Context) error {
		targetController := controller.Target{}

		return targetController.Delete(c, containerInstance)
	})

	// -------------------- APP -------------------------
	e.GET("/app/:id", func(c echo.Context) error {
		appController := controller.App{}

		return appController.Get(c, containerInstance)
	})

	e.PUT("/app/:id", func(c echo.Context) error {
		appController := controller.App{}

		return appController.Update(c, containerInstance)
	})

	e.POST("/app", func(c echo.Context) error {
		appController := controller.App{}

		return appController.Create(c, containerInstance)
	})

	e.GET("/app", func(c echo.Context) error {
		appController := controller.App{}

		return appController.List(c, containerInstance)
	})

	e.DELETE("/app/:id", func(c echo.Context) error {
		appController := controller.App{}

		return appController.Delete(c, containerInstance)
	})

	// -------------------- CATEGORY -------------------------

	e.GET("/category/:id", func(c echo.Context) error {
		catController := controller.Category{}

		return catController.Get(c, containerInstance)
	})

	e.PUT("/category/:id", func(c echo.Context) error {
		catController := controller.Category{}

		return catController.Update(c, containerInstance)
	})

	e.POST("/category", func(c echo.Context) error {
		catController := controller.Category{}

		return catController.Create(c, containerInstance)
	})

	e.GET("/category", func(c echo.Context) error {
		catController := controller.Category{}

		return catController.List(c, containerInstance)
	})

	e.DELETE("/category/:id", func(c echo.Context) error {
		catController := controller.Category{}

		return catController.Delete(c, containerInstance)
	})

	// -------------------- PRODUCT -------------------------

	e.GET("/product/:id", func(c echo.Context) error {
		prodController := controller.Product{}

		return prodController.Get(c, containerInstance)
	})

	e.PUT("/product/:id", func(c echo.Context) error {
		prodController := controller.Product{}

		return prodController.Update(c, containerInstance)
	})

	e.POST("/product", func(c echo.Context) error {
		prodController := controller.Product{}

		return prodController.Create(c, containerInstance)
	})

	e.GET("/product", func(c echo.Context) error {
		prodController := controller.Product{}

		return prodController.List(c, containerInstance)
	})

	e.DELETE("/product/:id", func(c echo.Context) error {
		prodController := controller.Product{}

		return prodController.Delete(c, containerInstance)
	})

	// -------------------- SELLER -------------------------

	e.GET("/seller/:id", func(c echo.Context) error {
		sellerController := controller.Seller{}

		return sellerController.Get(c, containerInstance)
	})

	e.PUT("/seller/:id", func(c echo.Context) error {
		sellerController := controller.Seller{}

		return sellerController.Update(c, containerInstance)
	})

	e.POST("/seller", func(c echo.Context) error {
		sellerController := controller.Seller{}

		return sellerController.Create(c, containerInstance)
	})

	e.GET("/seller", func(c echo.Context) error {
		sellerController := controller.Seller{}

		return sellerController.List(c, containerInstance)
	})

	e.DELETE("/seller/:id", func(c echo.Context) error {
		sellerController := controller.Seller{}

		return sellerController.Delete(c, containerInstance)
	})

	// -------------------- CART -------------------------

}

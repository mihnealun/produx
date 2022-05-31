package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/mindstand/gogm/v2"
	"net/http"
	"produx/domain/entity"
	"produx/domain/seller/dto"
	"produx/infrastructure/container"
	"produx/infrastructure/response"
)

type Seller struct{}

func (a Seller) Get(context echo.Context, c container.Container) error {
	return context.JSON(http.StatusOK, dto.NewSellerResponse(&entity.Seller{}))
}

func (a Seller) Create(context echo.Context, c container.Container) error {
	seller := entity.Seller{
		BaseUUIDNode: gogm.BaseUUIDNode{},
		Name:         context.FormValue("name"),
		Slug:         context.FormValue("slug"),
		Image:        context.FormValue("image"),
		Address:      context.FormValue("address"),
	}

	result := c.GetSellerService().Add(seller)

	return context.JSON(http.StatusOK, dto.NewSellerResponse(result))
}

func (a Seller) Update(context echo.Context, c container.Container) error {
	return context.JSON(http.StatusOK, dto.NewSellerResponse(&entity.Seller{}))
}

func (a Seller) Delete(context echo.Context, c container.Container) error {
	return context.JSON(http.StatusOK, response.NewSuccessResponse("Product deleted."))
}

func (a Seller) List(context echo.Context, c container.Container) error {
	sellers := c.GetSellerService().List()

	return context.JSON(http.StatusOK, dto.NewSellerListResponse(sellers))
}

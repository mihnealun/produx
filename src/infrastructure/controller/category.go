package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/mindstand/gogm/v2"
	"net/http"
	"produx/domain/entity"
	"produx/infrastructure/container"
	"produx/infrastructure/response"
)

type Category struct{}

func (a Category) Get(context echo.Context, c container.Container) error {

	return context.JSON(http.StatusOK, response.NewCategoryResponse(&entity.Category{}))
}

func (a Category) Create(context echo.Context, c container.Container) error {
	category := entity.Category{
		BaseUUIDNode: gogm.BaseUUIDNode{},
		Name:         context.FormValue("name"),
		Slug:         context.FormValue("slug"),
	}

	result := c.GetCategoryService().Add(category)

	return context.JSON(http.StatusOK, response.NewCategoryResponse(result))
}

func (a Category) Update(context echo.Context, c container.Container) error {
	return context.JSON(http.StatusOK, response.NewCategoryResponse(&entity.Category{}))
}

func (a Category) Delete(context echo.Context, c container.Container) error {
	return context.JSON(http.StatusOK, response.NewSuccessResponse("Category deleted."))
}

func (a Category) List(context echo.Context, c container.Container) error {
	categories := c.GetCategoryService().List()

	return context.JSON(http.StatusOK, response.NewCategoryListResponse(categories))
}

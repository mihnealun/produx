package controller

import (
	"fmt"
	"net/http"
	"produx/domain/entity"
	"produx/infrastructure/container"
	"produx/infrastructure/response"

	"github.com/labstack/echo/v4"
	"github.com/mindstand/gogm/v2"
)

type Category struct{}

func (a Category) Get(context echo.Context, c container.Container) error {
	category := c.GetCategoryService().Get(context.Param("id"))
	if len(category.UUID) == 0 {
		return context.JSON(
			http.StatusNotFound,
			response.NewErrorResponse(fmt.Sprintf("category not found %s", context.Param("id"))),
		)
	}
	return context.JSON(http.StatusOK, response.NewCategoryResponse(category))
}

func (a Category) Create(context echo.Context, c container.Container) error {
	cService := c.GetCategoryService()

	// check if category with that name already exists
	oldCategory := cService.GetByName(context.FormValue("name"))
	if oldCategory != nil {
		return context.JSON(
			http.StatusBadRequest,
			response.NewErrorResponse(fmt.Sprintf("category named '%s' already exists (%s)", context.FormValue("name"), oldCategory.UUID)),
		)
	}

	category := entity.Category{
		BaseUUIDNode: gogm.BaseUUIDNode{},
		Name:         context.FormValue("name"),
		Slug:         context.FormValue("slug"),
	}

	// find and link parent category
	if len(context.FormValue("parent")) > 0 {
		parent := cService.Get(context.FormValue("parent"))
		if len(parent.UUID) == 0 {
			return context.JSON(
				http.StatusBadRequest,
				response.NewErrorResponse(fmt.Sprintf("invalid category ID provided for parent %s", context.FormValue("parent"))),
			)
		}

		category.Parent = parent
	}

	result := cService.Add(category)

	return context.JSON(http.StatusOK, response.NewCategoryResponse(result))
}

func (a Category) Update(context echo.Context, c container.Container) error {
	return context.JSON(http.StatusOK, response.NewCategoryResponse(&entity.Category{}))
}

func (a Category) Delete(context echo.Context, c container.Container) error {
	category := c.GetCategoryService().Get(context.Param("id"))
	if len(category.UUID) == 0 {
		return context.JSON(
			http.StatusBadRequest,
			response.NewErrorResponse(fmt.Sprintf("invalid category ID provided %s", context.Param("id"))),
		)
	}

	if !c.GetCategoryService().Delete(category) {
		return context.JSON(
			http.StatusNotModified,
			response.NewErrorResponse(fmt.Sprintf("error deleting category with ID %s", context.Param("id"))),
		)
	}

	return context.JSON(http.StatusOK, response.NewSuccessResponse("Category deleted."))
}

func (a Category) List(context echo.Context, c container.Container) error {
	categories := c.GetCategoryService().List()

	return context.JSON(http.StatusOK, response.NewCategoryListResponse(categories))
}

func (a Category) Products(context echo.Context, c container.Container) error {
	category := c.GetCategoryService().Get(context.Param("id"))

	return context.JSON(http.StatusOK, response.NewProductListResponse(category.Products))
}

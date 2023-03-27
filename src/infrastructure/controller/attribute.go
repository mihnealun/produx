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

type Attribute struct{}

func (a Attribute) Get(context echo.Context, c container.Container) error {
	attribute := c.GetAttributeService().Get(context.Param("id"))
	if len(attribute.UUID) == 0 {
		return context.JSON(
			http.StatusNotFound,
			response.NewErrorResponse(fmt.Sprintf("attribute not found %s", context.Param("id"))),
		)
	}
	return context.JSON(http.StatusOK, response.NewAttributeResponse(attribute))
}

func (a Attribute) Create(context echo.Context, c container.Container) error {
	service := c.GetAttributeService()

	// check if attribute with that name already exists
	oldAttribute := service.GetByName(context.FormValue("name"))
	if oldAttribute != nil {
		return context.JSON(
			http.StatusBadRequest,
			response.NewErrorResponse(fmt.Sprintf("attribute named '%s' already exists (%s)", context.FormValue("name"), oldAttribute.UUID)),
		)
	}

	attribute := entity.Attribute{
		BaseUUIDNode: gogm.BaseUUIDNode{},
		Name:         context.FormValue("name"),
		Slug:         context.FormValue("slug"),
	}

	result := service.Add(attribute)

	return context.JSON(http.StatusOK, response.NewAttributeResponse(result))
}

func (a Attribute) Update(context echo.Context, c container.Container) error {
	return context.JSON(http.StatusOK, response.NewAttributeResponse(&entity.Attribute{}))
}

func (a Attribute) Delete(context echo.Context, c container.Container) error {
	attribute := c.GetAttributeService().Get(context.Param("id"))
	if len(attribute.UUID) == 0 {
		return context.JSON(
			http.StatusBadRequest,
			response.NewErrorResponse(fmt.Sprintf("invalid attribute ID provided %s", context.Param("id"))),
		)
	}

	if !c.GetAttributeService().Delete(*attribute) {
		return context.JSON(
			http.StatusNotModified,
			response.NewErrorResponse(fmt.Sprintf("error deleting attribute with ID %s", context.Param("id"))),
		)
	}

	return context.JSON(http.StatusOK, response.NewSuccessResponse("Attribute deleted."))
}

func (a Attribute) List(context echo.Context, c container.Container) error {
	attributes := c.GetAttributeService().List()

	return context.JSON(http.StatusOK, response.NewAttributeListResponse(attributes))
}

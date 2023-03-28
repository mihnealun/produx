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

type AttributeGroup struct{}

func (a AttributeGroup) Get(context echo.Context, c container.Container) error {
	attributeGroup := c.GetAttributeGroupService().Get(context.Param("id"))
	if len(attributeGroup.UUID) == 0 {
		return context.JSON(
			http.StatusNotFound,
			response.NewErrorResponse(fmt.Sprintf("attribute group not found %s", context.Param("id"))),
		)
	}

	return context.JSON(http.StatusOK, response.NewAttributeGroupResponse(attributeGroup))
}

func (a AttributeGroup) Create(context echo.Context, c container.Container) error {
	service := c.GetAttributeGroupService()

	// check if attribute with that name already exists
	oldAttributeGroup := service.GetByName(context.FormValue("name"))
	if oldAttributeGroup != nil {
		return context.JSON(
			http.StatusBadRequest,
			response.NewErrorResponse(fmt.Sprintf("attribute group named '%s' already exists (%s)", context.FormValue("name"), oldAttributeGroup.UUID)),
		)
	}

	attributeGroup := entity.AttributeGroup{
		BaseUUIDNode: gogm.BaseUUIDNode{},
		Name:         context.FormValue("name"),
		Slug:         context.FormValue("slug"),
	}

	result := service.Add(attributeGroup)

	return context.JSON(http.StatusOK, response.NewAttributeGroupResponse(result))
}

func (a AttributeGroup) Update(context echo.Context, c container.Container) error {
	return context.JSON(http.StatusOK, response.NewAttributeGroupResponse(&entity.AttributeGroup{}))
}

func (a AttributeGroup) Delete(context echo.Context, c container.Container) error {
	attributeGroup := c.GetAttributeGroupService().Get(context.Param("id"))
	if len(attributeGroup.UUID) == 0 {
		return context.JSON(
			http.StatusBadRequest,
			response.NewErrorResponse(fmt.Sprintf("invalid attribute group ID provided %s", context.Param("id"))),
		)
	}

	if !c.GetAttributeGroupService().Delete(*attributeGroup) {
		return context.JSON(
			http.StatusNotModified,
			response.NewErrorResponse(fmt.Sprintf("error deleting attribute group with ID %s", context.Param("id"))),
		)
	}

	return context.JSON(http.StatusOK, response.NewSuccessResponse("Attribute group deleted."))
}

func (a AttributeGroup) List(context echo.Context, c container.Container) error {
	attributeGroups := c.GetAttributeGroupService().List()

	return context.JSON(http.StatusOK, response.NewAttributeGroupListResponse(attributeGroups))
}

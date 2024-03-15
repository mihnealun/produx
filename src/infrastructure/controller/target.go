package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/mindstand/gogm/v2"
	"net/http"
	"produx/domain/entity"
	"produx/infrastructure/container"
	"produx/infrastructure/response"
)

type Target struct{}

func (t Target) Get(context echo.Context, c container.Container) error {
	target := c.GetTargetService().Get(context.Param("id"))

	return context.JSON(http.StatusOK, response.NewTargetResponse(target))
}

func (t Target) Create(context echo.Context, c container.Container) error {
	target := entity.Target{
		BaseUUIDNode: gogm.BaseUUIDNode{},
		Name:         context.FormValue("name"),
		Type:         "video",
		Url:          "active",
	}

	result := c.GetTargetService().Add(target)

	return context.JSON(http.StatusOK, response.NewTargetResponse(result))
}

func (t Target) Update(context echo.Context, c container.Container) error {

	return context.JSON(http.StatusOK, response.NewTargetResponse(&entity.Target{}))
}

func (t Target) Delete(context echo.Context, c container.Container) error {

	return context.JSON(http.StatusOK, response.NewSuccessResponse("Target deleted."))
}

func (t Target) List(context echo.Context, c container.Container) error {
	targets := c.GetTargetService().List()

	return context.JSON(http.StatusOK, response.NewTargetListResponse(targets))
}

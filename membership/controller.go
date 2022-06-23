package membership

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (c Controller) Read(ctx echo.Context) error {
	id := ctx.Param("id")
	res, _ := c.service.Read(id)

	return ctx.JSON(res.Code, res)
}

func (c Controller) Create(ctx echo.Context) error {
	var req CreateRequest
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, CreateResponse{Message: "invalid_request_format"})
	}
	res, _ := c.service.Create(req)

	return ctx.JSON(res.Code, res)
}

func (c Controller) Update(ctx echo.Context) error {
	var req UpdateRequest
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, UpdateResponse{Message: "invalid_request_format"})
	}
	res, _ := c.service.Update(req)
	return ctx.JSON(res.Code, res)
}

func (c Controller) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	res, _ := c.service.Delete(id)
	return ctx.JSON(res.Code, res)
}

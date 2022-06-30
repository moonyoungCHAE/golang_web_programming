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

func (controller Controller) Create(c echo.Context) error {
	var req CreateRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid_request_format")
	}
	res, err := controller.service.Create(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(res.Code, res)
}

func (controller Controller) Update(c echo.Context) error {
	var req UpdateRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid_request_format")
	}
	res, err := controller.service.Update(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(res.Code, res)
}

func (controller Controller) Delete(c echo.Context) error {
	var req string
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid_request_format")
	}
	res, err := controller.service.Delete(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(res.Code, res)
}

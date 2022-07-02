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
	req := CreateRequest{
		c.FormValue("UserName"),
		c.FormValue("MembershipType"),
	}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid_request_format")
	}
	res, _ := controller.service.Create(req)
	return c.JSON(res.Code, res)
}

func (controller Controller) Update(c echo.Context) error {
	req := UpdateRequest{
		c.FormValue("ID"),
		c.FormValue("UserName"),
		c.FormValue("MembershipType"),
	}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid_request_format")
	}
	res, _ := controller.service.Update(req)
	return c.JSON(res.Code, res)
}

func (controller Controller) Delete(c echo.Context) error {
	req := c.Param("id")
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid_request_format")
	}
	res, _ := controller.service.Delete(req)
	return c.JSON(res.Code, res)
}

func (controller Controller) GetByID(c echo.Context) error {
	req := c.Param("id")
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid_request_format")
	}
	res, _ := controller.service.GetByID(req)
	return c.JSON(res.Code, res)
}

func (controller Controller) GetSome(c echo.Context) error {
	offset := c.QueryParam("offset")
	limit := c.QueryParam("limit")
	res, _ := controller.service.GetSome(offset, limit)
	return c.JSON(res.Code, res)
}

func (controller Controller) GetAll(c echo.Context) error {
	req := c.Param("id")
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid_request_format")
	}
	res, _ := controller.service.GetAll(req)
	return c.JSON(res.Code, res)
}

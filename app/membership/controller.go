package membership

import (
	"github.com/labstack/echo"
	"net/http"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (controller Controller) Create(c echo.Context) error {
	var createReq *CreateRequest
	if err := c.Bind(&createReq); err != nil {
		return c.String(http.StatusBadRequest, "invalid_create_request_format")
	}

	resp := controller.service.Create(*createReq)
	resp.Filter(c)
	return c.JSON(resp.Code, resp)
}

func (controller Controller) GetByID(c echo.Context) error {
	id := c.Param("id")
	resp := controller.service.RemoveByID(id)
	resp.Filter(c)
	return c.JSON(resp.Code, resp)
}

func (controller Controller) RemoveByID(c echo.Context) error {
	id := c.Param("id")
	resp := controller.service.RemoveByID(id)
	resp.Filter(c)
	return c.JSON(resp.Code, resp)
}

func (controller Controller) ModifyMember(c echo.Context) error {
	var updatedReq *UpdateRequest
	if err := c.Bind(&updatedReq); err != nil {
		return c.String(http.StatusBadRequest, "invalid_create_request_format")
	}
	resp := controller.service.ModifyMember(*updatedReq)
	resp.Filter(c)
	return c.JSON(resp.Code, resp)
}

func (controller Controller) GetMembers(c echo.Context) error {
	offset := c.Param("offset")
	limit := c.Param("limit")
	resp := controller.service.GetMembers(offset, limit)
	resp.Filter(c)
	return c.JSON(resp.Code, resp)
}

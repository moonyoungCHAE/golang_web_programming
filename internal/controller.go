package internal

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (controller *Controller) Create(c echo.Context) error {
	var req CreateRequest // request와 일치하는 dto struct를 선언
	c.Bind(&req)          // 실제 넘겨받은 request를 struct에 담아줌
	res, err := controller.service.Create(req)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusCreated, res)
}

func (controller *Controller) Read(c echo.Context) error {
	id := c.Param("id")
	res, err := controller.service.GetByID(id)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, res)
}

func (controller *Controller) Update(c echo.Context) error {
	return nil
}

func (controller *Controller) Delete(c echo.Context) error {
	return nil
}

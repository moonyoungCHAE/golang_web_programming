package user

import "github.com/labstack/echo/v4"

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (controller Controller) Login(c echo.Context) error {
	name := c.FormValue("name")
	password := c.FormValue("password")

	res := controller.service.Login(name, password)
	return c.JSON(res.Code, res)
}

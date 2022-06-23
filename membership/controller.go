package membership

import (
	"github.com/labstack/echo/v4"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (c Controller) Read(ctx echo.Context) error {
	return nil
}

func (c Controller) Create(ctx echo.Context) error {
	return nil
}

func (c Controller) Update(ctx echo.Context) error {
	return nil
}

func (c Controller) Delete(ctx echo.Context) error {
	return nil
}

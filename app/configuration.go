package app

import (
	"./membership"
	"github.com/labstack/echo"
)

type Config struct {
	Controller membership.Controller
}

func DefaultConfig() *Config {
	data := map[string]membership.Membership{}
	service := membership.NewService(*membership.NewRepository(data))
	controller := membership.NewController(*service)
	return &Config{
		Controller: *controller,
	}
}

func NewEcho(config Config) *echo.Echo {
	e := echo.New()

	controller := config.Controller

	e.POST("/memberships", controller.Create)
	e.PUT("/memberships", controller.Update)

	return e
}

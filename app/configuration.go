package app

import (
	"github.com/boldfaced7/golang_web_programming/app/membership"
	"github.com/labstack/echo/v4"
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

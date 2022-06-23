package config

import (
	"GolangLivePT01/golang_web_programming/membership"
	"github.com/labstack/echo/v4"
	"sync"
)

var instance *AppConfig
var once sync.Once

type AppConfig struct {
	servicePort string
	echo        *echo.Echo
	group       *echo.Group
	controller  *membership.Controller
}

func GetInstance() *AppConfig {
	once.Do(func() {
		controller := initApplication()
		echo := echo.New()
		group := echo.Group("/api/v2")
		instance = &AppConfig{
			servicePort: "8080",
			echo:        echo,
			group:       group,
			controller:  controller,
		}
	})
	return instance
}

func initApplication() *membership.Controller {
	data := map[string]membership.Membership{}
	service := membership.NewService(*membership.NewRepository(data))
	controller := membership.NewController(*service)
	return controller
}

func (conf *AppConfig) GetServicePort() string {
	return conf.servicePort
}

func (conf *AppConfig) GetEcho() *echo.Echo {
	return conf.echo
}

func (conf *AppConfig) GetGroup() *echo.Group {
	return conf.group
}

func (conf *AppConfig) GetController() membership.Controller {
	return *conf.controller
}

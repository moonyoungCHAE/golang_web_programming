package config

import (
	"github.com/labstack/echo/v4"
	"sync"
)

var instance *AppConfig
var once sync.Once

type AppConfig struct {
	servicePort string
	echo        *echo.Echo
	group       *echo.Group
}

func GetInstance() *AppConfig {
	once.Do(func() {
		echo := echo.New()
		group := echo.Group("/api/v2")
		instance = &AppConfig{
			servicePort: "8080",
			echo:        echo,
			group:       group,
		}
	})
	return instance
}

func (conf *AppConfig) GetServicePort() string {
	return conf.servicePort
}

func (conf *AppConfig) GetGroup() *echo.Group {
	return conf.group
}

func (conf *AppConfig) GetEcho() *echo.Echo {
	return conf.echo
}

package routes

import (
	"GolangLivePT01/golang_web_programming/config"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strings"
)

func InitializeRoutes(cfg *config.AppConfig) *echo.Echo {

	controller := cfg.GetController()
	memberships := cfg.GetGroup().Group("/memberships")

	memberships.Use(middleware.BodyDump(func(c echo.Context, reqBody []byte, resBody []byte) {
		reqStr := strings.Trim(string(reqBody), "\n")
		resStr := strings.Trim(string(resBody), "\n")
		c.Logger().Output().Write([]byte(fmt.Sprintf("RequestBody:[%s]\nResponseBody:[%s]\n\n", reqStr, resStr)))
	}))
	memberships.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "URI:[${uri}], Method:[${method}], ResponseHttpStatusCode:[${status}]\n",
	}))

	memberships.POST("", controller.Create)
	memberships.GET("", controller.ReadAll)
	memberships.GET("/:id", controller.Read)
	memberships.PUT("", controller.Update)
	memberships.DELETE("/:id", controller.Delete)

	return cfg.GetEcho()
}

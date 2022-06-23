package routes

import (
	"GolangLivePT01/golang_web_programming/config"
	"github.com/labstack/echo/v4"
)

func InitializeRoutes(cfg *config.AppConfig) *echo.Echo {

	controller := cfg.GetController()
	memberships := cfg.GetGroup().Group("/memberships")
	memberships.POST("", controller.Create)
	memberships.GET("", controller.ReadAll)
	memberships.GET("/:id", controller.Read)
	memberships.PUT("", controller.Update)
	memberships.DELETE("/:id", controller.Delete)

	return cfg.GetEcho()
}

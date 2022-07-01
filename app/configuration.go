package app

import (
	"github.com/boldfaced7/golang_web_programming/app/membership"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
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
	LoggingMiddleware(e)

	e.POST("/memberships", controller.Create)
	e.PUT("/memberships", controller.Update)
	e.GET("/memberships", controller.GetByID)
	e.DELETE("/memberships", controller.Delete)
	// GET /memberships?offset=1,limit=3
	e.GET("/memberships", controller.GetSome)
	return e
}

func LoggingMiddleware(e *echo.Echo) {
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody []byte, resBody []byte) {
		log.Printf("URI=%v", c.Request().RequestURI)
		log.Printf("Method=%v", c.Request().Method)
		if string(reqBody) != "" {
			log.Printf("Request Body=%v", string(reqBody))
		}
		log.Printf("Status Code=%v", c.Response().Status)
		if string(resBody) != "" {
			log.Printf("Response Body=%v", string(resBody))
		}
	}))
}

package internal

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const _defaultPort = 8080

type Server struct {
	controller Controller
}

func NewDefaultServer() *Server {
	data := map[string]Membership{}
	service := NewService(*NewRepository(data))
	controller := NewController(*service)
	return &Server{
		controller: *controller,
	}
}

func (s *Server) Run() {
	e := echo.New()
	// 관련 doc - https://echo.labstack.com/middleware/body-dump/
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		log.Println("Request URI: ", c.Request().RequestURI)
		log.Println("Request Http Method: ", c.Request().Method)
		if string(reqBody) != "" {
			log.Println("Request Body: ", string(reqBody))
		}
		log.Println("Response Http Status Code: ", c.Response().Status)
		if string(resBody) != "" {
			log.Println("Response Body: ", string(resBody))
		}
	}))
	s.Routes(e)
	log.Fatal(e.Start(fmt.Sprintf(":%d", _defaultPort)))

}

func (s *Server) Routes(e *echo.Echo) {
	g := e.Group("/v1")
	RouteMemberships(g, s.controller)
}

func RouteMemberships(e *echo.Group, c Controller) {
	e.POST("/", c.Create)
	e.GET("/:id", c.Read)
	e.PUT("/:id", c.Update)
	e.DELETE("/:id", c.Delete)

	e.POST("/memberships", c.Create, middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		TargetHeader: "X-My-Request-Header",
	}))
}

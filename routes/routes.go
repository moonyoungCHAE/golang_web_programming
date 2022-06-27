package routes

import (
	"GolangLivePT01/golang_web_programming/membership"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strings"
)

func InitializeRoutes(e *echo.Group) {

	controller := getController()
	memberships := e.Group("/memberships")

	memberships.Use(middleware.BodyDump(func(c echo.Context, reqBody []byte, resBody []byte) {
		uri := c.Request().RequestURI
		method := c.Request().Method
		status := http.StatusText(c.Response().Status)
		reqStr := strings.Trim(string(reqBody), "\n")
		resStr := strings.Trim(string(resBody), "\n")
		c.Logger().Output().Write([]byte(fmt.Sprintf("URI:[%s], Method:[%s], StatusCode:[%s]\n"+
			"RequestBody:[%s]\nResponseBody:[%s]\n\n", uri, method, status, reqStr, resStr)))
	}))
	/*
		memberships.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "URI:[${uri}], Method:[${method}], ResponseHttpStatusCode:[${status}]\n",
		}))
	*/

	memberships.POST("", controller.Create)
	memberships.GET("", controller.ReadAll)
	memberships.GET("/:id", controller.Read)
	memberships.PUT("/:id", controller.Update)
	memberships.DELETE("/:id", controller.Delete)

}

func getController() *membership.Controller {
	data := map[string]membership.Membership{}
	service := membership.NewService(*membership.NewRepository(data))
	controller := membership.NewController(*service)
	return controller
}

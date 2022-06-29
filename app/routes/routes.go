package routes

import (
	"GolangLivePT01/golang_web_programming/app/logo"
	membership2 "GolangLivePT01/golang_web_programming/app/membership"
	user2 "GolangLivePT01/golang_web_programming/app/user"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strings"
)

type Controllers struct {
	MembershipController membership2.Controller
	LogoController       logo.Controller
	UserController       user2.Controller
}

type Middlewares struct {
	Middleware user2.Middleware
}

func InitializeRoutes(e *echo.Group) {

	c := initController()
	userMiddleware := initUserMiddleware().Middleware

	memberships := e.Group("/memberships")
	logo := e.Group("/logo")

	memberships.Use(middleware.BodyDump(func(c echo.Context, reqBody []byte, resBody []byte) {
		uri := c.Request().RequestURI
		method := c.Request().Method
		status := http.StatusText(c.Response().Status)
		reqStr := strings.Trim(string(reqBody), "\n")
		resStr := strings.Trim(string(resBody), "\n")
		c.Logger().Output().Write([]byte(fmt.Sprintf("URI:[%s], Method:[%s], Status:[%s]\n"+
			"RequestBody:[%s]\nResponseBody:[%s]\n\n", uri, method, status, reqStr, resStr)))
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "URI:[${uri}], Method:[${method}], StatusCode:[${status}]\n",
	}))

	jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{Claims: &user2.Claims{}, SigningKey: user2.DefaultSecret})

	membershipCtl := c.MembershipController
	logoCtl := c.LogoController
	userCtl := c.UserController

	memberships.POST("", membershipCtl.Create)
	memberships.GET("", membershipCtl.ReadAll, jwtMiddleware, userMiddleware.ValidateAdmin)
	memberships.GET("/:id", membershipCtl.Read, jwtMiddleware, userMiddleware.ValidateMemberOrAdmin)
	memberships.PUT("/:id", membershipCtl.Update, jwtMiddleware, userMiddleware.ValidateMember)
	memberships.DELETE("/:id", membershipCtl.Delete, jwtMiddleware, userMiddleware.ValidateMember)

	logo.GET("", logoCtl.Get)

	e.POST("/login", userCtl.Login)
}

func initController() *Controllers {
	data := map[string]membership2.Membership{}
	service := membership2.NewService(*membership2.NewRepository(data))
	controller := membership2.NewController(*service)

	return &Controllers{
		MembershipController: *controller,
		LogoController:       *logo.NewController(),
		UserController:       *user2.NewController(*user2.NewService(user2.DefaultSecret, *membership2.NewRepository(data))),
	}
}

func initUserMiddleware() *Middlewares {
	data := map[string]membership2.Membership{}
	return &Middlewares{
		Middleware: *user2.NewMiddleware(*membership2.NewRepository(data)),
	}
}

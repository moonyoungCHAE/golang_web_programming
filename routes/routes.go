package routes

import (
	"GolangLivePT01/golang_web_programming/logo"
	"GolangLivePT01/golang_web_programming/membership"
	"GolangLivePT01/golang_web_programming/user"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strings"
)

type Controllers struct {
	MembershipController membership.Controller
	LogoController       logo.Controller
	UserController       user.Controller
}

type Middlewares struct {
	Middleware user.Middleware
}

func InitializeRoutes(e *echo.Group) {

	c := initController()
	userMiddleware := initUserMiddleware().Middleware

	memberships := e.Group("/memberships")
	logo := e.Group("/logo")

	e.Use(middleware.BodyDump(func(c echo.Context, reqBody []byte, resBody []byte) {
		uri := c.Request().RequestURI
		method := c.Request().Method
		status := http.StatusText(c.Response().Status)
		reqStr := strings.Trim(string(reqBody), "\n")
		resStr := strings.Trim(string(resBody), "\n")
		c.Logger().Output().Write([]byte(fmt.Sprintf("URI:[%s], Method:[%s], StatusCode:[%s]\n"+
			"RequestBody:[%s]\nResponseBody:[%s]\n\n", uri, method, status, reqStr, resStr)))
	}))

	jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{Claims: &user.Claims{}, SigningKey: user.DefaultSecret})

	membershipCtl := c.MembershipController
	logoCtl := c.LogoController
	userCtl := c.UserController

	memberships.POST("", membershipCtl.Create)
	memberships.GET("", membershipCtl.ReadAll, jwtMiddleware, userMiddleware.ValidateAdmin)
	memberships.GET("/:id", membershipCtl.Read, jwtMiddleware, userMiddleware.ValidateAdmin)
	memberships.PUT("/:id", membershipCtl.Update, jwtMiddleware, userMiddleware.ValidateMember)
	memberships.DELETE("/:id", membershipCtl.Delete, jwtMiddleware, userMiddleware.ValidateMember)
	logo.GET("", logoCtl.Get)
	e.POST("/login", userCtl.Login)
}

func initController() *Controllers {
	data := map[string]membership.Membership{}
	service := membership.NewService(*membership.NewRepository(data))
	controller := membership.NewController(*service)

	return &Controllers{
		MembershipController: *controller,
		LogoController:       *logo.NewController(),
		UserController:       *user.NewController(*user.NewService(user.DefaultSecret, *membership.NewRepository(data))),
	}
}

func initUserMiddleware() *Middlewares {
	data := map[string]membership.Membership{}
	return &Middlewares{
		Middleware: *user.NewMiddleware(*membership.NewRepository(data)),
	}
}

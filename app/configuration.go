package app

import (
	"github.com/boldfaced7/golang_web_programming/app/logo"
	"github.com/boldfaced7/golang_web_programming/app/membership"
	"github.com/boldfaced7/golang_web_programming/app/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

type Config struct {
	MembershipController membership.Controller
	LogoController       logo.Controller
	UserController       user.Controller
	UserMiddleware       user.Middleware
}

func DefaultConfig() *Config {
	data := map[string]membership.Membership{}
	membershipRepository := membership.NewRepository(data)
	membershipService := membership.NewService(*membershipRepository)
	membershipController := membership.NewController(*membershipService)
	userService := user.NewService(user.DefaultSecret)

	return &Config{
		MembershipController: *membershipController,
		LogoController:       *logo.NewController(),
		UserController:       *user.NewController(*userService),
		//		UserMiddleware:       *user.NewMiddleware(*membership.NewRepository(data)),
		UserMiddleware: *user.NewMiddleware(*membershipRepository),
	}
}

func NewEcho(config Config) *echo.Echo {
	e := echo.New()

	membershipController := config.MembershipController
	userController := config.UserController

	userMiddleware := config.UserMiddleware
	jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{Claims: &user.Claims{}, SigningKey: user.DefaultSecret})

	e.GET("/memberships/:id", membershipController.GetByID, jwtMiddleware, userMiddleware.ValidateMember)
	e.GET("/memberships", membershipController.GetAll, jwtMiddleware, userMiddleware.ValidateAdmin)
	e.POST("/login", userController.Login)
	e.POST("/memberships", membershipController.Create)
	e.GET("/logo", config.LogoController.Get)

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

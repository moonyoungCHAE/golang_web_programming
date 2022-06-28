package user

import (
	"GolangLivePT01/golang_web_programming/membership"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"log"
)

type Middleware struct {
	membershipRepository membership.Repository
}

func NewMiddleware(membershipRepository membership.Repository) *Middleware {
	return &Middleware{membershipRepository: membershipRepository}
}

func (m Middleware) ValidateAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*Claims)
		if !claims.IsAdmin {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

func (m Middleware) ValidateMember(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*Claims)
		log.Println("request ID :::: ", c.Param("id"))
		log.Println("claims ID :::: ", claims.ID)
		if claims.ID != c.Param("id") {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

func (m Middleware) ValidateMemberOrAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*Claims)
		if claims.IsAdmin {
			return next(c)
		}
		if claims.ID == c.Param("id") {
			return next(c)
		}
		return echo.ErrUnauthorized
	}
}

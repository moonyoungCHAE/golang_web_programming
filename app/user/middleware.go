package user

import (
	"GolangLivePT01/golang_web_programming/app/membership"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
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
		if claims.IsAdmin || claims.ID == c.Param("id") {
			return next(c)
		}
		return echo.ErrUnauthorized
	}
}

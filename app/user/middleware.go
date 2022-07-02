package user

import (
	"github.com/boldfaced7/golang_web_programming/app/membership"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
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
		memberships := m.membershipRepository.GetAll()
		var isMember bool
		for _, mem := range memberships {
			if mem.ID == c.Param("id") {
				isMember = true
			}
		}
		if !isMember {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

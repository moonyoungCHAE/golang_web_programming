package user

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type Claims struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	IsAdmin bool   `json:"is_admin"`
	jwt.StandardClaims
}

func NewClaims(id, name string, isAdmin bool) Claims {
	return Claims{
		id,
		name,
		isAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}
}

func NewMemberClaims(id, name string) Claims {
	return NewClaims(id, name, false)
}

func NewAdminClaims(id, name string) Claims {
	return NewClaims(id, name, true)
}

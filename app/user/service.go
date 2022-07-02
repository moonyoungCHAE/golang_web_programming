package user

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

var DefaultSecret = []byte("secret")

type Service struct {
	secret []byte
}

func NewService(secret []byte) *Service {
	return &Service{secret: secret}
}

func (s Service) Login(name string, password string) LoginResponse {
	if name != password {
		return LoginResponse{Code: http.StatusBadRequest, Message: http.StatusText(http.StatusBadRequest)}
	}
	claims := NewMemberClaims(name)
	if name == "admin" {
		claims = NewAdminClaims(name)
	}
	token, err := s.createToken(claims)
	if err != nil {
		return LoginResponse{Code: http.StatusInternalServerError, Message: http.StatusText(http.StatusInternalServerError)}
	}
	return LoginResponse{Code: http.StatusOK, Message: http.StatusText(http.StatusOK), Token: token}
}

func (s Service) createToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secret)
}

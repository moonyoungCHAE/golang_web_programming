package user

import (
	"GolangLivePT01/golang_web_programming/app/membership"
	"errors"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

var DefaultSecret = []byte("secret")

type Service struct {
	secret     []byte
	repository membership.Repository
}

func NewService(secret []byte, repository membership.Repository) *Service {
	return &Service{
		secret:     secret,
		repository: repository,
	}
}

func (s Service) Login(name, password string) (LoginResponse, error) {

	if strings.Trim(name, " ") == "" || strings.Trim(password, " ") == "" {
		return LoginResponse{Code: http.StatusBadRequest, Message: "name or password is invalid"}, errors.New("name or password is invalid")
	}

	if name != password {
		return LoginResponse{Code: http.StatusBadRequest, Message: "wrong password"}, errors.New("wrong password")
	}

	if s.repository.ReadCountByName(name) <= 0 {
		return LoginResponse{Code: http.StatusBadRequest, Message: "your account is not exists"}, errors.New("your account is not exists")
	}

	id := s.repository.ReadIdByName(name)

	claims := NewMemberClaims(id, name)
	if name == "admin" {
		claims = NewAdminClaims(id, name)
	}

	token, err := s.createToken(claims)
	if err != nil {
		return LoginResponse{Code: http.StatusInternalServerError, Message: http.StatusText(http.StatusInternalServerError)}, err
	}
	return LoginResponse{Code: http.StatusOK, Message: http.StatusText(http.StatusOK), Token: token}, nil
}

func (s Service) createToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secret)
}

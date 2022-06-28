package user

import (
	m "GolangLivePT01/golang_web_programming/membership"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var data = make(map[string]m.Membership)

func createTestMembership(app *m.Service, userName, membershipType string) m.CreateResponse {
	res, _ := app.Create(m.CreateRequest{UserName: userName, MembershipType: membershipType})
	return res
}

func TestLogin(t *testing.T) {

	t.Run("로그인 처리를 합니다.", func(t *testing.T) {
		app := m.NewService(*m.NewRepository(data))
		createTestMembership(app, "test", "naver")

		service := NewService(DefaultSecret, *m.NewRepository(data))
		res, err := service.Login("test", "test")

		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		assert.Equal(t, res.Code, http.StatusOK)
	})

	t.Run("name 를 입력하지 않았을 때 예외 처리한다.", func(t *testing.T) {
		app := m.NewService(*m.NewRepository(data))
		createTestMembership(app, "test", "naver")

		service := NewService(DefaultSecret, *m.NewRepository(data))
		_, err := service.Login("", "test")

		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("name or password is invalid"), err)
	})

	t.Run("password 를 입력하지 않았을 때 예외 처리한다.", func(t *testing.T) {
		app := m.NewService(*m.NewRepository(data))
		createTestMembership(app, "test", "naver")

		service := NewService(DefaultSecret, *m.NewRepository(data))
		_, err := service.Login("test", "")

		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("name or password is invalid"), err)
	})

	t.Run("입력한 id가 존재하지 않을 때 예외 처리한다.", func(t *testing.T) {
		app := m.NewService(*m.NewRepository(data))
		createTestMembership(app, "test", "naver")

		service := NewService(DefaultSecret, *m.NewRepository(data))
		_, err := service.Login("test2", "test2")

		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("your account is not exists"), err)
	})

	t.Run("password 가 일치하지 않을 때 예외 처리한다.", func(t *testing.T) {
		app := m.NewService(*m.NewRepository(data))
		createTestMembership(app, "test", "naver")

		service := NewService(DefaultSecret, *m.NewRepository(data))
		_, err := service.Login("test", "1234")

		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("wrong password"), err)
	})

}

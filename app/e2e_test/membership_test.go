package e2e_test

import (
	"fmt"
	"github.com/boldfaced7/golang_web_programming/app"
	"github.com/boldfaced7/golang_web_programming/app/membership"
	"github.com/gavv/httpexpect"
	"net/http"
	"testing"
)

type loginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func TestUserGetById(t *testing.T) {
	handler := app.NewEcho(*app.DefaultConfig())

	e := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(handler),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	t.Run("멤버십의 주인만 멤버십을 조회할 수 있다.", func(t *testing.T) {
		// Given: 멤버십을 생성한다.
		membershipCreateRequest := e.POST("/memberships").
			WithJSON(membership.CreateRequest{
				UserName:       "jenny",
				MembershipType: "naver",
			}).
			Expect().
			Status(http.StatusCreated).
			JSON().Object()

		// When: 멤버십을 생성한 사용자가 로그인한다.
		userLoginRequest := e.POST("/login").
			WithJSON(loginRequest{
				Name:     fmt.Sprintf("%s", membershipCreateRequest.Value("ID").Raw()),
				Password: fmt.Sprintf("%s", membershipCreateRequest.Value("ID").Raw()),
			}).
			Expect().
			Status(http.StatusOK).
			JSON().Object()

		// Then: 사용자의 멤버십 단건 조회를 할 수 있다.
		e.GET(fmt.Sprintf("/memberships/%s", membershipCreateRequest.Value("ID").Raw())).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", userLoginRequest.Value("token").Raw())).
			Expect().
			Status(http.StatusOK).
			JSON().Object()

	})
}

func TestAdminGetAll(t *testing.T) {
	handler := app.NewEcho(*app.DefaultConfig())

	e := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(handler),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	t.Run("Admin 사용자는 멤버십 전체 조회를 할 수 있다.", func(t *testing.T) {
		// Given: 생성된 멤버십이 존재한다.
		e.POST("/memberships").
			WithJSON(membership.CreateRequest{
				UserName:       "jenny",
				MembershipType: "naver",
			}).
			Expect().
			Status(http.StatusCreated).
			JSON().Object()

		// When: Admin 사용자가 로그인한다.
		adminLoginRequest := e.POST("/login").
			WithJSON(loginRequest{
				Name:     "admin",
				Password: "admin",
			}).
			Expect().
			Status(http.StatusOK).
			JSON().Object()

		// Then: 멤버십 전체 조회를 할 수 있다.

		e.GET("/memberships").
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", adminLoginRequest.Value("token").Raw())).
			Expect().
			Status(http.StatusOK)
	})
}

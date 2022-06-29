package e2e_test

import (
	"GolangLivePT01/golang_web_programming/app/config"
	"GolangLivePT01/golang_web_programming/app/membership"
	"GolangLivePT01/golang_web_programming/app/routes"
	"fmt"
	"github.com/gavv/httpexpect"
	"net/http"
	"testing"
)

func initTestConfig(t *testing.T) *httpexpect.Expect {
	cfg := config.GetInstance()
	routes.InitializeRoutes(cfg.GetGroup())

	e := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(cfg.GetEcho()),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})
	return e
}

func TestReadOneMembership(t *testing.T) {

	e := initTestConfig(t)

	t.Run("멤버십의 주인만 멤버십을 조회할 수 있다", func(t *testing.T) {

		testName := "test1"
		//given: 멤버십을 생성한다
		createResponse := e.POST("/api/v2/memberships").
			WithJSON(membership.CreateRequest{UserName: testName, MembershipType: "naver"}).
			Expect().
			Status(http.StatusCreated).
			JSON().Object()

		//when: 멤버십을 생성한 사용자가 로그인한다
		loginResponse := e.POST("/api/v2/login").
			WithFormField("name", testName).
			WithFormField("password", testName).
			Expect().
			Status(http.StatusOK).
			JSON().Object()

		//then: 사용자의 멤버십 단건 조회를 할 수 있다.
		jwtToken := "Bearer " + loginResponse.Value("token").String().Raw()
		e.GET(fmt.Sprintf("/api/v2/memberships/%s", createResponse.Value("id").Raw())).
			WithHeader("authorization", jwtToken).
			Expect().
			Status(http.StatusOK).
			JSON().Object().
			Value("Membership").Object().
			Value("user_name").Equal(testName)
	})

}

func TestReadAllMembership(t *testing.T) {

	e := initTestConfig(t)

	t.Run("Admin 사용자는 멤버십 전체 조회를 할 수 있다", func(t *testing.T) {

		//given: 생성된 멤버십이 존재한다
		e.POST("/api/v2/memberships").
			WithJSON(membership.CreateRequest{UserName: "test1", MembershipType: "naver"}).
			Expect().
			Status(http.StatusCreated).
			JSON().Object()
		e.POST("/api/v2/memberships").
			WithJSON(membership.CreateRequest{UserName: "admin", MembershipType: "payco"}).
			Expect().
			Status(http.StatusCreated).
			JSON().Object()

		//when: admin 사용자가 로그인한다
		loginResponse := e.POST("/api/v2/login").
			WithFormField("name", "admin").
			WithFormField("password", "admin").
			Expect().
			Status(http.StatusOK).
			JSON().Object()

		//then: 멤버십 전체 조회를 할 수 있다
		jwtToken := "Bearer " + loginResponse.Value("token").String().Raw()
		e.GET(fmt.Sprintf("/api/v2/memberships")).
			WithHeader("authorization", jwtToken).
			Expect().
			Status(http.StatusOK).
			JSON().Object().Value("message").Equal("OK")
	})

}

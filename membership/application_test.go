package membership

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCreateMembership(t *testing.T) {
	t.Run("멤버십을 생성한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)

		log.Println(err)
		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID)
		assert.Equal(t, req.MembershipType, res.MembershipType)
	})

	t.Run("이미 등록된 사용자 이름이 존재할 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		_, _ = app.Create(CreateRequest{"jenny", "naver"})
		_, err := app.Create(CreateRequest{"jenny", "payco"})

		log.Println(err)
		assert.NotNil(t, err)
	})

	t.Run("사용자 이름을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		_, err := app.Create(CreateRequest{"", "payco"})
		log.Println(err)
		assert.NotNil(t, err)
	})

	t.Run("멤버십 타입을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		_, err := app.Create(CreateRequest{"jenny", ""})
		log.Println(err)
		assert.NotNil(t, err)
	})

	t.Run("naver/toss/payco 이외의 타입을 입력한 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		_, err := app.Create(CreateRequest{"jenny", "paybook"})
		log.Println(err)
		assert.NotNil(t, err)
	})
}

var testName, testType = "tester", "toss"

func TestUpdate(t *testing.T) {

	t.Run("membership 정보를 갱신한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		createResponse, _ := app.Create(CreateRequest{testName, testType})

		req := UpdateRequest{createResponse.ID, "ray", "payco"}
		res, err := app.Update(req)
		log.Println(err)

		assert.Nil(t, err)
		assert.Equal(t, req.MembershipType, res.MembershipType)
		assert.Equal(t, req.UserName, res.UserName)
		assert.Equal(t, req.ID, res.ID)
	})

	t.Run("수정하려는 사용자의 이름이 이미 존재하는 사용자 이름이라면 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		createResponse, _ := app.Create(CreateRequest{testName, testType})

		req := UpdateRequest{createResponse.ID, "ray", "payco"}
		_, err := app.Update(req)
		log.Println(err)
		assert.NotNil(t, err)
	})

	t.Run("멤버십 아이디를 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		_, _ = app.Create(CreateRequest{testName, testType})

		req := UpdateRequest{"", "ray", "payco"}
		_, err := app.Update(req)
		log.Println(err)
		assert.NotNil(t, err)
	})

	t.Run("사용자 이름을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		createResponse, _ := app.Create(CreateRequest{testName, testType})

		req := UpdateRequest{createResponse.ID, "", "payco"}
		_, err := app.Update(req)
		log.Println(err)
		assert.NotNil(t, err)
	})

	t.Run("멤버쉽 타입을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		createResponse, _ := app.Create(CreateRequest{testName, testType})

		req := UpdateRequest{createResponse.ID, "ray", ""}
		_, err := app.Update(req)
		log.Println(err)
		assert.NotNil(t, err)
	})

	t.Run("주어진 멤버쉽 타입이 아닌 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		createResponse, _ := app.Create(CreateRequest{testName, testType})

		req := UpdateRequest{createResponse.ID, "ray", "paybook"}
		_, err := app.Update(req)
		log.Println(err)
		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {

	t.Run("멤버십을 삭제한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		createResponse, _ := app.Create(CreateRequest{testName, testType})

		err := app.Delete(createResponse.ID)
		log.Println(err)
		assert.Nil(t, err)
	})

	t.Run("id를 입력하지 않았을 때 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		_, _ = app.Create(CreateRequest{testName, testType})

		err := app.Delete("")
		log.Println(err)
		assert.NotNil(t, err)
	})

	t.Run("입력한 id가 존재하지 않을 때 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		_, _ = app.Create(CreateRequest{testName, testType})
		err := app.Delete("nonexists")
		log.Println(err)

		assert.NotNil(t, err)
	})
}

func TestRead(t *testing.T) {

	t.Run("멤버십을 조회합니다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		createResponse, _ := app.Create(CreateRequest{testName, testType})
		res, err := app.Read(createResponse.ID)
		log.Println(err)

		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		assert.Equal(t, createResponse.ID, res.ID)
	})

	t.Run("id를 입력하지 않았을 때 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		_, _ = app.Create(CreateRequest{testName, testType})

		_, err := app.Read("")
		log.Println(err)
		assert.NotNil(t, err)
	})

	t.Run("입력한 id가 존재하지 않을 때 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		_, _ = app.Create(CreateRequest{testName, testType})

		_, err := app.Read("nonexists")
		log.Println(err)
		assert.NotNil(t, err)
	})

}

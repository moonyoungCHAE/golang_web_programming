package membership

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateMembership(t *testing.T) {
	t.Run("멤버십을 생성한다.", func(t *testing.T) {
		app := NewService(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)

		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID)
		assert.Equal(t, req.MembershipType, res.MembershipType)
	})

	t.Run("이미 등록된 사용자 이름이 존재할 경우 실패한다.", func(t *testing.T) {
		app := NewService(*NewRepository(map[string]Membership{}))
		app.Create(CreateRequest{"jenny", "naver"})
		_, err := app.Create(CreateRequest{"jenny", "payco"})

		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("[create] username is duplicated"), err)
	})

	t.Run("사용자 이름을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		app := NewService(*NewRepository(map[string]Membership{}))
		_, err := app.Create(CreateRequest{"", "payco"})
		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("[create] username or membership-type is not entered"), err)
	})

	t.Run("멤버십 타입을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		app := NewService(*NewRepository(map[string]Membership{}))
		_, err := app.Create(CreateRequest{"jenny", ""})
		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("[create] username or membership-type is not entered"), err)
	})

	t.Run("naver/toss/payco 이외의 타입을 입력한 경우 실패한다.", func(t *testing.T) {
		app := NewService(*NewRepository(map[string]Membership{}))
		_, err := app.Create(CreateRequest{"jenny", "paybook"})
		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("[create] membership type is invalid"), err)
	})
}

var testName, testType = "tester", "toss"

func createTestMembership(app *Service) CreateResponse {
	res, _ := app.Create(CreateRequest{testName, testType})
	return res
}

func TestUpdate(t *testing.T) {

	t.Run("membership 정보를 갱신한다.", func(t *testing.T) {
		app := NewService(*NewRepository(map[string]Membership{}))
		createResponse := createTestMembership(app)

		req := UpdateRequest{createResponse.ID, "ray", "payco"}
		res, err := app.Update(req)

		t.Log(err)
		assert.Nil(t, err)
		assert.Equal(t, req.MembershipType, res.Membership.MembershipType)
		assert.Equal(t, req.UserName, res.Membership.UserName)
		assert.Equal(t, req.ID, res.Membership.ID)
	})

	t.Run("수정하려는 사용자의 이름이 이미 존재하는 사용자 이름이라면 예외 처리한다.", func(t *testing.T) {
		app := NewService(*NewRepository(map[string]Membership{}))
		createResponse := createTestMembership(app)

		req := UpdateRequest{createResponse.ID, "tester", "payco"}
		_, err := app.Update(req)
		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("[update] username is duplicated"), err)
	})

	t.Run("멤버십 아이디를 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewService(*NewRepository(map[string]Membership{}))

		req := UpdateRequest{"", "ray", "payco"}
		_, err := app.Update(req)
		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("[update] ID or username, membership-type is not entered"), err)
	})

	t.Run("사용자 이름을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewService(*NewRepository(map[string]Membership{}))
		createResponse := createTestMembership(app)

		req := UpdateRequest{createResponse.ID, "", "payco"}
		_, err := app.Update(req)

		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("[update] ID or username, membership-type is not entered"), err)
	})

	t.Run("멤버쉽 타입을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewService(*NewRepository(map[string]Membership{}))
		createResponse := createTestMembership(app)

		req := UpdateRequest{createResponse.ID, "ray", ""}
		_, err := app.Update(req)

		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("[update] ID or username, membership-type is not entered"), err)
	})

	t.Run("주어진 멤버쉽 타입이 아닌 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewService(*NewRepository(map[string]Membership{}))
		createResponse := createTestMembership(app)

		req := UpdateRequest{createResponse.ID, "ray", "paybook"}
		_, err := app.Update(req)

		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("[update] membership type is invalid"), err)
	})

	t.Run("없는 ID 수정을 요청시, 예외 처리한다.", func(t *testing.T) {
		app := NewService(*NewRepository(map[string]Membership{}))
		createTestMembership(app)

		req := UpdateRequest{"test", "ray", "naver"}
		_, err := app.Update(req)

		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("[update] ID is not exists"), err)
	})
}

func TestDelete(t *testing.T) {

	t.Run("멤버십을 삭제한다.", func(t *testing.T) {
		app := NewService(*NewRepository(map[string]Membership{}))
		createResponse := createTestMembership(app)

		_, err := app.Delete(createResponse.ID)

		assert.Nil(t, err)
	})

	t.Run("id를 입력하지 않았을 때 예외 처리한다.", func(t *testing.T) {
		app := NewService(*NewRepository(map[string]Membership{}))
		_, err := app.Delete("")

		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("[delete] ID is not entered"), err)
	})

	t.Run("입력한 id가 존재하지 않을 때 예외 처리한다.", func(t *testing.T) {
		app := NewService(*NewRepository(map[string]Membership{}))
		_, err := app.Delete("nonexists")

		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("[delete] ID is invalid (non-exists)"), err)
	})
}

func TestRead(t *testing.T) {

	t.Run("멤버십을 조회합니다.", func(t *testing.T) {
		app := NewService(*NewRepository(map[string]Membership{}))
		createResponse := createTestMembership(app)
		res, err := app.Read(createResponse.ID)

		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		assert.Equal(t, createResponse.ID, res.Membership.ID)
	})

	t.Run("id를 입력하지 않았을 때 예외 처리한다.", func(t *testing.T) {
		app := NewService(*NewRepository(map[string]Membership{}))

		_, err := app.Read("")
		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("[read] ID is not entered"), err)
	})

	t.Run("입력한 id가 존재하지 않을 때 예외 처리한다.", func(t *testing.T) {
		app := NewService(*NewRepository(map[string]Membership{}))

		_, err := app.Read("nonexists")
		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("[read] ID is invalid (non-exists)"), err)
	})

}

func TestReadAll(t *testing.T) {

	t.Run("멤버십 전체를 조회합니다.", func(t *testing.T) {
		app := NewService(*NewRepository(map[string]Membership{}))
		createTestMembership(app)
		res, err := app.ReadAll("", "")

		t.Log(res)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("offset이 존재하지 않을 때 전체 조회됩니다.", func(t *testing.T) {
		app := NewService(*NewRepository(map[string]Membership{}))
		createTestMembership(app)

		res, err := app.ReadAll("", "10")

		t.Log(res)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("limit이 존재하지 않을 때 전체 조회됩니다.", func(t *testing.T) {
		app := NewService(*NewRepository(map[string]Membership{}))
		createTestMembership(app)

		res, err := app.ReadAll("1", "")

		t.Log(res)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("offset 타입변환 오류 발생시 예외 처리", func(t *testing.T) {
		app := NewService(*NewRepository(map[string]Membership{}))
		createTestMembership(app)

		res, err := app.ReadAll("test", "10")

		t.Log(res)
		assert.Error(t, err)
		assert.Equal(t, fmt.Errorf("invalid offset data"), err)

	})

	t.Run("limit 타입변환 오류 발생시 예외 처리", func(t *testing.T) {
		app := NewService(*NewRepository(map[string]Membership{}))
		createTestMembership(app)

		res, err := app.ReadAll("0", "test")

		t.Log(res)
		assert.Error(t, err)
		assert.Equal(t, fmt.Errorf("invalid limit data"), err)
	})
}

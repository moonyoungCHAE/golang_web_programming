package internal

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateMembership(t *testing.T) {
	t.Run("멤버십을 생성한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)
		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID)
		assert.Equal(t, req.MembershipType, res.MembershipType)
	})

	t.Run("이미 등록된 사용자 이름이 존재할 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		app.Create(req)
		_, err := app.Create(req)
		assert.EqualError(t, err, "user name already exists")
	})

	t.Run("사용자 이름을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"", "naver"}
		_, err := app.Create(req)
		assert.EqualError(t, err, "user name is required")
	})

	t.Run("멤버십 타입을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", ""}
		_, err := app.Create(req)
		assert.EqualError(t, err, "membership type is required")
	})

	t.Run("naver/toss/payco 이외의 타입을 입력한 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "facebook"}
		_, err := app.Create(req)
		assert.EqualError(t, err, "invalid membership type")
	})
}

func TestRead(t *testing.T) {
	t.Run("멤버십을 조회한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		res, _ := app.Create(CreateRequest{UserName: "jenny", MembershipType: "naver"})
		readReq := ReadRequest{res.ID}
		readRes, _ := app.Read(readReq)
		assert.Equal(t, res.ID, readRes.ID)
		assert.Equal(t, res.MembershipType, readRes.MembershipType)
	})

	t.Run("id를 입력하지 않았을 때 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := ReadRequest{""}
		_, err := app.Read(req)
		assert.EqualError(t, err, "id is required")
	})

	t.Run("입력한 id가 존재하지 않을 때 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := ReadRequest{"12345"}
		_, err := app.Read(req)
		assert.EqualError(t, err, "id not found")
	})
}

func TestUpdate(t *testing.T) {
	t.Run("internal 정보를 갱신한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		res, _ := app.Create(req)
		updateReq := UpdateRequest{UserName: "jenny", ID: res.ID, MembershipType: "toss"}
		updateRes, _ := app.Update(updateReq)
		assert.Equal(t, updateRes.MembershipType, "toss")
	})

	t.Run("수정하려는 사용자의 이름이 이미 존재하는 사용자 이름이라면 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		app.Create(req)
		req2 := CreateRequest{"sunny", "naver"}
		resp, _ := app.Create(req2)

		updateReq := UpdateRequest{
			UserName:       "jenny",
			ID:             resp.ID,
			MembershipType: "toss",
		}
		_, err := app.Update(updateReq)
		assert.EqualError(t, err, "user name already exists")
	})

	t.Run("멤버십 아이디를 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := UpdateRequest{UserName: "jenny", ID: "", MembershipType: "toss"}
		_, err := app.Update(req)
		assert.EqualError(t, err, "id is required")
	})

	t.Run("사용자 이름을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		id := uuid.New().String()
		req := UpdateRequest{id, "", "toss"}
		_, err := app.Update(req)
		assert.EqualError(t, err, "user name is required")
	})

	t.Run("멤버쉽 타입을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		id := uuid.New().String()
		req := UpdateRequest{id, "jenny", ""}
		_, err := app.Update(req)
		assert.EqualError(t, err, "membership type is required")
	})

	t.Run("주어진 멤버쉽 타입이 아닌 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		id := uuid.New().String()
		req := UpdateRequest{id, "jenny", "facebook"}
		_, err := app.Update(req)
		assert.EqualError(t, err, "invalid membership type")
	})
}

func TestDelete(t *testing.T) {
	t.Run("멤버십을 삭제한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		res, _ := app.Create(req)
		deleteReq := DeleteRequest{ID: res.ID}
		deleteRes, _ := app.Delete(deleteReq)
		assert.Equal(t, res.ID, deleteRes.ID)
	})

	t.Run("id를 입력하지 않았을 때 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{UserName: "jenny", MembershipType: "naver"}
		app.Create(req)
		deleteReq := DeleteRequest{ID: ""}
		_, err := app.Delete(deleteReq)
		assert.EqualError(t, err, "id is required")
	})

	t.Run("입력한 id가 존재하지 않을 때 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := DeleteRequest{ID: "12345"}
		_, err := app.Delete(req)
		assert.EqualError(t, err, "id not found")
	})
}

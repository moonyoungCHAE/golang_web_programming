package membership

import (
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
		res, err := app.Create(req)
		assert.Error(t, err)
		assert.Nil(t, res)
		//assert.NotEmpty(t, res.ID)
		//assert.Equal(t, req.MembershipType, res.MembershipType)
	})

	t.Run("사용자 이름을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("멤버십 타입을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("naver/toss/payco 이외의 타입을 입력한 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "ddd"}
		res, err := app.Create(req)
		assert.Error(t, err)
		assert.Nil(t, res)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("membership 정보를 갱신한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)

		//TODO Update 업데이트 id 수정할 이름,타입
		updReq := UpdateRequest{res.ID, "rusy", "toss"}
		updateRes, err := app.Update(updReq)

		assert.Nil(t, err)
		assert.NotEmpty(t, updateRes.ID)
		assert.Equal(t, updReq.MembershipType, updateRes.MembershipType)
	})

	t.Run("수정하려는 사용자의 이름이 이미 존재하는 사용자 이름이라면 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)

		//TODO Update 업데이트 id 수정할 이름,타입
		updReq := UpdateRequest{res.ID, "jenny", "toss"}
		updateRes, err := app.Update(updReq)

		assert.Error(t, err)
		assert.Nil(t, updateRes)
	})

	t.Run("멤버십 아이디를 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		_, err := app.Create(req)

		//TODO Update 업데이트 id 수정할 이름,타입
		updReq := UpdateRequest{"2", "jenny", "toss"}
		updateRes, err := app.Update(updReq)

		assert.Error(t, err)
		assert.Nil(t, updateRes)
	})

	t.Run("사용자 이름을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)

		//TODO Update 업데이트 id 수정할 이름,타입
		updReq := UpdateRequest{res.ID, "", "toss"}
		updateRes, err := app.Update(updReq)

		assert.Error(t, err)
		assert.Nil(t, updateRes)
	})

	t.Run("멤버쉽 타입을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)

		//TODO Update 업데이트 id 수정할 이름,타입
		updReq := UpdateRequest{res.ID, "", "sdfse"}
		updateRes, err := app.Update(updReq)

		assert.Error(t, err)
		assert.Nil(t, updateRes)

	})

	t.Run("주어진 멤버쉽 타입이 아닌 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)

		//TODO Update 업데이트 id 수정할 이름,타입
		updReq := UpdateRequest{res.ID, "", "sdfse"}
		updateRes, err := app.Update(updReq)

		assert.Error(t, err)
		assert.Nil(t, updateRes)
	})
}

func TestDelete(t *testing.T) {
	t.Run("멤버십을 삭제한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		res, _ := app.Create(req)

		//TODO Update 업데이트 id 수정할 이름,타입
		err := app.Delete(res.ID)
		assert.Nil(t, err)
	})

	t.Run("id를 입력하지 않았을 때 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		res, _ := app.Create(req)
		res.ID = ""
		//TODO Update 업데이트 id 수정할 이름,타입
		err := app.Delete(res.ID)
		assert.Error(t, err)
	})

	t.Run("입력한 id가 존재하지 않을 때 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		res, _ := app.Create(req)

		//TODO Update 업데이트 id 수정할 이름,타입
		_ = app.Delete(res.ID)
		err := app.Delete(res.ID)
		assert.Error(t, err)
	})
}

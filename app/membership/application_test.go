package membership

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateMembership(t *testing.T) {
	t.Run("멤버십을 생성한다.", func(t *testing.T) {
		service := NewService(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		res, err := service.Create(req)
		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID)
		assert.Equal(t, req.MembershipType, res.MembershipType)
	})

	t.Run("이미 등록된 사용자 이름이 존재할 경우 실패한다.", func(t *testing.T) {
		service := NewService(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		service.Create(req)
		req = CreateRequest{"jenny", "payco"}
		_, err := service.Create(req)
		assert.ErrorIs(t, SameNameErr, err)
	})

	t.Run("사용자 이름을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		service := NewService(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"", "naver"}
		_, err := service.Create(req)
		assert.ErrorIs(t, NoNameErr, err)
	})

	t.Run("멤버십 타입을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		service := NewService(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", ""}
		_, err := service.Create(req)
		assert.ErrorIs(t, NoMembershipErr, err)
	})

	t.Run("naver/toss/payco 이외의 타입을 입력한 경우 실패한다.", func(t *testing.T) {
		service := NewService(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "kakao"}
		_, err := service.Create(req)
		assert.ErrorIs(t, WrongMembershipErr, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("membership 정보를 갱신한다.", func(t *testing.T) {
		service := NewService(*NewRepository(map[string]Membership{}))
		creq := CreateRequest{"jenny", "naver"}
		cres, _ := service.Create(creq)
		ureq := UpdateRequest{cres.ID, "jenny", "toss"}
		ures, err := service.Update(ureq)
		assert.Nil(t, err)
		assert.Equal(t, ureq.ID, ures.ID)
		assert.Equal(t, ureq.UserName, ures.UserName)
		assert.Equal(t, ureq.MembershipType, ures.MembershipType)
	})

	t.Run("수정하려는 사용자의 이름이 이미 존재하는 사용자 이름이라면 예외 처리한다.", func(t *testing.T) {
		service := NewService(*NewRepository(map[string]Membership{}))
		creq := CreateRequest{"jenny", "naver"}
		service.Create(creq)
		creq = CreateRequest{"jisoo", "naver"}
		cres, _ := service.Create(creq)
		ureq := UpdateRequest{cres.ID, "jenny", "naver"}
		_, err := service.Update(ureq)
		assert.ErrorIs(t, SameNameErr, err)
	})

	t.Run("멤버십 아이디를 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		service := NewService(*NewRepository(map[string]Membership{}))
		creq := CreateRequest{"jenny", "naver"}
		service.Create(creq)
		ureq := UpdateRequest{"", "jenny", "toss"}
		_, err := service.Update(ureq)
		assert.ErrorIs(t, NoIdErr, err)
	})

	t.Run("사용자 이름을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		service := NewService(*NewRepository(map[string]Membership{}))
		creq := CreateRequest{"jenny", "naver"}
		cres, _ := service.Create(creq)
		ureq := UpdateRequest{cres.ID, "", "toss"}
		_, err := service.Update(ureq)
		assert.ErrorIs(t, NoNameErr, err)
	})

	t.Run("멤버쉽 타입을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		service := NewService(*NewRepository(map[string]Membership{}))
		creq := CreateRequest{"jenny", "naver"}
		cres, _ := service.Create(creq)
		ureq := UpdateRequest{cres.ID, "jenny", ""}
		_, err := service.Update(ureq)
		assert.ErrorIs(t, NoMembershipErr, err)
	})

	t.Run("주어진 멤버쉽 타입이 아닌 경우, 예외 처리한다.", func(t *testing.T) {
		service := NewService(*NewRepository(map[string]Membership{}))
		creq := CreateRequest{"jenny", "naver"}
		cres, _ := service.Create(creq)
		ureq := UpdateRequest{cres.ID, "jenny", "kakao"}
		_, err := service.Update(ureq)
		assert.ErrorIs(t, WrongMembershipErr, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("멤버십을 삭제한다.", func(t *testing.T) {
		service := NewService(*NewRepository(map[string]Membership{}))
		creq := CreateRequest{"jenny", "naver"}
		cres, _ := service.Create(creq)
		_, err := service.Delete(cres.ID)
		assert.Nil(t, err)
	})

	t.Run("id를 입력하지 않았을 때 예외 처리한다.", func(t *testing.T) {
		app := NewService(*NewRepository(map[string]Membership{}))
		creq := CreateRequest{"jenny", "naver"}
		app.Create(creq)
		_, err := app.Delete("")
		assert.ErrorIs(t, NoIdErr, err)
	})

	t.Run("입력한 id가 존재하지 않을 때 예외 처리한다.", func(t *testing.T) {
		app := NewService(*NewRepository(map[string]Membership{}))
		creq := CreateRequest{"jenny", "naver"}
		app.Create(creq)
		_, err := app.Delete("wrong_id")
		assert.ErrorIs(t, WrongIdErr, err)
	})
}

package membership

import (
	"fmt"
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
		_, _ = app.Create(req)
		req = CreateRequest{"jenny", "payco"}
		_, err := app.Create(req)
		assert.Equal(t, fmt.Errorf("same_name"), err)
	})

	t.Run("사용자 이름을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"", "naver"}
		_, err := app.Create(req)
		assert.Equal(t, fmt.Errorf("no_name"), err)
	})

	t.Run("멤버십 타입을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", ""}
		_, err := app.Create(req)
		assert.Equal(t, fmt.Errorf("no_membership"), err)
	})

	t.Run("naver/toss/payco 이외의 타입을 입력한 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "kakao"}
		_, err := app.Create(req)
		assert.Equal(t, fmt.Errorf("wrong_membership"), err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("membership 정보를 갱신한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		creq := CreateRequest{"jenny", "naver"}
		_, _ = app.Create(creq)
		ureq := UpdateRequest{"jenny", "jenny", "toss"}
		res, err := app.Update(ureq)
		assert.Nil(t, err)
		assert.Equal(t, ureq.ID, res.ID)
		assert.Equal(t, ureq.UserName, res.UserName)
		assert.Equal(t, ureq.MembershipType, res.MembershipType)
	})

	t.Run("수정하려는 사용자의 이름이 이미 존재하는 사용자 이름이라면 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		creq := CreateRequest{"jenny", "naver"}
		_, _ = app.Create(creq)
		creq = CreateRequest{"jisoo", "naver"}
		_, _ = app.Create(creq)
		ureq := UpdateRequest{"jisoo", "jenny", "naver"}
		_, err := app.Update(ureq)
		assert.Equal(t, fmt.Errorf("same_name"), err)
	})

	t.Run("멤버십 아이디를 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		creq := CreateRequest{"jenny", "naver"}
		_, _ = app.Create(creq)
		ureq := UpdateRequest{"", "jenny", "toss"}
		_, err := app.Update(ureq)
		assert.Equal(t, fmt.Errorf("no_ID"), err)
	})

	t.Run("사용자 이름을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		creq := CreateRequest{"jenny", "naver"}
		_, _ = app.Create(creq)
		ureq := UpdateRequest{"jenny", "", "toss"}
		_, err := app.Update(ureq)
		assert.Equal(t, fmt.Errorf("no_name"), err)
	})

	t.Run("멤버쉽 타입을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		creq := CreateRequest{"jenny", "naver"}
		_, _ = app.Create(creq)
		ureq := UpdateRequest{"jenny", "jenny", ""}
		_, err := app.Update(ureq)
		assert.Equal(t, fmt.Errorf("no_membership"), err)
	})

	t.Run("주어진 멤버쉽 타입이 아닌 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		creq := CreateRequest{"jenny", "naver"}
		_, _ = app.Create(creq)
		ureq := UpdateRequest{"jenny", "jenny", "kakao"}
		_, err := app.Update(ureq)
		assert.Equal(t, fmt.Errorf("wrong_membership"), err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("멤버십을 삭제한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		creq := CreateRequest{"jenny", "naver"}
		_, _ = app.Create(creq)
		err := app.Delete("jenny")
		assert.Nil(t, err)
	})

	t.Run("id를 입력하지 않았을 때 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		creq := CreateRequest{"jenny", "naver"}
		_, _ = app.Create(creq)
		err := app.Delete("")
		assert.Equal(t, fmt.Errorf("no_ID"), err)
	})

	t.Run("입력한 id가 존재하지 않을 때 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		creq := CreateRequest{"jenny", "naver"}
		_, _ = app.Create(creq)
		err := app.Delete("jisoo")
		assert.Equal(t, fmt.Errorf("wrong_ID"), err)
	})
}

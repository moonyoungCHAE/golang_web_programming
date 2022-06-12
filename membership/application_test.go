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

	})

	t.Run("사용자 이름을 입력하지 않은 경우 실패한다.", func(t *testing.T) {

	})

	t.Run("멤버십 타입을 입력하지 않은 경우 실패한다.", func(t *testing.T) {

	})

	t.Run("naver/toss/payco 이외의 타입을 입력한 경우 실패한다.", func(t *testing.T) {

	})
}

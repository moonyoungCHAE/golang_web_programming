package membership

import (
	"errors"
	"log"
	"sort"
)

type Repository struct {
	data map[string]Membership
}

func NewRepository(data map[string]Membership) *Repository {
	return &Repository{data: data}
}

func (r *Repository) Create(m Membership) (*Membership, error) {
	if isCheck, err := vailEmpty(m); isCheck {
		return nil, err
	}

	if _, ok := r.data[m.UserName]; ok {
		return nil, errors.New("이미 존재하는 사용자입니다.")
	}
	r.data[m.UserName] = m

	return &m, nil
}
func (r *Repository) RemoveByID(id string) (*Membership, error) {

	if isCheck, err := vailEmptyByID(id); isCheck {
		return nil, err
	}

	if _, ok := r.data[id]; !ok {
		return nil, errors.New("멤버쉽 ID가 존재하지 않습니다.")
	}

	var deleteUsetName string
	for _, membership := range r.data {
		if membership.ID == id {
			deleteUsetName = membership.UserName
		}
	}
	delete(r.data, deleteUsetName)
	var member Membership
	member.ID = id
	return &member, nil
}

func (r *Repository) ModifyMember(m Membership) (*Membership, error) {

	if isCheck, err := vailEmpty(m); isCheck {
		return nil, err
	}

	member, ok := r.data[m.UserName]
	if !ok {
		return nil, errors.New("멤버쉽 ID가 존재하지 않습니다.")
	}
	if member.UserName == m.UserName {
		return nil, errors.New("사용자 이름이 존재합니다.")
	}
	member.UserName = m.UserName
	member.MembershipType = m.MembershipType
	r.data[m.UserName] = member
	return &member, nil
}

func VaildType(memType string) bool {
	memTypeList := []string{"naver", "toss", "payco"}
	isType := false
	if contains(memTypeList, memType) {
		isType = true
	}
	log.Println(isType)
	return isType
}
func contains(s []string, searchterm string) bool {
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}

func vailEmpty(m Membership) (bool, error) {

	if len(m.UserName) == 0 {
		return true, errors.New("사용자 이름을 작성해주세요.")
	}
	if !VaildType(m.MembershipType) || len(m.MembershipType) == 0 {
		return true, errors.New("멤버쉽을 확인해주세요.")
	}
	return false, nil
}

func vailEmptyByID(id string) (bool, error) {
	if len(id) == 0 {
		return true, errors.New("ID를 작성해주세요.")
	}
	return false, nil
}

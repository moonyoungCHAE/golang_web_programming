package membership

import (
	"errors"
	"sort"
)

type Application struct {
	repository Repository
}

func NewApplication(repository Repository) *Application {
	return &Application{repository: repository}
}

func VaildType(memType string) bool {
	memTypeList := []string{"naver", "toss", "payco"}
	isType := false
	if contains(memTypeList, memType) {
		isType = true
	}
	return isType
}
func contains(s []string, searchterm string) bool {
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}
func (app *Application) Create(request CreateRequest) (*CreateResponse, error) {
	// TODO 1. 입력 값 검증
	if len(request.UserName) == 0 {
		return nil, errors.New("사용자 이름을 작성해주세요.")
	}
	if len(request.MembershipType) == 0 && VaildType(request.MembershipType) {
		return nil, errors.New("멤버쉽을 확인해주세요.")
	}

	// TODO 3. 멤버쉽 존재 여부
	for _, element := range app.repository.data {
		if element.UserName == request.UserName {
			return nil, errors.New("이미 존재하는 사용자입니다.")
		}
	}
	// TODO 2. 사용자 등록 "ID , 사용자 이름, 멤버쉽 타입"
	createMemberMap := map[string]Membership{"1": {"1", request.UserName, request.MembershipType}}
	app.repository.data = createMemberMap

	// TODO 4. 사용자 생성 정보 리턴
	return &CreateResponse{createMemberMap["1"].ID, createMemberMap["1"].MembershipType}, nil
}
func (app *Application) Update(request UpdateRequest) (*UpdateResponse, error) {
	//TODO 1. ID 기준으로 사용자 갱신
	if len(request.UserName) == 0 {
		return nil, errors.New("사용자 이름을 작성해주세요.")
	}
	if len(request.MembershipType) == 0 && VaildType(request.MembershipType) {
		return nil, errors.New("멤버쉽을 확인해주세요.")
	}

	member, ok := app.repository.data[request.ID]
	if !ok {
		return nil, errors.New("멤버쉽 ID가 존재하지 않습니다.")
	}

	if member.UserName == request.UserName {
		return nil, errors.New("사용자 이름이 존재합니다.")
	}

	member.UserName = request.UserName
	member.MembershipType = request.MembershipType
	app.repository.data[request.ID] = member

	return &UpdateResponse{request.ID, member.UserName, member.MembershipType}, nil
}

func (app *Application) Delete(id string) error {
	if len(id) == 0 {
		return errors.New("사용자 이름을 작성해주세요.")
	}
	_, ok := app.repository.data[id]
	if !ok {
		return errors.New("멤버쉽 ID가 존재하지 않습니다.")
	}

	delete(app.repository.data, id)

	return nil
}

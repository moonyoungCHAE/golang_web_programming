package membership

import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"strings"
)

type Application struct {
	repository Repository
}

func NewApplication(repository Repository) *Application {
	return &Application{repository: repository}
}

const validTypes string = "naver/toss/payco"

func (app *Application) Create(request CreateRequest) (CreateResponse, error) {

	// ID 생성
	randomId, err := uuid.NewGen().NewV4()
	if err != nil {
		panic(err)
	}
	userId := randomId.String()

	// 파라미터 검증
	if request.UserName == "" || request.MembershipType == "" {
		return CreateResponse{}, errors.New("[create] username or membership-type is not entered")
	}

	// 멤버십 타입 검증
	if !strings.Contains(validTypes, request.MembershipType) {
		return CreateResponse{}, errors.New("[create] membership type is invalid")
	}

	// 중복 확인
	if isDuplicateName(app.repository.data, request.UserName) {
		return CreateResponse{}, errors.New("[create] username is duplicated")
	}

	// Memory DB에 추가
	app.repository.data[userId] = Membership{
		userId, request.UserName, request.MembershipType,
	}

	return CreateResponse{userId, request.MembershipType}, nil
}

func (app *Application) Update(request UpdateRequest) (UpdateResponse, error) {

	// 파라미터 검증
	if request.ID == "" || request.UserName == "" || request.MembershipType == "" {
		return UpdateResponse{}, errors.New("[update] ID or username, membership-type is not entered")
	}

	// 멤버십 타입 검증
	if !strings.Contains(validTypes, request.MembershipType) {
		return UpdateResponse{}, errors.New("[update] membership type is invalid")
	}

	// 중복 확인
	if isDuplicateName(app.repository.data, request.UserName) {
		return UpdateResponse{}, errors.New("[update] username is duplicated")
	}

	// Memory DB에 수정
	app.repository.data[request.ID] = Membership{ID: request.ID, UserName: request.UserName, MembershipType: request.MembershipType}

	res, exists := app.repository.data[request.ID]
	if !exists {
		return UpdateResponse{}, errors.New("[update] after update, key is invalid")
	}
	fmt.Println(res)

	return UpdateResponse{res.ID, res.UserName, res.MembershipType}, nil
}

func (app *Application) Delete(id string) error {

	// 파라미터 검증
	if id == "" {
		return errors.New("[delete] ID is not entered")
	}

	// id 존재 여부 확인
	if _, exists := app.repository.data[id]; !exists {
		return errors.New("[delete] ID is invalid (non-exists)")
	}

	// Memory DB에 삭제
	delete(app.repository.data, id)

	return nil
}

func (app *Application) Read(id string) (ReadResponse, error) {

	// 파라미터 검증
	if id == "" {
		return ReadResponse{}, errors.New("[read] ID is not entered")
	}

	var membership, exists = app.repository.data[id]

	// id 존재 여부 확인
	if !exists {
		return ReadResponse{}, errors.New("[read] ID is invalid (non-exists)")
	}

	return ReadResponse{membership.ID, membership.UserName, membership.MembershipType}, nil
}

// isDuplicateName returns a bool value whether if username is duplicated or not
func isDuplicateName(data map[string]Membership, userName string) bool {
	for _, membership := range data {
		if membership.UserName == userName {
			return true
		}
	}
	return false
}

func (app *Application) Update(request UpdateRequest) (UpdateResponse, error) {
	return UpdateResponse{}, nil
}

func (app *Application) Delete(id string) error {
	return nil
}

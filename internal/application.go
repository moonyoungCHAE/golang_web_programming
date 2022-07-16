package internal

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrUserAlreadyExists        = errors.New("user already exists")
	ErrUserNameIsRequired       = errors.New("user name is required")
	ErrMembershipTypeIsRequired = errors.New("membership type is required")
	ErrInvalidMembershipType    = errors.New("invalid membership type")

	ErrUserIDNotFound = errors.New("user id not found")
)

type Application struct {
	repository Repository
}

func NewApplication(repository Repository) *Application {
	return &Application{repository: repository}
}

func (app *Application) Create(request CreateRequest) (CreateResponse, error) {
	if _, ok := app.repository.data[request.UserName]; ok {
		return CreateResponse{}, ErrUserAlreadyExists
	}
	if request.UserName == "" {
		return CreateResponse{}, ErrUserNameIsRequired
	}
	if request.MembershipType == "" {
		return CreateResponse{}, ErrMembershipTypeIsRequired
	}
	if request.MembershipType != "naver" && request.MembershipType != "toss" && request.MembershipType != "payco" {
		return CreateResponse{}, ErrInvalidMembershipType
	}

	app.repository.data[request.UserName] = Membership{
		ID:             uuid.New().String(),
		UserName:       request.UserName,
		MembershipType: request.MembershipType,
	}
	return CreateResponse{app.repository.data[request.UserName].ID, app.repository.data[request.UserName].MembershipType}, nil
}

func (app *Application) Read(request ReadRequest) (ReadResponse, error) {
	if request.ID == "" {
		return ReadResponse{}, ErrUserNameIsRequired
	}
	if _, ok := app.repository.data[request.ID]; !ok {
		return ReadResponse{}, ErrUserIDNotFound
	}

	return ReadResponse{app.repository.data[request.ID].ID, app.repository.data[request.ID].UserName, app.repository.data[request.ID].MembershipType}, nil
}

func (app *Application) Update(request UpdateRequest) (UpdateResponse, error) {
	if _, ok := app.repository.data[request.ID]; !ok {
		return UpdateResponse{}, ErrUserIDNotFound
	}
	if request.UserName == "" {
		return UpdateResponse{}, ErrUserNameIsRequired
	}
	if request.MembershipType == "" {
		return UpdateResponse{}, ErrMembershipTypeIsRequired
	}
	if request.MembershipType != "naver" && request.MembershipType != "toss" && request.MembershipType != "payco" {
		return UpdateResponse{}, ErrInvalidMembershipType
	}

	app.repository.data[request.ID] = Membership{
		ID:             request.ID,
		UserName:       request.UserName,
		MembershipType: request.MembershipType,
	}
	return UpdateResponse{ID: app.repository.data[request.ID].ID, UserName: app.repository.data[request.ID].UserName, MembershipType: app.repository.data[request.ID].MembershipType}, nil
}

func (app *Application) Delete(request DeleteRequest) (DeleteResponse, error) {
	if request.ID == "" {
		return DeleteResponse{}, ErrUserNameIsRequired
	}
	if _, ok := app.repository.data[request.ID]; !ok {
		return DeleteResponse{}, ErrUserIDNotFound
	}
	res := app.repository.data[request.ID]
	delete(app.repository.data, request.ID)

	return DeleteResponse{ID: res.ID, UserName: res.UserName, MembershipType: res.MembershipType}, nil
}

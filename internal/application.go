package internal

import (
	"errors"
)

var (
	ErrUserAlreadyExists        = errors.New("user already exists")
	ErrUserNameIsRequired       = errors.New("user name is required")
	ErrMembershipTypeIsRequired = errors.New("membership type is required")
	ErrInvalidMembershipType    = errors.New("invalid membership type")
	ErrUserIDNotFound           = errors.New("user id not found")
	ErrUserIDIsRequired         = errors.New("user id is required")
)

type Application struct {
	repository Repository
}

func NewApplication(repository Repository) *Application {
	return &Application{repository: repository}
}

func (app *Application) Create(request CreateRequest) (CreateResponse, error) {
	if request.UserName == "" {
		return CreateResponse{}, ErrUserNameIsRequired
	}
	if request.MembershipType == "" {
		return CreateResponse{}, ErrMembershipTypeIsRequired
	}
	if request.MembershipType != "naver" && request.MembershipType != "toss" && request.MembershipType != "payco" {
		return CreateResponse{}, ErrInvalidMembershipType
	}
	if app.repository.UserNameExists(request.UserName) {
		return CreateResponse{}, ErrUserAlreadyExists
	}

	return app.repository.Create(request)
}

func (app *Application) Read(request ReadRequest) (ReadResponse, error) {
	if request.ID == "" {
		return ReadResponse{}, ErrUserIDIsRequired
	}
	if _, ok := app.repository.GetByID(request.ID); !ok {
		return ReadResponse{}, ErrUserIDNotFound
	}

	return ReadResponse{app.repository.data[request.ID].ID, app.repository.data[request.ID].UserName, app.repository.data[request.ID].MembershipType}, nil
}

func (app *Application) Update(request UpdateRequest) (UpdateResponse, error) {
	if request.MembershipType == "" {
		return UpdateResponse{}, ErrMembershipTypeIsRequired
	}
	if request.UserName == "" {
		return UpdateResponse{}, ErrUserNameIsRequired
	}
	if request.MembershipType != "naver" && request.MembershipType != "toss" && request.MembershipType != "payco" {
		return UpdateResponse{}, ErrInvalidMembershipType
	}
	if app.repository.UserNameExists(request.UserName) {
		return UpdateResponse{}, ErrUserAlreadyExists
	}

	return app.repository.Update(request)
}

func (app *Application) Delete(request DeleteRequest) (DeleteResponse, error) {
	if request.ID == "" {
		return DeleteResponse{}, ErrUserIDIsRequired
	}
	if _, ok := app.repository.GetByID(request.ID); !ok {
		return DeleteResponse{}, ErrUserIDNotFound
	}

	return app.repository.Delete(request)
}

package membership

import (
	"errors"
	"github.com/gofrs/uuid"
	"net/http"
	"strings"
)

var validMemberships = [3]string{"toss", "naver", "payco"}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (app *Service) Create(request CreateRequest) (CreateResponse, error) {

	randomId, err := uuid.NewGen().NewV4()
	if err != nil {
		return CreateResponse{Code: http.StatusInternalServerError,
			Message: "[create] uuid create failed",
		}, errors.New("[create] uuid create failed")
	}
	userId := randomId.String()

	if strings.Trim(request.UserName, " ") == "" || strings.Trim(request.MembershipType, " ") == "" {
		return CreateResponse{}, errors.New("[create] username or membership-type is not entered")
	}

	if isInvalidMembership(request.MembershipType) {
		return CreateResponse{Code: http.StatusBadRequest,
			Message: "[create] membership type is invalid",
		}, errors.New("[create] membership type is invalid")
	}

	if app.isDuplicateName(request.UserName) {
		return CreateResponse{Code: http.StatusBadRequest,
			Message: "[create] username is duplicated",
		}, errors.New("[create] username is duplicated")
	}

	app.repository.Create(Membership{
		userId, request.UserName, request.MembershipType,
	})

	return CreateResponse{
		Code:           http.StatusCreated,
		Message:        "Created",
		ID:             userId,
		MembershipType: request.MembershipType,
	}, nil
}

func (app *Service) Update(request UpdateRequest) (UpdateResponse, error) {

	if strings.Trim(request.ID, " ") == "" || strings.Trim(request.UserName, " ") == "" || strings.Trim(request.MembershipType, " ") == "" {
		return UpdateResponse{
			Code:    http.StatusBadRequest,
			Message: "[update] ID or username, membership-type is not entered",
		}, errors.New("[update] ID or username, membership-type is not entered")
	}

	if isInvalidMembership(request.MembershipType) {
		return UpdateResponse{Code: http.StatusBadRequest,
			Message: "[update] membership type is invalid",
		}, errors.New("[update] membership type is invalid")
	}

	if app.isDuplicateName(request.UserName) {
		return UpdateResponse{Code: http.StatusBadRequest,
			Message: "[update] username is duplicated",
		}, errors.New("[update] username is duplicated")
	}

	res := app.repository.Update(Membership{ID: request.ID, UserName: request.UserName, MembershipType: request.MembershipType})

	if res.ID == "" {
		return UpdateResponse{
			Code:    http.StatusNoContent,
			Message: "[update] ID is not exists",
		}, errors.New("[update] ID is not exists")
	}

	return UpdateResponse{
		Code:           http.StatusCreated,
		Message:        "Created",
		ID:             res.ID,
		UserName:       res.UserName,
		MembershipType: res.MembershipType,
	}, nil
}

func (app *Service) Delete(id string) (DeleteResponse, error) {

	if strings.Trim(id, " ") == "" {
		return DeleteResponse{
			Code:    http.StatusBadRequest,
			Message: "[delete] ID is not entered",
		}, errors.New("[delete] ID is not entered")
	}
	err := app.repository.Delete(id)

	if err != nil {
		return DeleteResponse{
			Code:    http.StatusNoContent,
			Message: err.Error(),
		}, err
	}

	return DeleteResponse{
		Code:    http.StatusOK,
		Message: "OK",
	}, nil
}

func (app *Service) Read(id string) (ReadResponse, error) {

	if strings.Trim(id, " ") == "" {
		return ReadResponse{
			Code:    http.StatusBadRequest,
			Message: "[read] ID is not entered",
		}, errors.New("[read] ID is not entered")
	}
	res, err := app.repository.ReadById(id)

	if err != nil {
		return ReadResponse{
			Code:    http.StatusNoContent,
			Message: err.Error(),
		}, nil
	}

	return ReadResponse{
		Code:           http.StatusOK,
		Message:        "OK",
		ID:             res.ID,
		UserName:       res.UserName,
		MembershipType: res.MembershipType,
	}, nil
}

// isDuplicateName returns a bool value whether if username is duplicated or not
func (app *Service) isDuplicateName(userName string) bool {
	if app.repository.ReadCountByName(userName) > 0 {
		return true
	}
	return false
}

// isInvalidMembership returns a bool value whether if membershipType is valid or not
func isInvalidMembership(membershipType string) bool {
	for _, value := range validMemberships {
		if value == membershipType {
			return false
		}
	}
	return true
}

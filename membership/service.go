package membership

import (
	"errors"
	"github.com/gofrs/uuid"
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
		return CreateResponse{}, errors.New("[create] uuid create failed")
	}
	userId := randomId.String()

	if request.UserName == "" || request.MembershipType == "" {
		return CreateResponse{}, errors.New("[create] username or membership-type is not entered")
	}

	if isInvalidMembership(request.MembershipType) {
		return CreateResponse{}, errors.New("[create] membership type is invalid")
	}

	if app.isDuplicateName(request.UserName) {
		return CreateResponse{}, errors.New("[create] username is duplicated")
	}

	app.repository.Create(Membership{
		userId, request.UserName, request.MembershipType,
	})

	return CreateResponse{userId, request.MembershipType}, nil
}

func (app *Service) Update(request UpdateRequest) (UpdateResponse, error) {

	if request.ID == "" || request.UserName == "" || request.MembershipType == "" {
		return UpdateResponse{}, errors.New("[update] ID or username, membership-type is not entered")
	}

	if isInvalidMembership(request.MembershipType) {
		return UpdateResponse{}, errors.New("[update] membership type is invalid")
	}

	if app.isDuplicateName(request.UserName) {
		return UpdateResponse{}, errors.New("[update] username is duplicated")
	}

	res := app.repository.Update(Membership{ID: request.ID, UserName: request.UserName, MembershipType: request.MembershipType})

	if res.ID == "" {
		return UpdateResponse{}, errors.New("[update] ID is not exists")
	}

	return UpdateResponse{res.ID, res.UserName, res.MembershipType}, nil
}

func (app *Service) Delete(id string) error {

	if id == "" {
		return errors.New("[delete] ID is not entered")
	}
	err := app.repository.Delete(id)

	return err
}

func (app *Service) Read(id string) (ReadResponse, error) {

	if id == "" {
		return ReadResponse{}, errors.New("[read] ID is not entered")
	}
	res, err := app.repository.ReadById(id)

	return ReadResponse{res.ID, res.UserName, res.MembershipType}, err
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

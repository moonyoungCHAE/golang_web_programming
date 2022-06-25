package membership

import (
	"github.com/google/uuid"
)

type Application struct {
	repository Repository
}

func NewApplication(repository Repository) *Application {
	return &Application{repository: repository}
}

func (app *Application) Create(request CreateRequest) (*CreateResponse, error) {
	membership := Membership{uuid.New().String(), request.UserName, request.MembershipType}

	if _, err := app.repository.Create(membership); err != nil {
		return nil, err
	}
	return &CreateResponse{membership.ID, membership.MembershipType}, nil
}
func (app *Application) Update(request UpdateRequest) (*UpdateResponse, error) {
	membership := Membership{uuid.New().String(), request.UserName, request.MembershipType}
	if _, err := app.repository.ModifyMember(membership); err != nil {
		return nil, err
	}
	return &UpdateResponse{membership.ID, membership.UserName, membership.MembershipType}, nil
}

func (app *Application) Delete(id string) error {
	if _, err := app.repository.RemoveByID(id); err != nil {
		return err
	}
	return nil
}

package internal

import (
	"strconv"
)

type Application struct {
	repository Repository
}

func NewApplication(repository Repository) *Application {
	return &Application{repository: repository}
}

func (app *Application) Create(request CreateRequest) (CreateResponse, error) {
	membershipBuilder := NewMembershipGenerator()
	memberRepository := app.repository
	memberCount := len(memberRepository.data)

	id := strconv.Itoa(memberCount + 1)

	membershipBuilder.
		SetID(id).
		SetUserName(request.UserName).
		SetMembershipType(request.MembershipType)

	membership, err := membershipBuilder.GetMembership()
	if err != nil {
		return CreateResponse{}, err
	}

	_, err = app.repository.CreateMembership(*membership)
	if err != nil {
		return CreateResponse{}, err
	}

	return CreateResponse{membership.ID, membership.MembershipType}, nil
}

func (app *Application) Update(request UpdateRequest) (UpdateResponse, error) {
	newMembership, err := NewMembershipGenerator().
		SetID(request.ID).
		SetUserName(request.UserName).
		SetMembershipType(request.MembershipType).
		GetMembership()
	if err != nil {
		return UpdateResponse{}, err
	}

	_, err = app.repository.UpdateMembership(*newMembership)
	if err != nil {
		return UpdateResponse{}, err
	}

	return UpdateResponse{
		ID:             newMembership.ID,
		UserName:       newMembership.UserName,
		MembershipType: newMembership.MembershipType,
	}, nil
}

func (app *Application) Delete(request DeleteRequest) (DeleteResponse, error) {
	if request.ID == "" {
		return DeleteResponse{}, ErrUserIDIsRequired
	}
	if _, ok := app.repository.GetMembershipByID(request.ID); !ok {
		return DeleteResponse{}, ErrUserIDNotFound
	}

	return app.repository.DeleteMembership(request)
}

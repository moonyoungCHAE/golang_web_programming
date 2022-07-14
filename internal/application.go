package internal

import (
	"log"
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
	member_repository := app.repository
	member_count := len(member_repository.data)

	id := strconv.Itoa(member_count + 1)
	log.Println(id)

	membershipBuilder.
		SetID(id).
		SetUserName(request.UserName).
		SetMembershipType(request.MembershipType)

	membership, err := membershipBuilder.GetMembership()
	if err != nil {
		return CreateResponse{}, err
	}

	_, err = app.repository.
		CraateRepositoryData(*membership)
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

	_, err = app.repository.UpdateRepositoryData(*newMembership)
	if err != nil {
		return UpdateResponse{}, err
	}

	return UpdateResponse{
		ID:             newMembership.ID,
		UserName:       newMembership.UserName,
		MembershipType: newMembership.MembershipType,
	}, nil
}

func (app *Application) Delete(id string) error {
	m := app.repository.data[id]
	membership, err := NewMembershipGenerator().
		SetID(m.ID).
		SetUserName(m.UserName).
		SetMembershipType(m.MembershipType).
		GetMembership()
	if err != nil {
		return err
	}
	err = app.repository.DeleteRepositoryData(*membership)
	if err != nil {
		return err
	}
	return nil
}

package membership

import (
	"errors"
)

var (
	SameNameErr = errors.New("same_name")

	WrongIdErr         = errors.New("wrong_id")
	WrongMembershipErr = errors.New("wrong_membership")

	NoIdErr         = errors.New("no_ID")
	NoNameErr       = errors.New("no_name")
	NoMembershipErr = errors.New("no_membership")
)

func (app *Application) ValidateCreate(request CreateRequest) error {
	if _, exist := app.repository.data[request.UserName]; exist {
		return SameNameErr
	}

	if request.UserName == "" {
		return NoNameErr
	}

	if request.MembershipType == "" {
		return NoMembershipErr
	}

	if request.MembershipType != "naver" && request.MembershipType != "toss" && request.MembershipType != "payco" {
		return WrongMembershipErr
	}
	return nil
}

func (app *Application) ValidateUpdate(request UpdateRequest) error {
	if request.ID == "" {
		return NoIdErr
	}

	if request.UserName == "" {
		return NoNameErr
	}

	if request.MembershipType == "" {
		return NoMembershipErr
	}

	if request.MembershipType != "naver" && request.MembershipType != "toss" && request.MembershipType != "payco" {
		return WrongMembershipErr
	}

	if val, exist := app.repository.data[request.UserName]; exist && val.ID != request.ID {
		return SameNameErr
	}
	return nil
}

func (app *Application) ValidateDelete(id string) error {
	if id == "" {
		return NoIdErr
	}
	if _, exist := app.repository.data[id]; !exist {
		return WrongIdErr
	}
	return nil
}

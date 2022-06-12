package membership

import "fmt"

type Application struct {
	repository Repository
}

func NewApplication(repository Repository) *Application {
	return &Application{repository: repository}
}

func (app *Application) Create(request CreateRequest) (CreateResponse, error) {
	if _, exist := app.repository.data[request.UserName]; exist {
		return CreateResponse{}, fmt.Errorf("same_name")
	}

	if request.UserName == "" {
		return CreateResponse{}, fmt.Errorf("no_name")
	}

	if request.MembershipType == "" {
		return CreateResponse{}, fmt.Errorf("no_membership")
	}

	if request.MembershipType != "naver" && request.MembershipType != "toss" && request.MembershipType != "payco" {
		return CreateResponse{}, fmt.Errorf("wrong_membership")
	}

	membership := Membership{
		ID:             request.UserName,
		UserName:       request.UserName,
		MembershipType: request.MembershipType,
	}
	app.repository.data[membership.ID] = membership

	return CreateResponse{membership.ID, membership.MembershipType}, nil
}

func (app *Application) Update(request UpdateRequest) (UpdateResponse, error) {
	if request.ID == "" {
		return UpdateResponse{}, fmt.Errorf("no_ID")
	}

	if request.UserName == "" {
		return UpdateResponse{}, fmt.Errorf("no_name")
	}

	if request.MembershipType == "" {
		return UpdateResponse{}, fmt.Errorf("no_membership")
	}

	if request.MembershipType != "naver" && request.MembershipType != "toss" && request.MembershipType != "payco" {
		return UpdateResponse{}, fmt.Errorf("wrong_membership")
	}

	if val, exist := app.repository.data[request.UserName]; exist && val.ID != request.ID {
		return UpdateResponse{}, fmt.Errorf("same_name")
	}

	app.repository.data[request.UserName] = Membership{
		ID:             request.ID,
		UserName:       request.UserName,
		MembershipType: request.MembershipType,
	}

	return UpdateResponse{
		app.repository.data[request.UserName].ID,
		app.repository.data[request.UserName].UserName,
		app.repository.data[request.UserName].MembershipType,
	}, nil
}

func (app *Application) Delete(id string) error {
	if id == "" {
		return fmt.Errorf("no_ID")
	}
	if _, exist := app.repository.data[id]; !exist {
		return fmt.Errorf("wrong_ID")
	}

	delete(app.repository.data, id)
	return nil
}

package membership

import "github.com/google/uuid"

type Application struct {
	repository Repository
}

func NewApplication(repository Repository) *Application {
	return &Application{repository: repository}
}

func (app *Application) Create(request CreateRequest) (CreateResponse, error) {
	if err := app.ValidateCreate(request); err != nil {
		return CreateResponse{}, err
	}

	membership := Membership{
		ID:             request.UserName,
		UserName:       uuid.NewString(),
		MembershipType: request.MembershipType,
	}

	res := app.AddData(membership)

	return res, nil
}

func (app *Application) Update(request UpdateRequest) (UpdateResponse, error) {
	if err := app.ValidateUpdate(request); err != nil {
		return UpdateResponse{}, err
	}

	membership := Membership{
		ID:             request.ID,
		UserName:       request.UserName,
		MembershipType: request.MembershipType,
	}

	res := app.FixData(membership)

	return res, nil
}

func (app *Application) Delete(id string) error {
	if err := app.ValidateDelete(id); err != nil {
		return err
	}

	delete(app.repository.data, id)

	return nil
}

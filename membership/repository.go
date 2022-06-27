package membership

type Repository struct {
	data map[string]Membership
}

func NewRepository(data map[string]Membership) *Repository {
	return &Repository{data: data}
}
func (app *Application) AddData(membership Membership) CreateResponse {
	app.repository.data[membership.ID] = membership

	return CreateResponse{
		app.repository.data[membership.ID].ID,
		app.repository.data[membership.ID].MembershipType,
	}
}

func (app *Application) FixData(membership Membership) UpdateResponse {
	app.repository.data[membership.UserName] = membership

	return UpdateResponse{
		app.repository.data[membership.UserName].ID,
		app.repository.data[membership.UserName].UserName,
		app.repository.data[membership.UserName].MembershipType,
	}
}

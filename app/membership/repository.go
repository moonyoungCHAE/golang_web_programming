package membership

type Repository struct {
	data map[string]Membership
}

func NewRepository(data map[string]Membership) *Repository {
	return &Repository{data: data}
}
func (r *Repository) Create(membership Membership) {
	r.data[membership.ID] = membership
}

func (r *Repository) Update(membership Membership) {
	r.data[membership.UserName] = membership
}

package membership

import "strconv"

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

func (r *Repository) Delete(id string) {
	delete(r.data, id)
}

func (r *Repository) GetByID(id string) Membership {
	return Membership{
		ID:             r.data[id].ID,
		UserName:       r.data[id].UserName,
		MembershipType: r.data[id].MembershipType,
	}
}

func (r *Repository) GetSome(offset string, limit string) []Membership {
	var res []Membership
	o, _ := strconv.Atoi(offset)
	l, _ := strconv.Atoi(limit)
	for _, val := range r.data {
		res = append(res, val)
	}
	return res[o : o+l]
}

func (r *Repository) GetAll() []Membership {
	var res []Membership
	for _, val := range r.data {
		res = append(res, val)
	}
	return res
}

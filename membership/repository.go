package membership

import "errors"

type Repository struct {
	data map[string]Membership
}

func NewRepository(data map[string]Membership) *Repository {
	return &Repository{data: data}
}

func (r *Repository) Create(m Membership) Membership {
	r.data[m.ID] = m
	return r.data[m.ID]
}

func (r *Repository) Update(m Membership) Membership {
	for _, membership := range r.data {
		if membership.ID == m.ID {
			r.data[m.ID] = m
			break
		}
	}
	return r.data[m.ID]
}

func (r *Repository) Delete(id string) error {
	if _, exists := r.data[id]; !exists {
		return errors.New("[delete] ID is invalid (non-exists)")
	}
	delete(r.data, id)
	return nil
}

func (r *Repository) ReadById(id string) (Membership, error) {

	var membership, exists = r.data[id]
	if !exists {
		return membership, errors.New("[read] ID is invalid (non-exists)")
	}
	return membership, nil
}

func (r *Repository) ReadCountByName(name string) int {
	var readCount int = 0
	for _, m := range r.data {
		if m.UserName == name {
			readCount++
		}
	}
	return readCount
}

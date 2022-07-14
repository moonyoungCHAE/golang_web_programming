package internal

import "errors"

type Repository struct {
	data map[string]Membership
}

func NewRepository(data map[string]Membership) *Repository {
	return &Repository{data: data}
}

func (r *Repository) CraateRepositoryData(m Membership) (Membership, error) {
	for _, membership := range r.data {
		if membership.UserName == m.UserName {
			return Membership{}, errors.New("already existed user_name")
		}
	}
	r.data[m.ID] = m
	return m, nil
}

func (r *Repository) UpdateRepositoryData(m Membership) (Membership, error) {
	for _, membership := range r.data {
		if membership.ID == m.ID {
			continue
		}
		if membership.UserName == m.UserName {
			return Membership{}, errors.New("already existed name")
		}
	}
	r.data[m.ID] = m
	return m, nil
}

func (r *Repository) DeleteRepositoryData(membership Membership) error {
	_, ok := r.data[membership.ID]
	if ok {
		delete(r.data, membership.ID)
		return nil
	}
	return errors.New("not exist id")
}

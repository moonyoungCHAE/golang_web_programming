package internal

import "github.com/google/uuid"

type Repository struct {
	data map[string]Membership
}

func NewRepository(data map[string]Membership) *Repository {
	return &Repository{data: data}
}

func (r *Repository) UserNameExists(userName string) bool {
	for _, membership := range r.data {
		if membership.UserName == userName {
			return true
		}
	}
	return false
}

func (r *Repository) Create(req CreateRequest) (CreateResponse, error) {
	id := uuid.New().String()
	r.data[id] = Membership{ID: id, UserName: req.UserName, MembershipType: req.MembershipType}
	return CreateResponse{ID: id, MembershipType: req.MembershipType}, nil
}

func (r *Repository) GetByID(id string) (Membership, bool) {
	membership, ok := r.data[id]
	return membership, ok
}

func (r *Repository) Update(req UpdateRequest) (UpdateResponse, error) {
	r.data[req.ID] = Membership{ID: req.ID, UserName: req.UserName, MembershipType: req.MembershipType}
	return UpdateResponse{ID: req.ID, UserName: req.UserName, MembershipType: req.MembershipType}, nil
}

func (r *Repository) Delete(req DeleteRequest) (DeleteResponse, error) {
	delete(r.data, req.ID)
	return DeleteResponse{ID: req.ID, UserName: r.data[req.ID].UserName, MembershipType: r.data[req.ID].MembershipType}, nil
}

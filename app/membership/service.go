package membership

import (
	"github.com/google/uuid"
	"net/http"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (service *Service) Create(request CreateRequest) (CreateResponse, error) {
	if err := service.ValidateCreate(request); err != nil {
		return CreateResponse{}, err
	}
	membership := Membership{
		ID:             uuid.NewString(),
		UserName:       request.UserName,
		MembershipType: request.MembershipType,
	}
	service.repository.Create(membership)
	return CreateResponse{
		Code:           http.StatusCreated,
		Message:        http.StatusText(http.StatusCreated),
		ID:             membership.ID,
		MembershipType: membership.MembershipType,
	}, nil
}

func (service *Service) Update(request UpdateRequest) (UpdateResponse, error) {
	if err := service.ValidateUpdate(request); err != nil {
		return UpdateResponse{}, err
	}
	membership := Membership{
		ID:             request.ID,
		UserName:       request.UserName,
		MembershipType: request.MembershipType,
	}
	service.repository.Update(membership)
	return UpdateResponse{
		Code:           http.StatusOK,
		Message:        http.StatusText(http.StatusOK),
		ID:             membership.ID,
		UserName:       membership.UserName,
		MembershipType: membership.MembershipType,
	}, nil
}

func (service *Service) Delete(id string) (DeleteResponse, error) {
	if err := service.ValidateDelete(id); err != nil {
		return DeleteResponse{}, err
	}
	service.repository.Delete(id)
	return DeleteResponse{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	}, nil
}

func (service *Service) GetByID(id string) (GetResponse, error) {
	if err := service.ValidateGetByID(id); err != nil {
		return GetResponse{}, err
	}
	membership := service.repository.GetByID(id)
	return GetResponse{
		Code:           http.StatusOK,
		Message:        http.StatusText(http.StatusOK),
		ID:             membership.ID,
		UserName:       membership.UserName,
		MembershipType: membership.MembershipType,
	}, nil
}

func (service *Service) GetSome(offset string, limit string) (GetSomeResponse, error) {
	if err := service.ValidateGetSome(offset, limit); err != nil {
		return GetSomeResponse{}, err
	}
	var res []Membership
	res = service.repository.GetSome(offset, limit)
	return GetSomeResponse{
		Code:       http.StatusOK,
		Message:    http.StatusText(http.StatusOK),
		Membership: res,
	}, nil
}

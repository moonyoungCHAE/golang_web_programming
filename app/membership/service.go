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
		ID:             request.UserName,
		UserName:       uuid.NewString(),
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

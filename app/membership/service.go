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

func (service *Service) Create(request CreateRequest) GetResponse {
	membership := Membership{uuid.New().String(), request.UserName, request.MembershipType, request.Status}
	member, err := service.repository.Create(membership)
	if err != nil {
		return GetResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	return GetResponse{
		Code:    http.StatusCreated,
		Message: http.StatusText(http.StatusCreated),
		Member:  member,
	}
}

func (service *Service) GetByID(id string) GetResponse {
	membership, err := service.repository.GetById(id)
	if err != nil {
		return GetResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return GetResponse{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Member:  membership,
	}
}
func (service *Service) RemoveByID(id string) GetResponse {
	member, err := service.repository.RemoveByID(id)
	if err != nil {
		return GetResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return GetResponse{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Member:  member,
	}
}
func (service *Service) ModifyMember(request UpdateRequest) GetResponse {
	membership := Membership{uuid.New().String(), request.UserName, request.MembershipType, request.Status}
	member, err := service.repository.ModifyMember(membership)
	if err != nil {
		return GetResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return GetResponse{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Member:  member,
	}
}

func (service *Service) GetMembers(offset string, limit string) GetResponse {
	memberList, err := service.repository.GetMembers(offset, limit)
	if err != nil {
		return GetResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return GetResponse{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Member:  memberList,
	}
}

package internal

import (
	"github.com/google/uuid"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Create(request CreateRequest) (CreateResponse, error) {
	membership := Membership{uuid.New().String(), request.UserName, request.MembershipType}
	s.repository.CreateMembership(membership)
	return CreateResponse{
		ID:             membership.ID,
		MembershipType: membership.MembershipType,
	}, nil
}

func (s *Service) GetByID(id string) (GetResponse, error) {
	membership, err := s.repository.GetMembershipByID(id)
	if err != nil {
		return GetResponse{}, nil
	}
	return GetResponse{
		ID:             membership.ID,
		UserName:       membership.UserName,
		MembershipType: membership.MembershipType,
	}, nil
}

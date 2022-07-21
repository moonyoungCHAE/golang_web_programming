package internal

import (
	"net/http"
	"strconv"
	"strings"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (service *Service) IsDuplicateName(userName string) bool {
	_, err := service.repository.GetMembershipByName(userName)
	if err != nil {
		return true
	}
	return false
}

func (service *Service) Create(request CreateRequest) (CreateResponse, error) {
	membershipBuilder := NewMembershipGenerator()
	memberRepository := service.repository
	memberCount := len(memberRepository.data)

	id := strconv.Itoa(memberCount + 1)

	membershipBuilder.
		SetID(id).
		SetUserName(request.UserName).
		SetMembershipType(request.MembershipType)

	membership, err := membershipBuilder.GetMembership()
	if err != nil {
		return CreateResponse{
			Code:    http.StatusInternalServerError,
			Message: "MembershipBuilder Error",
		}, err
	}

	_, err = service.repository.CreateMembership(*membership)
	if err != nil {
		return CreateResponse{
			Code:    http.StatusInternalServerError,
			Message: "Repository : Create Error",
		}, err
	}

	return CreateResponse{
		Code:           http.StatusCreated,
		ID:             membership.ID,
		MembershipType: membership.MembershipType,
	}, nil
}

func (app *Service) Read(id string) (ReadResponse, error) {

	if strings.Trim(id, " ") == "" {
		return ReadResponse{
			Code:    http.StatusBadRequest,
			Message: "user id is required",
		}, ErrUserIDIsRequired
	}
	res, err := app.repository.GetMembershipByID(id)

	if err != nil {
		return ReadResponse{
			Code:    http.StatusInternalServerError,
			Message: "user not found",
		}, ErrUserNotFound
	}

	return ReadResponse{
		Code:           http.StatusOK,
		Message:        "Success",
		ID:             res.ID,
		UserName:       res.UserName,
		MembershipType: res.MembershipType,
	}, nil
}

func (service *Service) Update(request UpdateRequest) (UpdateResponse, error) {
	if strings.Trim(request.ID, " ") == "" || strings.Trim(request.UserName, " ") == "" || strings.Trim(request.MembershipType, " ") == "" {
		return UpdateResponse{
			Code:    http.StatusBadRequest,
			Message: "ID or username or membership_type is nil",
		}, ErrUpdateRequestIsRequired
	}
	if IsvalidMembership(request.MembershipType) {
		return UpdateResponse{Code: http.StatusBadRequest,
			Message: "choose membership type : naver, payco, toss",
		}, ErrInvalidMembershipType
	}

	if service.IsDuplicateName(request.UserName) {
		return UpdateResponse{Code: http.StatusBadRequest,
			Message: "username is duplicated",
		}, ErrUpdateRequestIsRequired
	}

	update_member, err := service.repository.UpdateMembership(Membership{ID: request.ID, UserName: request.UserName, MembershipType: request.MembershipType})
	if err != nil {
		return UpdateResponse{
			Code:    http.StatusInternalServerError,
			Message: "Can't update membership",
		}, err
	}

	return UpdateResponse{
		ID:             update_member.ID,
		UserName:       update_member.UserName,
		MembershipType: update_member.MembershipType,
	}, nil
}

func (service *Service) GetByID(id string) (GetResponse, error) {
	membership, err := service.repository.GetMembershipByID(id)
	if err != nil {
		return GetResponse{}, nil
	}
	return GetResponse{
		ID:             membership.ID,
		UserName:       membership.UserName,
		MembershipType: membership.MembershipType,
	}, nil
}

func (service *Service) Delete(request DeleteRequest) (DeleteResponse, error) {
	if request.ID == "" {
		return DeleteResponse{
			Code:    http.StatusBadRequest,
			Message: "user id is required",
		}, ErrUserIDIsRequired

	}
	if _, err := service.repository.GetMembershipByID(request.ID); err != nil {
		return DeleteResponse{
			Code:    http.StatusBadRequest,
			Message: "user id not found",
		}, ErrUserIDNotFound
	}

	deleteReq := DeleteRequest{ID: request.ID}

	res, err := service.repository.DeleteMembership(deleteReq)
	if err != nil {
		return DeleteResponse{
			Code:    http.StatusInternalServerError,
			Message: "Can't delete membership",
		}, err
	}

	return res, nil
}

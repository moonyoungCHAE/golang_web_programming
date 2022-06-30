package membership

import (
	"errors"
	"strconv"
)

var (
	SameNameErr   = errors.New("same_name")
	OutOfBoundErr = errors.New("out_of_bound")

	WrongIdErr = errors.New("wrong_id")
	// rongLimitErr      = errors.New("wrong_limit")
	WrongOffsetErr     = errors.New("wrong_offset")
	WrongMembershipErr = errors.New("wrong_membership")

	NoIdErr         = errors.New("no_ID")
	NoNameErr       = errors.New("no_name")
	NoLimitErr      = errors.New("no_limit")
	NoOffsetErr     = errors.New("no_offset")
	NoMembershipErr = errors.New("no_membership")
)

func (service *Service) ValidateCreate(request CreateRequest) error {
	if _, exist := service.repository.data[request.UserName]; exist {
		return SameNameErr
	}

	if request.UserName == "" {
		return NoNameErr
	}

	if request.MembershipType == "" {
		return NoMembershipErr
	}

	if request.MembershipType != "naver" && request.MembershipType != "toss" && request.MembershipType != "payco" {
		return WrongMembershipErr
	}
	return nil
}

func (service *Service) ValidateUpdate(request UpdateRequest) error {
	if request.ID == "" {
		return NoIdErr
	}

	if request.UserName == "" {
		return NoNameErr
	}

	if request.MembershipType == "" {
		return NoMembershipErr
	}

	if request.MembershipType != "naver" && request.MembershipType != "toss" && request.MembershipType != "payco" {
		return WrongMembershipErr
	}

	if val, exist := service.repository.data[request.UserName]; exist && val.ID != request.ID {
		return SameNameErr
	}
	return nil
}

func (service *Service) ValidateDelete(id string) error {
	if id == "" {
		return NoIdErr
	}
	if _, exist := service.repository.data[id]; !exist {
		return WrongIdErr
	}
	return nil
}

func (service *Service) ValidateGetByID(id string) error {
	if id == "" {
		return NoIdErr
	}
	if _, exist := service.repository.data[id]; !exist {
		return WrongIdErr
	}
	return nil
}

func (service *Service) ValidateGetSome(offset string, limit string) error {
	if offset == "" {
		return NoOffsetErr
	}
	if limit == "" {
		return NoLimitErr
	}
	o, _ := strconv.Atoi(offset)
	if o < len(service.repository.data) {
		return WrongOffsetErr
	}
	l, _ := strconv.Atoi(limit)
	if o+l < len(service.repository.data) {
		return OutOfBoundErr
	}
	return nil
}

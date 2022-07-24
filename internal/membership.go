package internal

type Membership struct {
	ID             string
	UserName       string
	MembershipType string
}

type MembershipGenerator struct {
	Membership *Membership
}

func NewMembershipGenerator() *MembershipGenerator {
	return &MembershipGenerator{Membership: &Membership{}}
}

func (m *MembershipGenerator) SetID(id string) *MembershipGenerator {
	m.Membership.ID = id
	return m
}

func (m *MembershipGenerator) SetUserName(userName string) *MembershipGenerator {
	m.Membership.UserName = userName
	return m
}

func (m *MembershipGenerator) SetMembershipType(membershipType string) *MembershipGenerator {
	m.Membership.MembershipType = membershipType
	return m
}

func (m *MembershipGenerator) GetMembership() (*Membership, error) {
	err := m.validateMembership()
	if err != nil {
		return nil, err
	}
	return m.Membership, nil
}

func (m *MembershipGenerator) validateMembership() error {
	if m.Membership.ID == "" {
		return ErrUserIDIsRequired
	}
	if m.Membership.UserName == "" {
		return ErrUserNameIsRequired
	}
	if m.Membership.MembershipType == "" {
		return ErrMembershipTypeIsRequired
	}
	if !(m.Membership.MembershipType == "naver" || m.Membership.MembershipType == "payco" || m.Membership.MembershipType == "toss") {
		return ErrInvalidMembershipType
	}
	return nil
}

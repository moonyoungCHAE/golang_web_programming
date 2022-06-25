package membership

type Membership struct {
	ID             string `json:"id,omitempty" logging:"true"`
	UserName       string `json:"userName,omitempty" logging:"true"`
	MembershipType string `json:"membershipType,omitempty" logging:"true"`
	Status         string `json:"status,omitempty" logging:"true"`
}

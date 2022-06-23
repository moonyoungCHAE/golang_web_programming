package membership

type Membership struct {
	ID             string `json:"id"`
	UserName       string `json:"user_name"`
	MembershipType string `json:"membership_type"`
}

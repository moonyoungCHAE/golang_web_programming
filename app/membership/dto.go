package membership

type CreateRequest struct {
	UserName       string `json:"user_name"`
	MembershipType string `json:"membership_type"`
}

type CreateResponse struct {
	Code           int    `json:"-"`
	Message        string `json:"message"`
	ID             string `json:"id"`
	MembershipType string `json:"membership_type"`
}

type UpdateRequest struct {
	ID             string `json:"-`
	UserName       string `json:"user_name"`
	MembershipType string `json:"membership_type"`
}

type UpdateResponse struct {
	Code       int        `json:"-"`
	Message    string     `json:"message"`
	Membership Membership `json:membership,omitempty`
}

type ReadResponse struct {
	Code       int        `json:"-"`
	Message    string     `json:"message"`
	Membership Membership `json:membership,omitempty`
}

type ReadAllResponse struct {
	Code        int          `json:"-"`
	Message     string       `json:"message"`
	Memberships []Membership `json:"memberships,omitempty"`
}

type DeleteResponse struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

package membership

type CreateRequest struct {
	UserName       string `json:"user-name"`
	MembershipType string `json:"membership-type"`
}

type CreateResponse struct {
	Code           int    `json:"-"` // 코드는 사용자에게 안 보여줄 거니까
	Message        string `json:"message"`
	ID             string
	MembershipType string
}

type UpdateRequest struct {
	ID             string `json:"id"`
	UserName       string `json:"user-name"`
	MembershipType string `json:"membership-type"`
}

type UpdateResponse struct {
	Code           int    `json:"-"` // 코드는 사용자에게 안 보여줄 거니까
	Message        string `json:"message"`
	ID             string
	UserName       string
	MembershipType string
}

type DeleteResponse struct {
	Code    int    `json:"-"` // 코드는 사용자에게 안 보여줄 거니까
	Message string `json:"message"`
}

type GetResponse struct {
	Code           int    `json:"-"` // 코드는 사용자에게 안 보여줄 거니까
	Message        string `json:"message"`
	ID             string
	UserName       string
	MembershipType string
}

type GetSomeResponse struct {
	Code       int          `json:"-"` // 코드는 사용자에게 안 보여줄 거니까
	Message    string       `json:"message"`
	Membership []Membership `json:"membership"`
}

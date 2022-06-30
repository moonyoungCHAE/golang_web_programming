package membership

type CreateRequest struct {
	UserName       string
	MembershipType string
}

type CreateResponse struct {
	Code           int    `json:"-"` // 코드는 사용자에게 안 보여줄 거니까
	Message        string `json:"message"`
	ID             string
	MembershipType string
}

type UpdateRequest struct {
	ID             string
	UserName       string
	MembershipType string
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

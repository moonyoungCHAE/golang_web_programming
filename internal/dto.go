package internal

type CreateRequest struct {
	UserName       string
	MembershipType string
}

type CreateResponse struct {
	ID             string
	MembershipType string
}

type UpdateRequest struct {
	ID             string
	UserName       string
	MembershipType string
}

type UpdateResponse struct {
	ID             string
	UserName       string
	MembershipType string
}

type DeleteRequest struct {
	ID string
}

type DeleteResponse struct {
	ID             string
	UserName       string
	MembershipType string
}

type GetResponse struct {
	ID             string `json:"id"`
	UserName       string `json:"user_name"`
	MembershipType string `json:"membership_type"`
}

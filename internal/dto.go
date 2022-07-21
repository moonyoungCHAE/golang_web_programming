package internal

type CreateRequest struct {
	UserName       string `json:"user_name"`
	MembershipType string `json:"membership_type"`
}

type CreateResponse struct {
	ID             string `json:"id"`
	MembershipType string `json:"membership_type"`
}

type UpdateRequest struct {
	ID             string `json:"id"`
	UserName       string `json:"user_name"`
	MembershipType string `json:"membership_type"`
}

type UpdateResponse struct {
	ID             string `json:"id"`
	UserName       string `json:"user_name"`
	MembershipType string `json:"membership_type"`
}

type DeleteRequest struct {
	ID string `json:"id"`
}

type DeleteResponse struct {
	ID             string `json:"id"`
	UserName       string `json:"user_name"`
	MembershipType string `json:"membership_type"`
}

type GetResponse struct {
	ID             string `json:"id"`
	UserName       string `json:"user_name"`
	MembershipType string `json:"membership_type"`
}

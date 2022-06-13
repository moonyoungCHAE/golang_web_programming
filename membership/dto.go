package membership

// FindResponse : response for find func
type FindResponse struct {
	ID             string
	UserName       string
	MembershipType string
}

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

package internal

type CreateRequest struct {
	UserName       string
	MembershipType string
}

type CreateResponse struct {
	ID             string
	MembershipType string
}

type ReadRequest struct {
	ID string
}

type ReadResponse struct {
	ID             string
	UserName       string
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

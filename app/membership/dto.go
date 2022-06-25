package membership

import (
	"github.com/labstack/echo"
)

type CreateRequest struct {
	UserName       string `json:"userName,omitempty"`
	MembershipType string `json:"membershipType,omitempty"`
	Status         string `json:"status,omitempty"`
}

type CreateResponse struct {
	Code           int    `json:"code,omitempty"`
	Message        string `json:"message,omitempty"`
	ID             string `json:"id,omitempty"`
	MembershipType string `json:"membershipType,omitempty"`
	Status         string `json:"status,omitempty"`
}

type UpdateRequest struct {
	ID             string `json:"id,omitempty"`
	UserName       string `json:"userName,omitempty"`
	MembershipType string `json:"membershipType,omitempty"`
	Status         string `json:"status,omitempty"`
}

type UpdateResponse struct {
	ID             string `json:"id,omitempty"`
	UserName       string `json:"userName,omitempty"`
	MembershipType string `json:"membershipType,omitempty"`
	Status         string `json:"status,omitempty"`
}

type GetResponse struct {
	Code    int         `json:"code,omitempty" logging:"true"`
	Message string      `json:"message,omitempty" logging:"true"`
	Member  interface{} `json:"member,omitempty" logging:"true"`
	//MemberList []*Membership `json:"memberList,omitempty"`
}

func (m *GetResponse) Filter(c echo.Context) {
	c.Set("filter", m)
}

package user

type LoginResponse struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
}

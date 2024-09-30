package ecomapi

type (
	SignUpUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	SignUpUserResponse struct {
		Message string `json:"message"`
	}
)

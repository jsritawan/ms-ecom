package itf

import (
	"context"

	"github.com/jsritawan/ms-ecom/ecomapi"
)

type (
	IUserService interface {
		AddNewUser(c context.Context, req *ecomapi.SignUpUserRequest) (*ecomapi.SignUpUserResponse, error)
	}
)

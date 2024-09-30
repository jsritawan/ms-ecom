package handler

import (
	"context"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/jsritawan/ms-ecom/ecomapi"
	"github.com/jsritawan/ms-ecom/internal/model"
)

func (u *UserService) AddNewUser(ctx context.Context, req *ecomapi.SignUpUserRequest) (*ecomapi.SignUpUserResponse, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &model.User{
		Code:     uuid.New().String(),
		Email:    req.Email,
		Password: string(pw),
	}
	user, err := u.UserStorage.Save(ctx, newUser)
	if err != nil {
		return nil, err
	}

	newProfile := &model.Profile{
		UserId: user.Id,
		Status: "A",
		Sex:    "O",
	}
	if _, err := u.ProfileStorage.Save(ctx, newProfile); err != nil {
		return nil, err
	}

	return &ecomapi.SignUpUserResponse{Message: "Sign-Up Successfully"}, nil
}

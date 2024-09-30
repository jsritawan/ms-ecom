package httpadapter

import (
	"github.com/jsritawan/ms-ecom/internal/handler"
	"github.com/jsritawan/ms-ecom/internal/handler/itf"
)

type Adapter struct {
	profileService itf.IProfileService
	userService    itf.IUserService
}

func New(s *handler.Service) *Adapter {
	return &Adapter{
		profileService: handler.NewProfileService(s),
		userService:    handler.NewUserService(s),
	}
}

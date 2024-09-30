package handler

import "github.com/jsritawan/ms-ecom/internal/storage"

type (
	Service struct {
		ProfileStorage storage.IProfileStorage
		UserStorage    storage.IUserStorage
	}

	ProfileService struct {
		*Service
	}

	UserService struct {
		*Service
	}
)

func New(db *storage.Storage) *Service {
	return &Service{
		ProfileStorage: storage.NewProfileStorage(db),
		UserStorage:    storage.NewUserStorage(db),
	}
}

func NewProfileService(s *Service) *ProfileService {
	return &ProfileService{
		Service: s,
	}
}

func NewUserService(s *Service) *UserService {
	return &UserService{
		Service: s,
	}
}

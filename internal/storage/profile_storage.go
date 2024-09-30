package storage

import (
	"github.com/jsritawan/ms-ecom/internal/model"
)

type (
	IProfileStorage interface {
		IStorage[*model.Profile]
	}

	ProfileStorage struct {
		AbstractStorage[*model.Profile]
	}
)

func NewProfileStorage(s *Storage) *ProfileStorage {
	return &ProfileStorage{
		AbstractStorage: AbstractStorage[*model.Profile]{

			db:        s.db,
			tableName: model.ProfileTableName,
		},
	}
}

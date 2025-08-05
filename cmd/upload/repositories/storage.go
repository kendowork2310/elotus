package repositories

import (
	"elotus/cmd/common/daos"
	"elotus/pkg/db"

	"gorm.io/gorm"
)

type Storage interface {
	CreateUpload(upload *daos.Upload) error
}

type storage struct {
	clientDB *gorm.DB
}

func NewStorage() Storage {
	return &storage{
		clientDB: db.NewMySQLClient(),
	}
}

func (s *storage) CreateUpload(upload *daos.Upload) error {
	return s.clientDB.Create(upload).Error
}

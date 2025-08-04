package repositories

import (
	"elotus/pkg/db"
	"gorm.io/gorm"
)

type Storage interface {
}

type storage struct {
	clientDB *gorm.DB
}

func NewStorage() Storage {
	return &storage{
		clientDB: db.NewMySQLClient(),
	}
}

package repositories

import (
	"elotus/cmd/common/daos"
	"elotus/pkg/db"
	"gorm.io/gorm"
)

type Storage interface {
	CreateUser(user *daos.User) error
	GetUserByUsername(username string) (*daos.User, error)
}

type storage struct {
	clientDB *gorm.DB
}

func NewStorage() Storage {
	return &storage{
		clientDB: db.NewMySQLClient(),
	}
}

func (s *storage) CreateUser(user *daos.User) error {
	return s.clientDB.Create(user).Error
}

func (s *storage) GetUserByUsername(username string) (*daos.User, error) {
	var user daos.User
	err := s.clientDB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

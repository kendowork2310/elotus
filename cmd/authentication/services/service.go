package services

import (
	"elotus/cmd/authentication/repositories"
	"elotus/cmd/common/daos"
	"elotus/cmd/common/errs"
	"elotus/pkg/jwt"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type ServiceAuthen interface {
	Register(username, password string) error
	Login(username, password string) (string, error)
}

type service struct {
	repo repositories.Storage
}

func NewServices() ServiceAuthen {
	return &service{
		repo: repositories.NewStorage(),
	}
}

func (s *service) Register(username, password string) error {
	// Check if user already exists
	existingUser, err := s.repo.GetUserByUsername(username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if existingUser != nil {
		return errs.NewCustomError(errs.ErrUsernameExists)
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errs.NewCustomError(errs.ErrInternalServer)
	}

	// Create user
	user := &daos.User{
		Username: username,
		Password: string(hashedPassword),
	}

	return s.repo.CreateUser(user)
}

func (s *service) Login(username, password string) (string, error) {
	// Get user by username
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errs.NewCustomError(errs.ErrInvalidCredentials)
		}
		return "", err
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errs.NewCustomError(errs.ErrInvalidCredentials)
	}

	// Generate JWT token
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", errs.NewCustomError(errs.ErrInternalServer)
	}

	return token, nil
}

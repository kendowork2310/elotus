package services

import (
	"elotus/cmd/authentication/repositories"
	"go.uber.org/zap"
)

type ServiceAuthen interface {
}

type services struct {
	repo repositories.Storage
	logz *zap.Logger
}

func NewServices() ServiceAuthen {
	loggerZap, err := zap.NewProduction()
	if err != nil {
		panic("cannot init logzap")
	}
	return &services{
		repo: repositories.NewStorage(),
		logz: loggerZap,
	}
}

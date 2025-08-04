package services

import (
	"elotus/cmd/upload/repositories"
	"go.uber.org/zap"
)

type ServiceUpload interface {
}

type services struct {
	repo repositories.Storage
	logz *zap.Logger
}

func NewServices() ServiceUpload {
	loggerZap, err := zap.NewProduction()
	if err != nil {
		panic("cannot init logzap")
	}
	return &services{
		repo: repositories.NewStorage(),
		logz: loggerZap,
	}
}

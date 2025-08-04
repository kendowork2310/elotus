package handlers

import (
	"elotus/cmd/authentication/services"
)

type HandlerAuthen interface {
}

type handlers struct {
	srv services.ServiceAuthen
}

func NewHandler() HandlerAuthen {
	return &handlers{
		srv: services.NewServices(),
	}
}

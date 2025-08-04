package handlers

import "elotus/cmd/upload/services"

type HandlerUpload interface {
}

type handlers struct {
	srv services.ServiceUpload
}

func NewHandler() HandlerUpload {
	return &handlers{
		srv: services.NewServices(),
	}
}

package upload

import (
	"elotus/cmd/upload/handlers"
	"elotus/cmd/upload/middleware"
	"elotus/pkg/cfg"
	"elotus/pkg/server"
	"fmt"
)

const (
	App        = "upload"
	ServerPort = "UPLOAD_SERVER_PORT"
	ServerMode = "SERVER_MODE"
)

type UploadServer struct {
	srv *server.Server
}

func NewUploadServer() *UploadServer {
	serverConfig := newServerConfiguration()
	httpSrv := server.NewHTTPServer(serverConfig)

	r := &UploadServer{
		srv: httpSrv,
	}
	return r
}

func (s *UploadServer) Run() {
	s.route()
	s.srv.Run()
}

func (s *UploadServer) route() {
	handler := handlers.NewHandler()
	// Upload routes with authentication middleware
	s.srv.Route.Use(middleware.TokenAuthentication())
	v1 := s.srv.Route.Group("/v1")
	{
		v1.POST("/upload", handler.UploadFile)
	}
}

func newServerConfiguration() *server.HTTPServerConfiguration {
	return &server.HTTPServerConfiguration{
		Port: cfg.Reader().MustGetString(ServerPort),
		Mode: cfg.Reader().MustGetString(ServerMode),
		App:  fmt.Sprintf("%s", App),
	}
}

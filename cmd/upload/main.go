package upload

import (
	"elotus/cmd/upload/handlers"
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
	_ = handlers.NewHandler()
}

func newServerConfiguration() *server.HTTPServerConfiguration {
	return &server.HTTPServerConfiguration{
		Port: cfg.Reader().MustGetString(ServerPort),
		Mode: cfg.Reader().MustGetString(ServerMode),
		App:  fmt.Sprintf("%s", App),
	}
}

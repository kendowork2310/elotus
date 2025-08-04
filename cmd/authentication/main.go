package authentication

import (
	"elotus/cmd/authentication/handlers"
	"elotus/pkg/cfg"
	"elotus/pkg/server"
	"fmt"
)

const (
	App        = "authentication"
	ServerPort = "SERVER_PORT"
	ServerMode = "SERVER_MODE"
)

type AuthenticationServer struct {
	srv *server.Server
}

func NewAuthenticationServer() *AuthenticationServer {
	serverConfig := newServerConfiguration()
	httpSrv := server.NewHTTPServer(serverConfig)

	r := &AuthenticationServer{
		srv: httpSrv,
	}
	return r
}

func (s *AuthenticationServer) Run() {
	s.route()
	s.srv.Run()
}

func (s *AuthenticationServer) route() {
	_ = handlers.NewHandler()
}

func newServerConfiguration() *server.HTTPServerConfiguration {
	return &server.HTTPServerConfiguration{
		Port: cfg.Reader().MustGetString(ServerPort),
		Mode: cfg.Reader().MustGetString(ServerMode),
		App:  fmt.Sprintf("%s", App),
	}
}

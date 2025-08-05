package authentication

import (
	"elotus/cmd/authentication/handlers"
	"elotus/pkg/cfg"
	"elotus/pkg/server"
	"fmt"
)

const (
	App        = "authentication"
	ServerPort = "AUTHEN_SERVER_PORT"
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
	handler := handlers.NewHandler()
	v1 := s.srv.Route.Group("/v1")
	{
		v1.POST("/register", handler.Register)
		v1.POST("/login", handler.Login)
		v1.POST("/refresh", handler.RefreshToken)
	}
}

func newServerConfiguration() *server.HTTPServerConfiguration {
	return &server.HTTPServerConfiguration{
		Port: cfg.Reader().MustGetString(ServerPort),
		Mode: cfg.Reader().GetStringWithDefault(ServerMode, "debug"),
		App:  fmt.Sprintf("%s", App),
	}
}

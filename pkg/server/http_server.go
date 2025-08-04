package server

import (
	"elotus/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HTTPServerConfiguration struct {
	Port   string
	Mode   string
	Module string
	App    string
}

type Server struct {
	Route *gin.RouterGroup

	server *gin.Engine
	port   string
}

func NewHTTPServer(c *HTTPServerConfiguration) *Server {
	r := gin.New()
	if err := r.SetTrustedProxies(nil); err != nil {
		panic(fmt.Sprintf("could not initialize gin server cause %s", err.Error()))
	}

	r.Use(
		gin.Recovery(),
		logger.RequestInfo(c.App),
		logger.PopulateLogger(),
	)

	rg := r.Group(fmt.Sprintf("%s/%s", c.Module, c.App))
	rg.GET("/status", func(c *gin.Context) {
		c.Status(http.StatusOK)
		return
	})

	return &Server{
		Route:  rg,
		server: r,
		port:   c.Port,
	}
}

func (_this *Server) Run() {
	if err := _this.server.Run(fmt.Sprintf(":%s", _this.port)); err != nil {
		panic(fmt.Sprintf("could not start server cause %s", err))
	}
}

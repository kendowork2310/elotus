package handlers

import (
	"elotus/cmd/authentication/dtos"
	"elotus/cmd/authentication/services"
	"elotus/cmd/common/apif"
	"elotus/cmd/common/errs"
	"elotus/pkg/logger"
	"fmt"

	"github.com/gin-gonic/gin"
)

type HandlerAuthen interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	RefreshToken(c *gin.Context)
}

type handlers struct {
	srv services.ServiceAuthen
}

func NewHandler() HandlerAuthen {
	return &handlers{
		srv: services.NewServices(),
	}
}

func (h *handlers) Register(c *gin.Context) {
	var req dtos.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Gin(c).Error("Failed to bind JSON in Register: %v", err)
		apif.Respond(c, nil, errs.NewCustomError(errs.ErrBadRequest))
		return
	}

	err := h.srv.Register(req.Username, req.Password)
	if err != nil {
		logger.Gin(c).Error("Failed to register user %s: %v", req.Username, err)
	}
	apif.Respond(c, nil, err)
}

func (h *handlers) Login(c *gin.Context) {
	var req dtos.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Gin(c).Error("Failed to bind JSON in Login: %v", err)
		apif.Respond(c, nil, errs.NewCustomError(errs.ErrBadRequest))
		return
	}

	tokenPair, err := h.srv.Login(req.Username, req.Password)
	if err != nil {
		logger.Gin(c).Error("Failed to login user %s: %v", req.Username, err)
		apif.Respond(c, nil, errs.NewCustomError(errs.ErrUnauthorized))
		return
	}

	apif.Respond(c, gin.H{
		"access_token":  tokenPair.AccessToken,
		"refresh_token": tokenPair.RefreshToken,
		"expires_in":    fmt.Sprintf("%ds", tokenPair.ExpiresIn),
		"token_type":    "Bearer",
	}, nil)
}

func (h *handlers) RefreshToken(c *gin.Context) {
	var req dtos.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Gin(c).Error("Failed to bind JSON in RefreshToken: %v", err)
		apif.Respond(c, nil, errs.NewCustomError(errs.ErrBadRequest))
		return
	}

	newAccessToken, err := h.srv.RefreshToken(req.RefreshToken)
	if err != nil {
		logger.Gin(c).Error("Failed to refresh token: %v", err)
		apif.Respond(c, nil, errs.NewCustomError(errs.ErrUnauthorized))
		return
	}

	apif.Respond(c, gin.H{
		"access_token": newAccessToken,
		"expires_in":   fmt.Sprintf("%ds", 15*60), // 15mins
		"token_type":   "Bearer",
	}, nil)
}

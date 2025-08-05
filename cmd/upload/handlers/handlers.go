package handlers

import (
	"elotus/cmd/common/apif"
	"elotus/cmd/common/errs"
	"elotus/cmd/upload/services"
	"elotus/pkg/logger"

	"github.com/gin-gonic/gin"
)

type HandlerUpload interface {
	UploadFile(c *gin.Context)
}

type handlers struct {
	srv services.ServiceUpload
}

func NewHandler() HandlerUpload {
	return &handlers{
		srv: services.NewServices(),
	}
}

func (h *handlers) UploadFile(c *gin.Context) {
	// Get file from form
	file, err := c.FormFile("data")
	if err != nil {
		logger.Gin(c).Error("Failed to get file from form: %v", err)
		apif.Respond(c, nil, errs.NewCustomError(errs.ErrBadRequest))
		return
	}

	// Get username from context (set by auth middleware)
	username, exists := c.Get("username")
	if !exists {
		logger.Gin(c).Error("User not authenticated in upload")
		apif.Respond(c, nil, errs.NewCustomError(errs.ErrUnauthorized))
		return
	}

	// Get user agent
	userAgent := c.GetHeader("User-Agent")

	// Upload file
	upload, err := h.srv.UploadFile(file, username.(string), userAgent)
	if err != nil {
		logger.Gin(c).Error("Failed to upload file for user %s: %v", username, err)
		apif.Respond(c, nil, errs.NewCustomError(errs.ErrBadRequest))
		return
	}

	apif.Respond(c, gin.H{
		"upload": gin.H{
			"id":           upload.ID,
			"filename":     upload.Filename,
			"content_type": upload.ContentType,
			"size":         upload.Size,
			"upload_time":  upload.UploadTime,
			"user":         upload.User,
		},
	}, nil)
}

package services

import (
	"elotus/cmd/common/daos"
	"elotus/cmd/common/errs"
	"elotus/cmd/upload/repositories"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
	"time"
)

const (
	MaxFileSize = 8 * 1024 * 1024
)

type ServiceUpload interface {
	UploadFile(file *multipart.FileHeader, username string, userAgent string) (*daos.Upload, error)
}

type service struct {
	repo repositories.Storage
}

func NewServices() ServiceUpload {
	return &service{
		repo: repositories.NewStorage(),
	}
}

func (s *service) UploadFile(file *multipart.FileHeader, username string, userAgent string) (*daos.Upload, error) {
	// Validate file size
	if file.Size > MaxFileSize {
		return nil, errs.NewCustomError(errs.ErrFileTooLarge)
	}

	// Validate content type
	contentType := file.Header.Get("Content-Type")
	if !isImageContentType(contentType) {
		return nil, errs.NewCustomError(errs.ErrInvalidFileType)
	}

	// Create temporary file
	tempFile, err := os.CreateTemp("/tmp", "upload-*")
	if err != nil {
		return nil, errs.NewCustomErrWithMsg(errs.ErrFileProcessing, fmt.Sprintf("failed to create temp file: %v", err))
	}

	defer tempFile.Close()

	// Open uploaded file
	src, err := file.Open()
	if err != nil {
		return nil, errs.NewCustomErrWithMsg(errs.ErrFileProcessing, fmt.Sprintf("failed to open uploaded file: %v", err))
	}
	defer src.Close()

	// Copy file content to temp file
	_, err = io.Copy(tempFile, src)
	if err != nil {
		return nil, errs.NewCustomErrWithMsg(errs.ErrFileProcessing, fmt.Sprintf("failed to copy file content: %v", err))
	}

	// Create upload record
	upload := &daos.Upload{
		Filename:    file.Filename,
		ContentType: contentType,
		Size:        file.Size,
		UploadTime:  time.Now(),
		User:        username,
		UserAgent:   userAgent,
	}

	err = s.repo.CreateUpload(upload)
	if err != nil {
		return nil, errs.NewCustomErrWithMsg(errs.ErrFileProcessing, fmt.Sprintf("failed to save upload record: %v", err))
	}

	return upload, nil
}

func isImageContentType(contentType string) bool {
	imageTypes := []string{
		"image/jpeg",
		"image/jpg",
		"image/png",
		"image/gif",
		"image/bmp",
		"image/webp",
		"image/tiff",
		"image/svg+xml",
	}

	for _, imageType := range imageTypes {
		if strings.HasPrefix(contentType, imageType) {
			return true
		}
	}
	return false
}

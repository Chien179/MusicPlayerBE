package util

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

type MediaUpload interface {
	FileUpload(file multipart.File, uploadPath string, fileName string) (string, error)
}

type media struct {
	*Config
}

func NewMediaUpload(config *Config) MediaUpload {
	return &media{
		Config: config,
	}
}

func (m *media) FileUpload(file multipart.File, uploadPath string, fileName string) (string, error) {
	url, err := FileUploadHandler(file, uploadPath, fileName, m.Config)

	if err != nil {
		return "", err
	}

	return url, nil
}

func FileUploadHandler(file interface{}, uploadPath string, fileName string, config *Config) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	cld, err := cloudinary.NewFromParams(config.CloudinaryCloudName, config.CloudinaryAPIKey, config.CloudinaryAPISecret)

	if err != nil {
		return "", err
	}

	uploadParam, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{Folder: uploadPath, PublicID: fileName})

	if err != nil {
		return "", err
	}
	return uploadParam.SecureURL, nil
}

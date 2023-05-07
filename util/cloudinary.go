package util

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

type MediaUpload interface {
	FileUpload(file multipart.File, uploadPath string, config Config) (string, error)
}

type media struct{}

func NewMediaUpload() MediaUpload {
	return &media{}
}

func (*media) FileUpload(file multipart.File, uploadPath string, config Config) (string, error) {
	url, err := FileUploadHandler(file, uploadPath, config)

	if err != nil {
		return "", err
	}

	return url, nil
}

func FileUploadHandler(file interface{}, uploadPath string, config Config) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cld, err := cloudinary.NewFromParams(config.CloudinaryCloudName, config.CloudinaryAPIKey, config.CloudinaryAPISecret)

	if err != nil {
		return "", err
	}

	uploadParam, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{Folder: uploadPath})

	if err != nil {
		return "", err
	}
	return uploadParam.SecureURL, nil
}

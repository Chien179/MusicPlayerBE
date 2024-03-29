package api

import (
	"database/sql"
	"errors"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func isForUser(UserID int64, reqUserID int64, ctx *gin.Context) bool {
	if UserID != reqUserID {
		err := errors.New("this resource doesn't belong to the authenticated user")

		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return false
	}

	return true
}

func isGetFieldError(err error, ctx *gin.Context) bool {
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return true
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return true
	}

	return false
}

func isViolationError(err error, ctx *gin.Context) bool {
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusBadRequest, errorResponse(err))
				return true
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return true
	}

	return false
}

func (server *Server) uploadFile(ctx *gin.Context, fileHeader *multipart.FileHeader, filePath string, fileName string) (string, error) {
	file, err := fileHeader.Open()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return "", err
	}

	fileUrl, err := server.uploader.FileUpload(file, filePath, fileName)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return "", err
	}

	return fileUrl, nil
}

package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
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

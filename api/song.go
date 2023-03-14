package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) getSong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, struct{}{})
}

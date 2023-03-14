package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) getPlaylist(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, struct{}{})
}

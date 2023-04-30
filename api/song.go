package api

import (
	"net/http"

	db "github.com/Chien179/MusicPlayerBE/db/sqlc"
	"github.com/gin-gonic/gin"
)

type getSongsRequest struct {
	Page  int32 `form:"page" binding:"required,min=1"`
	Limit int32 `form:"limit" binding:"required,min=10,max=100"`
}

func (server *Server) getSongs(ctx *gin.Context) {
	var req getSongsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.GetSongsParams{
		Limit:  req.Limit,
		Offset: (req.Page - 1) * req.Limit,
	}

	songs, err := server.store.GetSongs(ctx, args)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, songs)
}

func (Server *Server) createSong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, struct{}{})
}

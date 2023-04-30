package api

import (
	"net/http"

	db "github.com/Chien179/MusicPlayerBE/db/sqlc"
	"github.com/gin-gonic/gin"
)

type getGenresRequest struct {
	Page  int32 `form:"page" binding:"required,min=1"`
	Limit int32 `form:"limit" binding:"required,min=5,max=20"`
}

func (server *Server) getGenres(ctx *gin.Context) {
	var req getGenresRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.GetGenresParams{
		Limit:  req.Limit,
		Offset: (req.Page - 1) * req.Limit,
	}

	genres, err := server.store.GetGenres(ctx, args)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, genres)
}

package api

import (
	"mime/multipart"
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

type createSongRequest struct {
	Name     string                `form:"name" binding:"required"`
	Singer   string                `form:"singer" binding:"required"`
	Image    *multipart.FileHeader `form:"image" binding:"required"`
	File     *multipart.FileHeader `form:"file" binding:"required"`
	Duration int64                 `form:"duration" binding:"required,min=1"`
}

func (server *Server) createSong(ctx *gin.Context) {
	var req createSongRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	image, err := req.Image.Open()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	file, err := req.File.Open()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	imgUrl, err := server.uploader.FileUpload(image, "B2CDMusic/Image")

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	fileUrl, err := server.uploader.FileUpload(file, "B2CDMusic/Music")

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	song, err := server.store.CreateSong(ctx, db.CreateSongParams{
		Name:     req.Name,
		Singer:   req.Singer,
		Image:    imgUrl,
		FileUrl:  fileUrl,
		Duration: req.Duration,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, song)
}

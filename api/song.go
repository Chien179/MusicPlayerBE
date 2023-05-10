package api

import (
	"mime/multipart"
	"net/http"
	"strconv"

	db "github.com/Chien179/MusicPlayerBE/db/sqlc"
	"github.com/gin-gonic/gin"
)

func (server *Server) getSongs(ctx *gin.Context) {
	var req getPaginationRequest

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
	Genres   []int64               `form:"genres" binding:"required"`
}

func (server *Server) createSong(ctx *gin.Context) {
	var req createSongRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Create song to get song's id
	song, err := server.store.CreateSongTx(ctx, db.CreateSongTxParams{
		Name:     req.Name,
		Singer:   req.Singer,
		Duration: req.Duration,
		FileUrl:  "",
		Image:    "",
		Genres:   req.Genres,
	})

	if isGetFieldError(err, ctx) {
		return
	}

	var imgUrl, fileUrl string

	if imgUrl, err = server.uploadFile(ctx, req.Image, "B2CDMusic/Image/Songs", strconv.Itoa(int(song.ID))); err != nil {
		return
	}

	if fileUrl, err = server.uploadFile(ctx, req.File, "B2CDMusic/Music", strconv.Itoa(int(song.ID))); err != nil {
		return
	}

	// Update song's file after upload to cloud
	song, err = server.store.UpdateSongFile(ctx, db.UpdateSongFileParams{
		ID:      song.ID,
		Image:   imgUrl,
		FileUrl: fileUrl,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, song)
}

type updateSongRequest struct {
	Name     string                `form:"name"`
	Singer   string                `form:"singer"`
	Image    *multipart.FileHeader `form:"image"`
	File     *multipart.FileHeader `form:"file"`
	Duration int64                 `form:"duration"`
}

func (server *Server) updateSong(ctx *gin.Context) {
	var req idURI

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var body updateSongRequest

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// get song to check if this song existed in db
	song, err := server.store.GetSong(ctx, req.ID)

	if isGetFieldError(err, ctx) {
		return
	}

	// update song req
	songUpdateReq := db.UpdateSongParams{
		ID:       song.ID,
		Name:     body.Name,
		Duration: body.Duration,
		Singer:   body.Singer,
	}

	// check if image file exists, upload new one to cloud and set to update request
	if body.Image != nil {
		var imgUrl string

		if imgUrl, err = server.uploadFile(ctx, body.Image, "B2CDMusic/Image/Songs", strconv.Itoa(int(song.ID))); err != nil {
			return
		}

		songUpdateReq.Image = imgUrl
	} else {
		songUpdateReq.Image = song.Image
	}

	if body.File != nil {
		var fileUrl string

		if fileUrl, err = server.uploadFile(ctx, body.File, "B2CDMusic/Music", strconv.Itoa(int(song.ID))); err != nil {
			return
		}

		songUpdateReq.FileUrl = fileUrl
	} else {
		songUpdateReq.FileUrl = song.FileUrl
	}

	song, _ = server.store.UpdateSong(ctx, songUpdateReq)

	ctx.JSON(http.StatusOK, song)
}

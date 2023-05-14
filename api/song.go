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

	if isViolationError(err, ctx) {
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
	Genres   []int64               `form:"genres"`
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
	songUpdateReq := db.UpdateSongTxParams{
		ID:       song.ID,
		Name:     body.Name,
		Duration: body.Duration,
		Singer:   body.Singer,
		Genres:   body.Genres,
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

	songtx, err := server.store.UpdateSongTx(ctx, songUpdateReq)

	if isViolationError(err, ctx) {
		return
	}

	ctx.JSON(http.StatusOK, songtx)
}

type getPrevOrNextSongURI struct {
	Index *int32 `uri:"index" binding:"required,min=0"`
}

func (server *Server) getPrevOrNextSong(ctx *gin.Context) {
	var req getPrevOrNextSongRequest
	var reqURI getPrevOrNextSongURI

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var song db.Song
	var err error
	var offset int32

	if req.Direction == "prev" {
		if *reqURI.Index == 0 {
			offset = 0
		} else {
			offset = *reqURI.Index - 1
		}
	} else {
		offset = *reqURI.Index + 1
	}

	song, err = server.store.GetSongWithOffset(ctx, offset)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, song)
}

func (server *Server) getSongShuffle(ctx *gin.Context) {
	var req idURI

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	song, err := server.store.GetRandomSong(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, song)
}

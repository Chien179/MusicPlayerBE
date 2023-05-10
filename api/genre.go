package api

import (
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	db "github.com/Chien179/MusicPlayerBE/db/sqlc"
	"github.com/gin-gonic/gin"
)

func (server *Server) getGenres(ctx *gin.Context) {
	var req getPaginationRequest

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

type getGenreResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	Songs     []db.Song `json:"songs"`
}

func (server *Server) getGenre(ctx *gin.Context) {
	var req idURI

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	genre, err := server.store.GetGenre(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	songs, err := server.store.GetGenreSongs(ctx, genre.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := getGenreResponse{
		ID:        genre.ID,
		Name:      genre.Name,
		Image:     genre.Image,
		CreatedAt: genre.CreatedAt,
		Songs:     songs,
	}

	ctx.JSON(http.StatusOK, resp)
}

type createGenreRequest struct {
	Name  string                `form:"name" binding:"required"`
	Image *multipart.FileHeader `form:"image" binding:"required"`
}

func (server *Server) createGenre(ctx *gin.Context) {
	var req createGenreRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	genre, err := server.store.CreateGenre(ctx, db.CreateGenreParams{
		Name: req.Name,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var imgUrl string

	if imgUrl, err = server.uploadFile(ctx, req.Image, "B2CDMusic/Image/Genres", strconv.Itoa(int(genre.ID))); err != nil {
		return
	}

	genre, err = server.store.Updategenre(ctx, db.UpdategenreParams{
		ID:    genre.ID,
		Name:  genre.Name,
		Image: imgUrl,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, genre)
}

type updateGenreRequest struct {
	Name  string                `form:"name"`
	Image *multipart.FileHeader `form:"image"`
}

func (server *Server) updateGenre(ctx *gin.Context) {
	var req idURI

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var body updateGenreRequest

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	genre, err := server.store.GetGenre(ctx, req.ID)

	if isGetFieldError(err, ctx) {
		return
	}

	genreUpdateReq := db.UpdategenreParams{
		ID:   genre.ID,
		Name: body.Name,
	}

	if body.Image != nil {
		var imgUrl string

		if imgUrl, err = server.uploadFile(ctx, body.Image, "B2CDMusic/Image/Genres", strconv.Itoa(int(genre.ID))); err != nil {
			return
		}

		genreUpdateReq.Image = imgUrl
	} else {
		genreUpdateReq.Image = genre.Image
	}

	genre, _ = server.store.Updategenre(ctx, genreUpdateReq)

	ctx.JSON(http.StatusOK, genre)
}

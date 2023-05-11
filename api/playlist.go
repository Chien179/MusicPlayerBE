package api

import (
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	db "github.com/Chien179/MusicPlayerBE/db/sqlc"
	"github.com/Chien179/MusicPlayerBE/token"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (server *Server) getUserPlaylists(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	playlists, err := server.store.GetUserPlaylists(ctx, authPayload.UserID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, playlists)
}

type playlistDetailResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	Songs     []db.Song `json:"songs"`
}

func (server *Server) getUserPlaylistDetail(ctx *gin.Context) {
	var req idURI

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	playlist, err := server.store.GetUserPlaylist(ctx, req.ID)

	if isGetFieldError(err, ctx) {
		return
	}

	if !isForUser(playlist.UsersID, authPayload.UserID, ctx) {
		return
	}

	songs, err := server.store.GetUserPlaylistSongs(ctx, playlist.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := playlistDetailResponse{
		ID:        playlist.ID,
		Name:      playlist.Name,
		Image:     playlist.Image,
		CreatedAt: playlist.CreatedAt,
		Songs:     songs,
	}

	ctx.JSON(http.StatusOK, res)
}

type songAndPlaylistUri struct {
	ID     int64 `uri:"id" binding:"required,min=1"`
	SongID int64 `uri:"song_id" binding:"required,min=1"`
}

func (server *Server) addSongToPlaylist(ctx *gin.Context) {
	var req songAndPlaylistUri

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	playlist, err := server.store.GetUserPlaylist(ctx, req.ID)

	if isGetFieldError(err, ctx) {
		return
	}

	if !isForUser(playlist.UsersID, authPayload.UserID, ctx) {
		return
	}

	song, err := server.store.GetSong(ctx, req.SongID)

	if isGetFieldError(err, ctx) {
		return
	}

	res, err := server.store.AddSongToPlaylist(ctx, db.AddSongToPlaylistParams{
		PlaylistsID: playlist.ID,
		SongsID:     song.ID,
	})

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (server *Server) removeSongFromPlaylist(ctx *gin.Context) {
	var req songAndPlaylistUri

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	playlist, err := server.store.GetUserPlaylist(ctx, req.ID)

	if isGetFieldError(err, ctx) {
		return
	}

	if !isForUser(playlist.UsersID, authPayload.UserID, ctx) {
		return
	}

	song, err := server.store.GetSong(ctx, req.SongID)

	if isGetFieldError(err, ctx) {
		return
	}

	err = server.store.RemoveSongFromPlaylist(ctx, db.RemoveSongFromPlaylistParams{
		PlaylistsID: playlist.ID,
		SongsID:     song.ID,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}

type createPlaylistRequest struct {
	Name  string `json:"name" binding:"required"`
	Image string `json:"image" binding:"required"`
}

func (server *Server) createUserPlaylist(ctx *gin.Context) {
	var req createPlaylistRequest

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	playlist, err := server.store.CreatePlaylist(ctx, db.CreatePlaylistParams{
		Name:    req.Name,
		UsersID: authPayload.UserID,
		Image:   req.Image,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, playlist)
}

type updatePlaylistRequest struct {
	Name  string                `form:"name"`
	Image *multipart.FileHeader `form:"image"`
}

func (server *Server) updateUserPlaylist(ctx *gin.Context) {
	var req idURI
	var body updatePlaylistRequest

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	playlist, err := server.store.GetUserPlaylist(ctx, req.ID)

	if isGetFieldError(err, ctx) {
		return
	}

	if !isForUser(playlist.UsersID, authPayload.UserID, ctx) {
		return
	}

	updateReq := db.UpdatePlaylistParams{
		ID:   playlist.ID,
		Name: body.Name,
	}

	if body.Image != nil {
		var imgUrl string

		if imgUrl, err = server.uploadFile(ctx, body.Image, "B2CDMusic/Image/Playlists", strconv.Itoa(int(playlist.ID))); err != nil {
			return
		}

		updateReq.Image = imgUrl
	}

	playlist, err = server.store.UpdatePlaylist(ctx, updateReq)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, playlist)
}

func (server *Server) deleteUserPlaylist(ctx *gin.Context) {
	var req idURI

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	playlist, err := server.store.GetUserPlaylist(ctx, req.ID)

	if isGetFieldError(err, ctx) {
		return
	}

	if !isForUser(playlist.UsersID, authPayload.UserID, ctx) {
		return
	}

	err = server.store.DeletePlaylist(ctx, playlist.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}

package api

import (
	"github.com/gin-gonic/gin"
)

func (server *Server) setupRouter() {
	router := gin.Default()
	router.Use(CORSMiddleware())

	// No-auth
	router.POST("/register", server.register)
	router.POST("/login", server.login)
	router.GET("/songs", server.getSongs)
	router.GET("/genres", server.getGenres)
	router.GET("/genres/:id", server.getGenre)

	// Require auth
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("/playlists", server.getUserPlaylists)
	authRoutes.GET("/playlists/:id", server.getUserPlaylistDetail)
	authRoutes.POST("/playlists", server.createUserPlaylist)
	authRoutes.PUT("/playlists/:id", server.updateUserPlaylist)
	authRoutes.DELETE("/playlists/:id", server.deleteUserPlaylist)
	authRoutes.POST("/playlists/:id/songs/:song_id", server.addSongToPlaylist)
	authRoutes.DELETE("/playlists/:id/songs/:song_id", server.removeSongFromPlaylist)

	// Admin only
	adminRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker), adminAuthorizeMiddleware)
	adminRoutes.POST("/songs", server.createSong)
	adminRoutes.PUT("/songs/:id", server.updateSong)
	adminRoutes.POST("/genres", server.createGenre)
	adminRoutes.PUT("/genres/:id", server.updateGenre)

	server.router = router
}

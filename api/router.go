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
	router.GET("/songs/index/:index", server.getPrevOrNextSong)
	router.GET("/songs/shuffle/:id", server.getSongShuffle)
	router.GET("/genres", server.getGenres)
	router.GET("/genres/:id", server.getGenre)

	// Require auth
	authRoutes := router.Group("/playlists").Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("", server.getUserPlaylists)
	authRoutes.GET("/:id", server.getUserPlaylistDetail)
	authRoutes.GET("/:id/songs/index/:index", server.getPrevOrNextPlaylistSong)
	authRoutes.GET("/:id/songs/shuffle/:song_id", server.getPlaylistSongShuffle)
	authRoutes.POST("", server.createUserPlaylist)
	authRoutes.PUT("/:id", server.updateUserPlaylist)
	authRoutes.DELETE("/:id", server.deleteUserPlaylist)
	authRoutes.POST("/:id/songs/:song_id", server.addSongToPlaylist)
	authRoutes.DELETE("/:id/songs/:song_id", server.removeSongFromPlaylist)

	userRoutes := router.Group("/user").Use(authMiddleware(server.tokenMaker))
	userRoutes.GET("", server.getUser)
	userRoutes.PUT("", server.updateUser)
	userRoutes.DELETE("", server.deleteUser)

	// Admin only
	adminRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker), adminAuthorizeMiddleware)
	adminRoutes.POST("/songs", server.createSong)
	adminRoutes.PUT("/songs/:id", server.updateSong)
	adminRoutes.POST("/genres", server.createGenre)
	adminRoutes.PUT("/genres/:id", server.updateGenre)
	adminRoutes.GET("/users", server.getUsers)

	server.router = router
}

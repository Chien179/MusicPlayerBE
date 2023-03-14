package api

import "github.com/gin-gonic/gin"

func (server *Server) setupRouter() {
	router := gin.Default()

	// No-auth
	router.POST("/register", server.register)
	router.POST("/login", server.login)

	// Require auth
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("/playlists", server.getPlaylist)

	// Admin only
	adminRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker), adminAuthorizeMiddleware)
	adminRoutes.GET("/songs", server.getSong)

	server.router = router
}

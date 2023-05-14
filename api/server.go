package api

import (
	"fmt"

	db "github.com/Chien179/MusicPlayerBE/db/sqlc"
	"github.com/Chien179/MusicPlayerBE/token"
	"github.com/Chien179/MusicPlayerBE/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves all HTTP request for Music APIs
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
	uploader   util.MediaUpload
}

// NewServer creates a new server
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
		uploader:   util.NewMediaUpload(&config),
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("direction", validDirection)
	}

	server.setupRouter()

	return server, nil
}

// Start runs HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

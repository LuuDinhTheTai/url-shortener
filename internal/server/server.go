package server

import (
	"url-shortener/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	r   *gin.Engine
	cfg config.Config
	db  *gorm.DB
}

func NewServer(cfg config.Config, db *gorm.DB) *Server {
	return &Server{
		cfg: cfg,
		db:  db,
	}
}

func (s *Server) Run() error {
	s.r = gin.Default()

	if err := s.MapHandlers(); err != nil {
		return err
	}

	if err := s.r.Run(":" + s.cfg.Server.Port); err != nil {
		return err
	}

	return nil
}

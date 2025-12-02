package server

import (
	"url-shortener/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	r   *gin.Engine
	cfg config.Config
}

func NewServer(cfg config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Run() error {
	s.r = gin.Default()

	if err := s.MapHandlers(s.r); err != nil {
		return err
	}

	if err := s.r.Run(":" + s.cfg.Server.Port); err != nil {
		return err
	}

	return nil
}

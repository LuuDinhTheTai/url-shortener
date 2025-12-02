package server

import (
	"url-shortener/internal/url/delivery/http"
	"url-shortener/internal/url/repository"
	"url-shortener/internal/url/usecase"
)

func (s *Server) MapHandlers() error {
	// Init repository
	urlRepository := repository.NewUrlRepository(s.db)

	// Init usecase
	urlUseCase := usecase.NewUrlUseCase(s.cfg, urlRepository)

	// Init handler
	urlHandler := http.NewUrlHandler(urlUseCase)

	http.MapUrlRoutes(s.r, urlHandler)

	return nil
}

package server

import (
	"url-shortener/internal/url/delivery/http"
	"url-shortener/internal/url/repository"
	"url-shortener/internal/url/usecase"

	"github.com/gin-gonic/gin"
)

func (s *Server) MapHandlers(r *gin.Engine) error {
	// Init repository
	urlRepository := repository.NewUrlRepository()

	// Init usecase
	urlUseCase := usecase.NewUrlUseCase(urlRepository)

	// Init handler
	urlHandler := http.NewUrlHandler(urlUseCase)

	http.MapUrlRoutes(r, urlHandler)

	return nil
}

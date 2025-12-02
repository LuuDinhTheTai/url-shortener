package usecase

import (
	"url-shortener/internal/url"

	"github.com/gin-gonic/gin"
)

type urlUseCase struct {
	urlRepository url.Repository
}

func NewUrlUseCase(urlRepository url.Repository) url.UseCase {
	return &urlUseCase{
		urlRepository: urlRepository,
	}
}

func (u *urlUseCase) Shorten(ctx *gin.Context, url string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (u *urlUseCase) Redirect(ctx *gin.Context, code string) (string, error) {
	//TODO implement me
	panic("implement me")
}

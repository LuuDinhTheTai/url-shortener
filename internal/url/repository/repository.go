package repository

import (
	"url-shortener/internal/models"
	"url-shortener/internal/url"

	"github.com/gin-gonic/gin"
)

type urlRepository struct {
}

func NewUrlRepository() url.Repository {
	return &urlRepository{}
}

func (u *urlRepository) Create(ctx *gin.Context, url models.Url) (*models.Url, error) {
	//TODO implement me
	panic("implement me")
}

func (u *urlRepository) FindByCode(ctx *gin.Context, code string) (*models.Url, error) {
	//TODO implement me
	panic("implement me")
}

package repository

import (
	"url-shortener/internal/model"
	"url-shortener/internal/url"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type urlRepository struct {
	db *gorm.DB
}

func NewUrlRepository(db *gorm.DB) url.Repository {
	return &urlRepository{
		db: db,
	}
}

func (u *urlRepository) Save(ctx *gin.Context, url model.Url) (*model.Url, error) {
	if err := u.db.WithContext(ctx).Create(&url).Error; err != nil {
		return nil, err
	}

	return &url, nil
}

func (u *urlRepository) FindByCode(ctx *gin.Context, code string) (*model.Url, error) {
	//TODO implement me
	panic("implement me")
}

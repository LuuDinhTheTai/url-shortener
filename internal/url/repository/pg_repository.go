package repository

import (
	"context"
	"fmt"
	"url-shortener/internal/model"
	"url-shortener/internal/url/interface"

	"gorm.io/gorm"
)

type urlRepository struct {
	db *gorm.DB
}

func NewUrlRepository(db *gorm.DB) _interface.Repository {
	return &urlRepository{
		db: db,
	}
}

func (u *urlRepository) Save(ctx context.Context, url model.Url) (*model.Url, error) {
	if err := u.db.WithContext(ctx).Create(&url).Error; err != nil {
		return nil, fmt.Errorf("save: failed to save url: %w\n", err)
	}

	return &url, nil
}

func (u *urlRepository) FindByCode(ctx context.Context, code string) (*model.Url, error) {
	//TODO implement me
	panic("implement me")
}

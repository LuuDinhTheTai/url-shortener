package _interface

import (
	"context"
	"url-shortener/internal/model"
)

type Repository interface {
	Save(ctx context.Context, url model.Url) (*model.Url, error)
	FindByCode(ctx context.Context, code string) (*model.Url, error)
	Update(ctx context.Context, newUrl model.Url) (*model.Url, error)
}

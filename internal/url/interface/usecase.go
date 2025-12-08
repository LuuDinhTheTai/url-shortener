package _interface

import (
	"context"
	"url-shortener/internal/url/delivery/dto"
)

type UseCase interface {
	Shorten(ctx context.Context, url string) (*dto.ShortenResponse, error)
	Redirect(ctx context.Context, code string) (string, error)
}

package usecase

import (
	"context"
	"fmt"
	"time"
	"url-shortener/config"
	"url-shortener/internal/encrypt"
	"url-shortener/internal/model"
	"url-shortener/internal/url/delivery/dto"
	"url-shortener/internal/url/interface"

	"github.com/google/uuid"
)

type urlUseCase struct {
	cfg           config.Config
	urlRepository _interface.Repository
}

func NewUrlUseCase(cfg config.Config, urlRepository _interface.Repository) _interface.UseCase {
	return &urlUseCase{
		cfg:           cfg,
		urlRepository: urlRepository,
	}
}

func (u *urlUseCase) Shorten(ctx context.Context, url string) (*dto.ShortenResponse, error) {
	newId := uuid.New()

	code := encrypt.EncodeBase62(newId)

	domain := u.cfg.Server.Host + u.cfg.Server.Port

	newUrl := &model.Url{
		Id:        newId,
		Domain:    domain,
		Code:      code,
		LongUrl:   url,
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}

	result, err := u.urlRepository.Save(ctx, *newUrl)
	if err != nil {
		return &dto.ShortenResponse{}, fmt.Errorf("shorten: failed to save url: %w\n", err)
	}

	return &dto.ShortenResponse{
		ShortUrl:  result.Domain + "/" + result.Code,
		LongUrl:   result.LongUrl,
		ExpiresAt: result.ExpiresAt,
	}, nil
}

func (u *urlUseCase) Redirect(ctx context.Context, code string) (string, error) {
	url, err := u.urlRepository.FindByCode(ctx, code)
	if err != nil {
		return "", fmt.Errorf("redirect: failed to find url: %w\n", err)
	}

	if url.ExpiresAt.Before(time.Now()) {
		return "", fmt.Errorf("redirect: url expired")
	}

	if url.LongUrl == "" {
		return "", fmt.Errorf("redirect: url not found")
	}

	url.UpdatedAt = time.Now()
	if _, err := u.urlRepository.Update(ctx, *url); err != nil {
		return "", fmt.Errorf("redirect: failed to update url: %w\n", err)
	}

	return url.LongUrl, nil
}

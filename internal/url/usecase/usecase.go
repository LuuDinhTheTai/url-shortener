package usecase

import (
	"log"
	"time"
	"url-shortener/config"
	"url-shortener/internal/encrypt"
	"url-shortener/internal/model"
	"url-shortener/internal/url"
	"url-shortener/internal/url/delivery/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type urlUseCase struct {
	cfg           config.Config
	urlRepository url.Repository
}

func NewUrlUseCase(cfg config.Config, urlRepository url.Repository) url.UseCase {
	return &urlUseCase{
		cfg:           cfg,
		urlRepository: urlRepository,
	}
}

func (u *urlUseCase) Shorten(ctx *gin.Context, url string) (*dto.ShortenResponse, error) {
	newId := uuid.New()

	code := encrypt.EncodeBase62(newId)

	domain := u.cfg.Server.Host

	newUrl := &model.Url{
		Id:        newId,
		Domain:    domain,
		Code:      code,
		LongUrl:   url,
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}

	result, err := u.urlRepository.Save(ctx, *newUrl)
	if err != nil {
		log.Println(err)
		return &dto.ShortenResponse{}, err
	}

	return &dto.ShortenResponse{
		ShortUrl:  result.Domain + "/" + result.Code,
		LongUrl:   result.LongUrl,
		ExpiresAt: result.ExpiresAt,
	}, nil
}

func (u *urlUseCase) Redirect(ctx *gin.Context, code string) (string, error) {
	//TODO implement me
	panic("implement me")
}

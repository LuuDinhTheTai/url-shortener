package _interface

import (
	"url-shortener/internal/url/delivery/dto"

	"github.com/gin-gonic/gin"
)

type UseCase interface {
	Shorten(ctx *gin.Context, url string) (*dto.ShortenResponse, error)
	Redirect(ctx *gin.Context, code string) (string, error)
}

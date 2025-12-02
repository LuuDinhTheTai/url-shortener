package http

import (
	"net/http"
	"url-shortener/internal/url"

	"github.com/gin-gonic/gin"
)

type urlHandler struct {
	urlUseCase url.UseCase
}

func NewUrlHandler(urlUseCase url.UseCase) url.Handler {
	return &urlHandler{
		urlUseCase: urlUseCase,
	}
}

func (u *urlHandler) Shorten(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"result": "shortened url",
	})
}

func (u *urlHandler) Redirect(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

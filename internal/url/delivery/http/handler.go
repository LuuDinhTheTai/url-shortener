package http

import (
	"log"
	"net/http"
	"url-shortener/internal/url/delivery/dto"
	"url-shortener/internal/url/interface"

	"github.com/gin-gonic/gin"
)

type urlHandler struct {
	urlUseCase _interface.UseCase
}

func NewUrlHandler(urlUseCase _interface.UseCase) _interface.Handler {
	return &urlHandler{
		urlUseCase: urlUseCase,
	}
}

func (u *urlHandler) Shorten(ctx *gin.Context) {
	var req dto.ShortenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := u.urlUseCase.Shorten(ctx, req.LongURL)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (u *urlHandler) Redirect(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

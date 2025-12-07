package http

import (
	"log"
	"net/http"
	"strings"
	"url-shortener/internal/url"
	"url-shortener/internal/url/delivery/dto"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	code := ctx.Param("code")

	if strings.HasSuffix(code, "favicon.ico") {
		ctx.Status(http.StatusNoContent)
		return
	}

	result, err := u.urlUseCase.Redirect(ctx, code)
	if err != nil {
		log.Println(err)

		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "short url not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, result)
}

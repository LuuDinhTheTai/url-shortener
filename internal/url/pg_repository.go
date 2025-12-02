package url

import (
	"url-shortener/internal/model"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	Save(ctx *gin.Context, url model.Url) (*model.Url, error)
	FindByCode(ctx *gin.Context, code string) (*model.Url, error)
}

package url

import (
	"url-shortener/internal/models"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	Create(ctx *gin.Context, url models.Url) (*models.Url, error)
	FindByCode(ctx *gin.Context, code string) (*models.Url, error)
}

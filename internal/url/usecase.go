package url

import "github.com/gin-gonic/gin"

type UseCase interface {
	Shorten(ctx *gin.Context, url string) (string, error)
	Redirect(ctx *gin.Context, code string) (string, error)
}

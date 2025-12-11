package http

import (
	"url-shortener/internal/url/interface"

	"github.com/gin-gonic/gin"
)

func MapUrlRoutes(r *gin.Engine, handler _interface.Handler) {
	r.POST("/shorten", handler.Shorten)
	r.GET("/:code", handler.Redirect)
}

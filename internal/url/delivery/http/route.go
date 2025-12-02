package http

import (
	"url-shortener/internal/url"

	"github.com/gin-gonic/gin"
)

func MapUrlRoutes(r *gin.Engine, handler url.Handler) {
	r.POST("/shorten", handler.Shorten)
	r.GET("/:code", handler.Redirect)
}

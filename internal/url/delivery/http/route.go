package http

import (
	"url-shortener/internal/url"

	"github.com/gin-gonic/gin"
)

func MapUrlRoutes(r *gin.Engine, handler url.Handler) {
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(204) // 204 No Content
	})
	r.POST("/shorten", handler.Shorten)
	r.GET("/:code", handler.Redirect)
}

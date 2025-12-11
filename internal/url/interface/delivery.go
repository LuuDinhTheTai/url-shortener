package _interface

import "github.com/gin-gonic/gin"

type Handler interface {
	Shorten(ctx *gin.Context)
	Redirect(ctx *gin.Context)
}

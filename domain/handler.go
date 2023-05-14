package domain

import "github.com/gin-gonic/gin"

type Handler interface {
	Register(router *gin.Engine)
}

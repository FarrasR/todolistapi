package router

import (
	"todolistapi/domain"
	"todolistapi/entity/response"

	"github.com/gin-gonic/gin"
)

func BuildHandler(handlers ...domain.Handler) *gin.Engine {
	router := gin.New()

	for _, handler := range handlers {
		handler.Register(router)
	}
	return router
}

func RunServer(router *gin.Engine) {
	router.RedirectFixedPath = true
	router.HandleMethodNotAllowed = true

	router.GET("/", func(c *gin.Context) {
		healthCheck(c)
	})

	err := router.Run(":3030")
	if err != nil {
		panic("Error To Start")
	}
}

func healthCheck(c *gin.Context) {
	response.OKHelloWorld(c)
}

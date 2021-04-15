package api

import (
	"github.com/gin-gonic/gin"
	"oi.io/apps/discover/handler"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.POST("api/register", handler.RegisterHandler)
	return router
}

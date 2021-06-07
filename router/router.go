package router

import (
	"PrintCloud/api"
	_ "PrintCloud/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Http(router *gin.Engine) {
	printerRouter := router.Group("/api/printers/")
	{
		printerRouter.GET("/", api.Printers)
	}
	swaggerRouter := router.Group("/api/swagger/")
	{
		swaggerRouter.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

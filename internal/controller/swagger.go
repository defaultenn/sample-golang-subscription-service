package controller

import (
	"test_task/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitSwagger(engine *gin.Engine) {
	docs.SwaggerInfo.Title = "Subscription Service"
	docs.SwaggerInfo.Description = "Subscription Service"
	docs.SwaggerInfo.Version = "0.1.0"
	engine.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER"))
}

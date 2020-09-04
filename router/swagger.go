package router

import (
	"github.com/gin-gonic/gin"
	"github.com/marshhu/ma-novel-crawler/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func useSwaggerRouters(eng *gin.Engine) {
	docs.SwaggerInfo.Title = "github.com/marshhu/ma-novel-crawler API"
	docs.SwaggerInfo.Description = "ma小说爬虫."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	eng.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

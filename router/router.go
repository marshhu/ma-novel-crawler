package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// Creates a router without any middleware by default
	r := gin.New()

	r.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	userBizRouters(r)
	useSwaggerRouters(r)

	gin.SetMode("debug")
	return r
}

//// useStaticRouters 静态资源
//func useStaticRouters(eng *gin.Engine) {
//	staticGroup := eng.Group("/static")
//	rootDir := utils.RootDir()
//	staticGroup.Static("", filepath.Join(rootDir, "assets"))
//}

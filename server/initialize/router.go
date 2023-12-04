package initialize

import (
	"net/http"
	"quickSearch/global"
	r "quickSearch/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	router := gin.Default()
	codeMemoryRouter := r.RouterGroupApp.System
	//跨域
	router.Use(Cors())
	//未找到路由
	router.NoRoute(NotFound)
	APIGroup := router.Group(global.CONFIG.System.RouterPrefix)
	{
		codeMemoryRouter.InitFileInfoRouter(APIGroup)
		codeMemoryRouter.InitInfoRouter(APIGroup)
		codeMemoryRouter.InitCategoryRouter(APIGroup)
	}
	return router
}
func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"error":  "404 ,page not exists!",
	})
}

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		origin := context.GetHeader("Origin")
		context.Header("Access-Control-Allow-Origin", origin)
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}

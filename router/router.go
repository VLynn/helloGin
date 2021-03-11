package router

import (
	"github.com/gin-gonic/gin"
	"helloGin/controller"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("view/*")

	// 注册
	router.GET("/register", controller.RegisterGet)
	router.POST("/register", controller.RegisterPost)

	return router
}

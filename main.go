package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"net/http"
)

func main() {
	// 创建一个服务
	ginServer := gin.Default()
	ginServer.Use(favicon.New("./favicon.ico"))

	// 加载静态页面
	ginServer.LoadHTMLGlob("templates/*")
	// 加载资源文件
	ginServer.Static("static", "./static")

	// 访问地址，处理我们的请求   Request   Respose
	ginServer.GET("/index", func(context *gin.Context) {
		// context.JSON()
		context.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "这是go后台的数据",
		})
	})

	// url?urlid=xxx
	ginServer.GET("user/info", func(context *gin.Context) {
		userid := context.Query("userid")
		username := context.Query("username")

		context.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})
	})

	// /user/info/1/yangwb
	ginServer.GET("/user/info/:userid/:username", func(context *gin.Context) {
		userid := context.Param("userid")
		username := context.Param("username")
		context.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})
	})

	// 支持函数式编程
	ginServer.POST("/user/add", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		context.JSON(http.StatusOK, gin.H{
			"msg":      "ok",
			"username": username,
			"password": password,
		})
	})

	// ginServer

	// 服务器端口
	ginServer.Run(":8082")
}

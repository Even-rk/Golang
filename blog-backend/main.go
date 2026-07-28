package main

import (
	"blog-backend/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化 Gin 路由
	r := gin.Default()

	// 注册路由
	router.RegisterRoutes(r)

	// 启动服务器
	r.Run(":8080")
}

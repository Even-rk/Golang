package main

import (
	"blog-backend/router"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 链接数据库
	db, err := gorm.Open(mysql.Open("root:102466@tcp(127.0.0.1:3306)/gorm"), &gorm.Config{})
	if err != nil {
		fmt.Printf("数据库连接失败: %v", err)
	}
	fmt.Println("数据库连接成功", db)
	// 初始化 Gin 路由
	r := gin.New()
	// 注册路由
	router.RegisterRoutes(r, db)
	// 启动服务器
	r.Run(":8080")
}

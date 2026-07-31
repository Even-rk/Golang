package main

import (
	"blog-backend/router"
	model "blog-backend/types"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GetDBURL 获取数据库连接URL
func GetDBURL() string {
	return os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME")
}

func main() {
	// 获取当前环境，默认开发环境
	appEnv := os.Getenv("MODE")
	// 根据当前环境加载对应的配置文件
	envFile := ".env." + appEnv
	err := godotenv.Load(envFile)

	// 链接数据库
	db, err := gorm.Open(mysql.Open(GetDBURL()), &gorm.Config{})
	if err != nil {
		fmt.Printf("数据库连接失败: %v", err)
	}
	fmt.Println("数据库连接成功", db)
	// 自动迁移模型
	db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
	// 初始化 Gin 路由
	r := gin.New()
	// 注册路由
	router.RegisterRoutes(r, db)
	// 启动服务器
	r.Run(":" + os.Getenv("SERVER_PORT"))
}

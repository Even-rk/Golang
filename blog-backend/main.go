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
	DBuser := os.Getenv("DB_USER")
	DBpassword := os.Getenv("DB_PASSWORD")
	DBhost := os.Getenv("DB_HOST")
	DBport := os.Getenv("DB_PORT")
	DBname := os.Getenv("DB_NAME")
	dsn := "charset=utf8mb4&parseTime=true&loc=Local"
	return DBuser + ":" + DBpassword + "@tcp(" + DBhost + ":" + DBport + ")/" + DBname + "?" + dsn
}

func main() {
	// 获取当前环境，默认开发环境
	appEnv := os.Getenv("MODE")
	if appEnv == "" {
		appEnv = "development"
	}
	// 根据当前环境加载对应的配置文件
	envFile := ".env." + appEnv
	err := godotenv.Load(envFile)
	if err != nil {
		fmt.Printf("加载配置文件失败: %v\n", err)
		return
	}

	// 链接数据库
	db, err := gorm.Open(mysql.Open(GetDBURL()), &gorm.Config{})
	if err != nil {
		fmt.Printf("数据库连接失败: %v", err)
		return
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

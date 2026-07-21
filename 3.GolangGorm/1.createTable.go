package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 模型定义
// 创建user模型：用户模型，包含用户ID、用户名、邮箱字段
type User struct {
	ID    int    `gorm:"primaryKey"`
	Name  string `gorm:"column:name"`
	Email string `gorm:"column:email"`
}

// 创建post模型 ：文章模型，包含文章ID、标题、内容、用户ID字段
type Post struct {
	ID      int    `gorm:"primaryKey"`
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
	UserID  int    `gorm:"column:user_id"`
}

// 创建comment模型 ：评论模型，包含评论ID、文章ID、用户ID、内容字段
type Comment struct {
	ID      int    `gorm:"primaryKey"`
	PostID  int    `gorm:"column:post_id"`
	UserID  int    `gorm:"column:user_id"`
	Content string `gorm:"column:content"`
}

func CreateTable() *gorm.DB {
	// 初始化数据库连接
	db, err := gorm.Open(mysql.Open("root:102466@tcp(127.0.0.1:3306)/gorm"), &gorm.Config{})
	if err != nil {
		fmt.Printf("数据库连接失败: %v", err)
	}
	// 创建数据库表
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	// 清空数据
	db.Exec("TRUNCATE TABLE users")
	db.Exec("TRUNCATE TABLE posts")
	db.Exec("TRUNCATE TABLE comments")
	// 创建实例对应User模型
	// var users []User
	// 添加多个用户到数据库
	var userList = []User{
		{Name: "张三", Email: "zhangsan@example.com"},
		{Name: "李四", Email: "lisi@example.com"},
	}
	db.Create(&userList)
	// 创建实例对应Post模型
	// var posts []Post
	// 添加多个文章到数据库
	var postList = []Post{
		{Title: "第一篇文章", Content: "这是用户1的第一篇文章的内容", UserID: 1},
		{Title: "第二篇文章", Content: "这是用户2的第二篇文章的内容", UserID: 2},
		{Title: "第三篇文章", Content: "这是用户1的第二篇文章的内容", UserID: 1},
		{Title: "第四篇文章", Content: "这是用户1的第三篇文章的内容", UserID: 1},
		{Title: "第五篇文章", Content: "这是用户1的第四篇文章的内容", UserID: 1},
	}
	db.Create(&postList)
	// 创建实例对应Comment模型
	// var comments []Comment
	// 添加多个评论到数据库
	var commentList = []Comment{
		{PostID: 1, UserID: 1, Content: "这是一篇文章的第一条评论"},
		{PostID: 2, UserID: 1, Content: "这是第二篇文章的第一条评论"},
		{PostID: 3, UserID: 1, Content: "这是第三篇文章的第一条评论"},
		{PostID: 4, UserID: 2, Content: "这是第四篇文章的第一条评论"},
		{PostID: 5, UserID: 2, Content: "这是第五篇文章的第一条评论"},
		{PostID: 1, UserID: 2, Content: "这是第一篇文章的第二条评论"},
		{PostID: 2, UserID: 2, Content: "这是第二篇文章的第二条评论"},
		{PostID: 3, UserID: 2, Content: "这是第三篇文章的第二条评论"},
		{PostID: 3, UserID: 2, Content: "这是第三篇文章的第三条评论"},
		{PostID: 3, UserID: 2, Content: "这是第三篇文章的第四条评论"},
	}
	db.Create(&commentList)
	// // 查询用户表
	// db.Find(&users)
	// // 查询文章表
	// db.Find(&posts)
	// // 查询评论表
	// db.Find(&comments)
	// // 打印查询结果
	// fmt.Printf("查询到的用户:%v", users)
	// fmt.Printf("查询到的文章:%v", posts)
	// fmt.Printf("查询到的评论:%v", comments)

	return db
}

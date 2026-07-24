package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 1. 初始化并创建数据库表
	db, err := gorm.Open(mysql.Open("root:102466@tcp(127.0.0.1:3306)/gorm"), &gorm.Config{})
	if err != nil {
		fmt.Printf("数据库连接失败: %v", err)
	}
	// 创建数据
	CreateTable(db)
	// 2.1 查询某个用户发布的所有文章及其对应的评论信息
	userList := []User{}
	db.Find(&userList)
	for _, user := range userList {
		postsWithComments := QueryUserPostsAndComments(db, user.ID)
		// 用户信息
		fmt.Printf("用户ID：%d, 用户名：%s\n", user.ID, user.Name)
		for _, post := range postsWithComments {
			// 文章信息
			fmt.Printf("文章ID：%d, 标题：%s, 内容：%s\n", post.Post.ID, post.Post.Title, post.Post.Content)
			// 评论信息
			for _, comment := range post.Comments {
				fmt.Printf("评论ID：%d, 内容：%s\n", comment.ID, comment.Content)
			}
			// 文章分隔线
			fmt.Println("-----------------")
		}
		// 用户分隔线
		fmt.Println("==================================")
	}

	// 2.2 查询评论数量最多的文章信息
	post := QueryMostCommentedPost(db)
	// 评论数量最多的文章信息
	fmt.Printf("评论数量最多的文章ID：%d, 标题：%s, 内容：%s\n", post.ID, post.Title, post.Content)
	// 3.删除评论
	var comment Comment
	db.Model(&Comment{}).Limit(1).First(&comment)
	// 删除评论
	db.Unscoped().Where("id = ?", comment.ID).Delete(&comment)
}

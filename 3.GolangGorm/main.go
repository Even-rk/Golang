package main

import "fmt"

func main() {
	// 1. 初始化并创建数据库表
	db := CreateTable()

	// 2.1 查询某个用户发布的所有文章及其对应的评论信息
	userList := []User{}
	db.Find(&userList)
	for _, user := range userList {
		postsWithComments := QueryUserPostsAndComments(db, user.ID)
		for _, post := range postsWithComments {
			// 文章信息
			fmt.Printf("文章ID：%d, 标题：%s, 内容：%s\n", post.Post.ID, post.Post.Title, post.Post.Content)
			// 评论信息
			for _, comment := range post.Comments {
				fmt.Printf("评论ID：%d, 内容：%s, 用户ID：%d\n", comment.ID, comment.Content, comment.UserID)
			}
		}
	}

	// 2.2 查询评论数量最多的文章信息
	post := QueryMostCommentedPost(db)
	// 评论数量最多的文章信息
	fmt.Printf("评论数量最多的文章ID：%d, 标题：%s, 内容：%s\n", post.ID, post.Title, post.Content)
}

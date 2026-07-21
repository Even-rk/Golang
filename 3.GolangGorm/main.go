package main

import "fmt"

func main() {
	// 1. 初始化并创建数据库表
	db := CreateTable()

	// 2.1 查询某个用户发布的所有文章及其对应的评论信息
	userList := []User{}
	db.Find(&userList)
	for _, user := range userList {
		posts := QueryUserPostsAndComments(db, user.ID)
		for _, post := range posts {
			// 文章信息
			fmt.Printf("文章信息：%v\n", post)
		}
	}

	// 2.2 查询评论数量最多的文章信息
	post := QueryMostCommentedPost(db)
	// 评论数量最多的文章信息
	fmt.Printf("评论数量最多的文章信息：%v\n", post)
}

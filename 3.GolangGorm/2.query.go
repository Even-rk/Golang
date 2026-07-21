package main

import "gorm.io/gorm"

// 1.使用Gorm查询某个用户发布的所有文章及其对应的评论信息
func QueryUserPostsAndComments(db *gorm.DB, userID int) []Post {
	var posts []Post
	db.Model(&Post{}).Where("user_id = ?", userID).Find(&posts)
	return posts
}

// 2.编写Go代码，使用Gorm查询评论数量最多的文章信息。
func QueryMostCommentedPost(db *gorm.DB) Post {
	var post Post
	db.Model(&Post{}).Joins("LEFT JOIN comments ON posts.id = comments.post_id").
		Group("posts.id").
		Order("COUNT(comments.id) DESC").
		Limit(1).
		First(&post)
	return post
}

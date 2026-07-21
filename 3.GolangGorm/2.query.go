package main

import "gorm.io/gorm"

type PostWithComments struct {
	Post
	Comments []Comment // 一篇文章对应多条评论
}

// 1.使用Gorm查询某个用户发布的所有文章及其对应的评论信息
func QueryUserPostsAndComments(db *gorm.DB, userID int) []PostWithComments {
	// 定义返回的结构体切片
	var postsWithComments []PostWithComments
	// 查询用户发布的所有文章
	var posts []Post
	db.Model(&Post{}).Where("user_id = ?", userID).Find(&posts)

	// 查询文章对应的评论
	var comments []Comment
	for _, post := range posts {
		db.Model(&Comment{}).Where("post_id = ?", post.ID).Find(&comments)
		// 将文章和评论信息打包到结构体中
		postsWithComments = append(postsWithComments, PostWithComments{
			Post:     post,
			Comments: comments,
		})
	}
	return postsWithComments
}

// 2.编写Go代码，使用Gorm查询评论数量最多的文章信息。
func QueryMostCommentedPost(db *gorm.DB) Post {
	var post Post
	db.Model(&Post{}).Joins("LEFT JOIN comments ON posts.id = comments.post_id"). // 连表查询
											Group("posts.id").                // 按文章ID分组
											Order("COUNT(comments.id) DESC"). // 按评论数量降序排序
											Limit(1).                         // 只查询评论数量最多的文章
											First(&post)                      // 首条记录
	return post
}

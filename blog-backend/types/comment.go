package model

import "gorm.io/gorm"

// 评论模型
type Comment struct {
	gorm.Model
	UserID  int    `gorm:"column:user_id"`    // 关联用户ID
	PostID  int    `gorm:"column:post_id"`    // 关联文章ID
	Content string `gorm:"column:content"`    // 评论内容
	User    User   `gorm:"foreignKey:UserID"` // 关联用户
	Post    Post   `gorm:"foreignKey:PostID"` // 关联文章
}

// 创建评论请求结构
type CreateCommentRequest struct {
	PostID  int    `json:"post_id" binding:"required"`
	UserID  int    `json:"user_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}
